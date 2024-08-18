// main.go
package main

import (
    "log"
    "compensation-api/config"
    "compensation-api/elasticsearch"
    "compensation-api/repositories"
    "compensation-api/services"
    "compensation-api/handlers"
    "compensation-api/server"
	"github.com/joho/godotenv"
	"fmt"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }

    cfg := config.LoadConfig()

	// Print out the loaded configuration to verify
    fmt.Printf("Elasticsearch URL: %s", cfg.ElasticsearchURL)
    fmt.Printf("Index Name: %s", cfg.IndexName)
    fmt.Printf("Port: %s", cfg.Port)

    esClient, err := elasticsearch.GetClient(cfg)
    if err != nil {
        log.Fatalf("Failed to create Elasticsearch client: %s", err)
    }

    // Use the index name from the configuration
    compensationRepo := repositories.NewCompensationRepository(esClient, cfg.IndexName)
    compensationService := services.NewCompensationService(compensationRepo)
    compensationHandler := handlers.NewCompensationHandler(compensationService)

    // Start the server on the configured port
    server.StartServer(compensationHandler, cfg.Port)
}
