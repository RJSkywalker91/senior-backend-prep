package config

import (
	"errors"
	"time"
)

type Config struct {
	AppEnv       string
	HTTPAddr     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration

	PGHost     string
	PGPort     int
	PGUser     string
	PGPassword string
	PGDB       string
	PGSSLMODE  string

	JWTSecret string
}

func defaultConfig() Config {
	return Config{
		AppEnv:       "dev",
		HTTPAddr:     ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		PGHost:       "localhost",
		PGPort:       5432,
		PGUser:       "goapi",
		PGPassword:   "", // TODO: Add way to dynamically get password
		PGDB:         "matchmaking_db",
		PGSSLMODE:    "disable",
	}
}

func Load() (Config, error) {
	cfg := defaultConfig()

	if cfg.AppEnv == "prod" && cfg.JWTSecret == "" {
		return cfg, errors.New("JWT_SECRET required in prod")
	}

	return cfg, nil
}
