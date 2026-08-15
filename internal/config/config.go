package config

import (
	"os"
)

type Config struct {
	DatabaseURL string
	ServerPort  string
}

func Load() *Config {
	return &Config{
		DatabaseURL: getEnv("DATABASE_URL", "host=localhost port=5432 user=nilmadhab dbname=ecommerce sslmode=disable"),
		ServerPort:  getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
