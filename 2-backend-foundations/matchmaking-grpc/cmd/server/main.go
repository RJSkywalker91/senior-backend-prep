package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	pb "matchmaking/cmd/proto"
	"matchmaking/internal/config"
	"matchmaking/internal/player"
	"matchmaking/internal/storage/postgres"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func main() {
	// Set up Context
	initCtx, initCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer initCancel()

	// Load Config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("failed to load application configuration")
	}

	// Create DB pool
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", cfg.PGHost, cfg.PGPort, cfg.PGUser, cfg.PGPassword, cfg.PGDB, cfg.PGSSLMODE)
	pool, err := pgxpool.New(initCtx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	if err := pool.Ping(initCtx); err != nil {
		log.Fatal(err)
	}

	// Set up layers with dependency injection
	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 60 * time.Second,
			MaxConnectionAge:  0,
			Time:              30 * time.Second, // ping interval
			Timeout:           10 * time.Second, // ping timeout
		}),
		grpc.ConnectionTimeout(10*time.Second), // handshake timeout
	)
	playerRepo := postgres.NewPlayerRepoPG(pool)
	pb.RegisterPlayerServiceServer(s, player.NewPlayerServiceServer(playerRepo))

	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		log.Printf("gRPC server listening at %s", lis.Addr())
		if err := s.Serve(lis); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
			log.Fatalf("server error: %v", err)
		}
	}()

	// --- Graceful shutdown on SIGINT/SIGTERM
	runCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	<-runCtx.Done()

	// Give in-flight RPCs up to 10s to finish
	done := make(chan struct{})
	go func() {
		s.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		log.Println("server stopped gracefully")
	case <-time.After(10 * time.Second):
		log.Println("grace period expired; forcing stop")
		s.Stop()
	}
}
