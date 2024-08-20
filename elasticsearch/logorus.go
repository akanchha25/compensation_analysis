// logrus_elasticsearch_hook.go
package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/esapi"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/sirupsen/logrus"
)

// ElasticsearchHook sends logrus logs to Elasticsearch
type ElasticsearchHook struct {
    client *elasticsearch.Client
    index  string
}

func NewElasticsearchHook(client *elasticsearch.Client, index string) *ElasticsearchHook {
    return &ElasticsearchHook{client: client, index: index}
}

func (hook *ElasticsearchHook) Levels() []logrus.Level {
    return logrus.AllLevels
}

func (hook *ElasticsearchHook) Fire(entry *logrus.Entry) error {
    logData, err := json.Marshal(entry)
    if err != nil {
        return err
    }

    req := esapi.IndexRequest{
        Index:      hook.index,
        Body:       bytes.NewReader(logData),
        Refresh:    "true",
    }

    res, err := req.Do(context.Background(), hook.client)
    if err != nil {
        return err
    }
    defer res.Body.Close()

    if res.IsError() {
        return fmt.Errorf("failed to index document: %s", res.String())
    }

    return nil
}
