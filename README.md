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


## Installation

#### 1. Clone the Repository

```bash
  git clone https://github.com/yourusername/compensation-api.git
  cd my-project
```
#### 2. Environment Configuration Create a .env file in the root of the project with the following content from .example.env:

```bash
# The URL of your Elasticsearch instance, which includes the protocol (http/https), 
# host (domain or IP address), and optionally the port number.
ELASTICSEARCH_URL=

# The username for accessing your secured Elasticsearch instance.
ELASTIC_USERNAME=

# The password corresponding to the above username for Elasticsearch authentication.
ELASTIC_PASSWORD=

# The name of the Elasticsearch index where your data will be stored.
INDEX_NAME=e

# Whether to verify SSL/TLS certificates when connecting to an HTTPS Elasticsearch instance.
# Typically set to "true" for production environments to ensure security, and "false" for development.
VERIFY_CERTS=

# The port on which your application will run.
PORT=
```

#### 3. Install Dependencies
```bash
 go mod tidy
```
#### 4. Run the Application
```bash
 go run main.go

# You should see logs confirming that the server is running: Starting server on port 8080...
```
#### 5. Health Check 
   ```bash
 curl http://localhost:8080/health
```
## API Reference

#### Get Compensation Data

```http
  GET /compensation_data
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `salary[gte]` | `string` | **Required**. Filter for salaries greater than or equal to a value |
| `zip_code` | `string` | **Required**. Filter by zip code.|
| `timestamp` | `string` | **Required**. Sort by timestamp. |




## Deployment

To deploy this project run

```bash
  npm run deploy
```






   
