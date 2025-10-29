package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"playerapi/internal/config"
	"playerapi/internal/player"
	"playerapi/internal/storage/postgres"
	"playerapi/internal/transport"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	// Set up Context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Load Config
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("failed to load application configuration")
	}

	// Create DB pool
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", cfg.PGHost, cfg.PGPort, cfg.PGUser, cfg.PGPassword, cfg.PGDB, cfg.PGSSLMODE)
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()
	if err := pool.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	// Set up layers with dependency injection
	playerRepo := postgres.NewPlayerRepoPG(pool)
	playerService := player.NewService(playerRepo)
	playerHandler := transport.NewPlayerHandler(playerService)

	// Create Server
	mux := http.NewServeMux()
	mux.HandleFunc("POST /register", playerHandler.CreatePlayer)
	mux.HandleFunc("GET /players/{id}", playerHandler.GetPlayer)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
		// Good production defaults:
		ReadHeaderTimeout: 5 * time.Second,  // mitigate slowloris
		ReadTimeout:       10 * time.Second, // whole request body
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	go func() {
		log.Printf("listening on %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Graceful shutdown on SIGINT/SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	}
	log.Println("server stopped")
}
