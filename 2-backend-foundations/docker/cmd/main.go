package main

import (
	"context"
	"docker/internal/handlers"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	// Create Server
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", handlers.HealthCheck)
	mux.HandleFunc("/", handlers.Echo)

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
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("graceful shutdown failed: %v", err)
	}
	log.Println("server exited gracefully")
}
