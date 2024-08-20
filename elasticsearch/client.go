
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
