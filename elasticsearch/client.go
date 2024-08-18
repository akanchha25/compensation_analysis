// elasticsearch/client.go
// package elasticsearch

// import (
//     "log"
//     "sync"

//     "github.com/elastic/go-elasticsearch/v8"
//     "compensation-api/config"
// )

// var once sync.Once
// var client *elasticsearch.Client

// func GetClient(cfg *config.Config) *elasticsearch.Client {
//     once.Do(func() {
//         es, err := elasticsearch.NewClient(elasticsearch.Config{
//             Addresses: []string{cfg.ElasticsearchURL},
//         })
//         if err != nil {
//             log.Fatalf("Error creating the Elasticsearch client: %s", err)
//         }
//         client = es
//     })
//     return client
// }


// elasticsearch/elasticsearch.go
// package elasticsearch

// import (
//     "log"
//     "github.com/elastic/go-elasticsearch/v8"
// )

// func GetClient() (*elasticsearch.Client, error) {
//     cfg := elasticsearch.Config{
//         Addresses: []string{
//             "http://localhost:9200", // Replace with your Elasticsearch endpoint
//         },
//     }

//     es, err := elasticsearch.NewClient(cfg)
//     if err != nil {
//         log.Fatalf("Error creating the client: %s", err)
//         return nil, err
//     }

//     return es, nil
// }



package elasticsearch

import (
    "compensation-api/config"
    "crypto/tls"
    "log"
    "net/http"

    "github.com/elastic/go-elasticsearch/v8"
)

func GetClient(cfg *config.Config) (*elasticsearch.Client, error) {
    esCfg := elasticsearch.Config{
        Addresses: []string{cfg.ElasticsearchURL},
        Username:  cfg.Username,
        Password:  cfg.Password,
        Transport: &http.Transport{
            TLSClientConfig: &tls.Config{
                InsecureSkipVerify: !cfg.VerifyCerts, // Skip certificate verification if VerifyCerts is false
            },
        },
    }

    es, err := elasticsearch.NewClient(esCfg)
    if err != nil {
        log.Fatalf("Error creating the client: %s", err)
        return nil, err
    }

    return es, nil
}
