package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Addr        string
	DatabaseURL string
	LogLevel    string
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		Addr:        getenv("ADDR", ":8080"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		LogLevel:    getenv("LOG_LEVEL", "info"),
	}
}
