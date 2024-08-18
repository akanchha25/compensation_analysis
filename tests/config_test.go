// package tests

// import (
//     "os"
//     "testing"
//     "compensation-api/config"
// )

// func TestLoadConfig(t *testing.T) {
//     os.Setenv("ELASTICSEARCH_URL", "http://localhost:9200")
//     os.Setenv("INDEX_NAME", "test_index")
//     os.Setenv("PORT", "8080")

//     cfg := config.LoadConfig()

//     if cfg.ElasticsearchURL != "http://localhost:9200" {
//         t.Errorf("Expected Elasticsearch URL to be 'http://localhost:9200', got '%s'", cfg.ElasticsearchURL)
//     }
//     if cfg.IndexName != "test_index" {
//         t.Errorf("Expected Index Name to be 'test_index', got '%s'", cfg.IndexName)
//     }
//     if cfg.Port != "8080" {
//         t.Errorf("Expected Port to be '8080', got '%s'", cfg.Port)
//     }
// }


package tests

import (
    "os"
    "testing"
    "compensation-api/config"
    "github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
    os.Setenv("ELASTICSEARCH_URL", "https://localhost:9200")
    os.Setenv("INDEX_NAME", "compensation_data")
    os.Setenv("PORT", "8080")
    os.Setenv("VERIFY_CERTS", "false")
    os.Setenv("ELASTIC_USERNAME", "elastic")
    os.Setenv("ELASTIC_PASSWORD", "OMdt*psB6iJNUFHTOsQi")

    cfg := config.LoadConfig()

    assert.Equal(t, "https://localhost:9200", cfg.ElasticsearchURL)
    assert.Equal(t, "compensation_data", cfg.IndexName)
    assert.Equal(t, "8080", cfg.Port)
}
