
package config

import (
    "os"
)

type Config struct {
    ElasticsearchURL string
    Username         string
    Password         string
    IndexName        string
    Port             string
    VerifyCerts      bool
}

func LoadConfig() *Config {
    return &Config{
        ElasticsearchURL: os.Getenv("ELASTICSEARCH_URL"),
        Username:         os.Getenv("ELASTIC_USERNAME"),
        Password:         os.Getenv("ELASTIC_PASSWORD"),
        IndexName:        os.Getenv("INDEX_NAME"),
        Port:             getEnv("PORT", "8080"),
        VerifyCerts:      getEnv("VERIFY_CERTS", "false") == "true",
    }
}

func getEnv(key, fallback string) string {
    value := os.Getenv(key)
    if value == "" {
        return fallback
    }
    return value
}

