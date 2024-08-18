// utils/query_builder.go
package utils

import (
    "bytes"
    "encoding/json"
)

func BuildQuery(filters map[string]interface{}, sortFields map[string]string) *bytes.Buffer {
    var buf bytes.Buffer
    query := map[string]interface{}{
        "query": map[string]interface{}{
            "bool": map[string]interface{}{
                "must": []interface{}{},
            },
        },
    }

    // Add filters to the query
    for field, value := range filters {
        switch v := value.(type) {
        case map[string]interface{}: // Handle range queries
            rangeQuery := map[string]interface{}{
                "range": map[string]interface{}{
                    field: v,
                },
            }
            query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(
                query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]interface{}),
                rangeQuery,
            )
        case float64, int: // Numeric fields
            query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(
                query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]interface{}),
                map[string]interface{}{
                    "range": map[string]interface{}{
                        field: map[string]interface{}{
                            "gte": v,
                        },
                    },
                },
            )
        case string: // Text or keyword fields
            query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"] = append(
                query["query"].(map[string]interface{})["bool"].(map[string]interface{})["must"].([]interface{}),
                map[string]interface{}{
                    "match": map[string]interface{}{
                        field: v,
                    },
                },
            )
        }
    }

    // Add sorting to the query
    if len(sortFields) > 0 {
        var sorts []map[string]interface{}
        for field, order := range sortFields {
            sorts = append(sorts, map[string]interface{}{
                field: map[string]interface{}{
                    "order": order,
                },
            })
        }
        query["sort"] = sorts
    }

    json.NewEncoder(&buf).Encode(query)
    return &buf
}
