package services

import (
	"os"

	"github.com/gox7/shorturl/model"
)

func NewConfig(config *model.Config) {
	*config = model.Config{
		SERVER_HOST:       get("SERVER_HOST", "0.0.0.0"),
		SERVER_PORT:       get("SERVER_PORT", "8080"),
		SERVER_PASSWORD:   get("SERVER_PASSWORD", "cipher-password"),
		POSTGRES_HOST:     get("POSTGRES_HOST", "127.0.0.1"),
		POSTGRES_PORT:     get("POSTGRES_PORT", "5432"),
		POSTGRES_USER:     get("POSTGRES_USER", "user"),
		POSTGRES_PASSWORD: get("POSTGRES_PASSWORD", "password"),
		POSTGRES_NAME:     get("POSTGRES_NAME", "dbname"),
	}
}

func get(key string, def string) string {
	value := os.Getenv(key)
	if value == "" {
		return def
	}
	return value
}
