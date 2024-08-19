// repositories/compensation_repository.go
package repositories

import (
	"bytes"
	"compensation-api/models"
	"context"
	"encoding/json"
	"errors"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
)

type CompensationRepository interface {
    Search(query *bytes.Buffer) ([]models.Compensation, error)
    FindByID(id string, fields []string) (*models.Compensation, error)
}

type compensationRepository struct {
    client    *elasticsearch.Client
    indexName string
}

func NewCompensationRepository(client *elasticsearch.Client, indexName string) CompensationRepository {
    return &compensationRepository{client: client, indexName: indexName}
}

func (r *compensationRepository) Search(query *bytes.Buffer) ([]models.Compensation, error) {
    req := esapi.SearchRequest{
        Index: []string{r.indexName},
        Body:  query,
    }

    res, err := req.Do(context.Background(), r.client)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    if res.IsError() {
        return nil, errors.New(res.String())
    }

    var result map[string]interface{}
    if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
        return nil, err
    }

    var compensations []models.Compensation
    hits := result["hits"].(map[string]interface{})["hits"].([]interface{})
    for _, hit := range hits {
        source := hit.(map[string]interface{})["_source"]
        var compensation models.Compensation
        data, _ := json.Marshal(source)
        json.Unmarshal(data, &compensation)
        compensations = append(compensations, compensation)
    }

    return compensations, nil
}

func (r *compensationRepository) FindByID(id string, fields []string) (*models.Compensation, error) {
    var sourceIncludes []string
    if len(fields) > 0 {
        sourceIncludes = fields
    }

    req := esapi.GetRequest{
        Index:      r.indexName,
        DocumentID: id,
        Source:     sourceIncludes,
    }

    res, err := req.Do(context.Background(), r.client)
    if err != nil {
        return nil, err
    }
    defer res.Body.Close()

    if res.IsError() {
        return nil, errors.New(res.String())
    }

    var result map[string]interface{}
    if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
        return nil, err
    }

    sourceData := result["_source"]
    if sourceData == nil {
        return nil, errors.New("document not found")
    }

    var compensation models.Compensation
    data, _ := json.Marshal(sourceData)
    json.Unmarshal(data, &compensation)

    return &compensation, nil
}
