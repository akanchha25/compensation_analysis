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

```
## Prerequisites

Go: Ensure you have Go installed (version 1.18 or higher).
Elasticsearch: Ensure you have an Elasticsearch cluster running.


## Setup
1. Clone the Repository
    ```git clone https://github.com/yourusername/compensation-api.git```
    ```cd compensation-api```

2. Environment Configuration
   Create a .env file in the root of the project with the following content from .example.env:

3. Install Dependencies
   ```go mod tidy```

4. Run the Application
   ```go run main.go```
You should see logs confirming that the server is running:
Starting server on port 8080...

5. Health Check
```curl http://localhost:8080/health```




   
