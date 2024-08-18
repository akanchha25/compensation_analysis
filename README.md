# Compensation API

The Compensation API is a Go-based RESTful service designed to query compensation data stored in an Elasticsearch cluster. This API supports filtering, sorting, and retrieving compensation records based on various fields.

## Project Structure

```plaintext
/compensation-api
├── main.go                    # Application entry point
├── config
│   └── config.go              # Configuration loading and environment setup
├── server
│   └── server.go              # Server setup and route handling
├── handlers
│   └── compensation_handler.go # HTTP handlers for the API endpoints
├── services
│   └── compensation_service.go # Business logic and service layer
├── repositories
│   └── compensation_repository.go # Data access layer for Elasticsearch
├── models
│   └── compensation.go        # Data models representing compensation records
├── elasticsearch
│   └── client.go              # Elasticsearch client setup and configuration
└── utils
    └── query_builder.go       # Utilities for building Elasticsearch queries

