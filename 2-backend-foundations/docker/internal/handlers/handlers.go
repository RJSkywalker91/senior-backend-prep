package handlers

import (
	"fmt"
	"net/http"
	"os"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("ok"))
}

func Echo(w http.ResponseWriter, r *http.Request) {
	msg := os.Getenv("APP_MESSAGE")
	if msg == "" {
		msg = "hello from Go in Docker 👋"
	}
	fmt.Fprintln(w, msg)
}
