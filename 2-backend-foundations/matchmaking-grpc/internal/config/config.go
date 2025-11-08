package config

import (
	"errors"
	"os"
	"time"

	"github.com/joho/godotenv"
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
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		PGHost:       "localhost",
		PGPort:       5432,
		PGSSLMODE:    "disable",
	}
}

func Load() (Config, error) {
	_ = godotenv.Load()
	cfg := defaultConfig()
	cfg.HTTPAddr = os.Getenv("PORT")
	cfg.PGUser = os.Getenv("PG_USER")
	cfg.PGPassword = os.Getenv("PG_PASSWORD")
	cfg.PGDB = os.Getenv("PG_DB")
	cfg.JWTSecret = os.Getenv("JWT_SECRET")

	if cfg.AppEnv == "prod" && cfg.JWTSecret == "" {
		return cfg, errors.New("JWT_SECRET required in prod")
	}

	return cfg, nil
}
