package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port     string
	LogLevel string
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Print("PORT environment not set, using default 8080")
		port = "8080"
	}

	logLevel := os.Getenv("LOG_LEVEL")
	if logLevel == "" {
		log.Print("LOG_LEVEL environment not set, using default INFO")
		logLevel = "INFO"
	}

	return Config{
		Port:     port,
		LogLevel: logLevel,
	}
}
