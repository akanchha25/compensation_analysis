package main

import (
    "log"
    "fmt"
    "compensation-api/config"
    "compensation-api/elasticsearch"
    "compensation-api/repositories"
    "compensation-api/services"
    "compensation-api/handlers"
    "compensation-api/server"
    "github.com/joho/godotenv"
    "github.com/sirupsen/logrus"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }

    cfg := config.LoadConfig()

    // Print out the loaded configuration to verify
    fmt.Printf("Elasticsearch URL: %s\n", cfg.ElasticsearchURL)
    fmt.Printf("Index Name: %s\n", cfg.IndexName)
    fmt.Printf("Port: %s\n", cfg.Port)

    // Initialize Elasticsearch client
    esClient, err := elasticsearch.GetClient(cfg)
    if err != nil {
        log.Fatalf("Failed to create Elasticsearch client: %s", err)
    }

    // Set up logrus to send logs to Elasticsearch
    hook := elasticsearch.NewElasticsearchHook(esClient, "application-logs")
    logrus.AddHook(hook)

    // Use the index name from the configuration
    repo := repositories.NewCompensationRepository(esClient, cfg.IndexName)
    service := services.NewCompensationService(repo)
    filterFactory := handlers.NewDefaultFilterFactory()
    handler := handlers.NewCompensationHandler(service, filterFactory)

    // Start the server on the configured port
    server.StartServer(handler, cfg.Port)
}
