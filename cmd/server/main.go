package main

import (
	"log"
	"value-index/api"
	"value-index/internal/config"
	"value-index/internal/provider"
	"value-index/internal/search"
	"value-index/logger"
)

func main() {
	cfg := config.Load()

	logLevel, err := logger.LogLevelFromString(cfg.LogLevel)
	if err != nil {
		log.Fatalf("failed to parse log level")
	}
	appLogger := logger.NewStdLogger(logLevel, nil)

	fileProvider := provider.NewFileProvider("input.txt")

	searchService, err := search.NewSearchService(fileProvider, appLogger)
	if err != nil {
		log.Fatalf("failed to load data: %v", err)
	}

	apiService := api.NewRest(searchService, cfg.Port)
	log.Fatal(apiService.Start())
}
