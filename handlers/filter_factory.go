// handlers/filter_factory.go
package handlers

import (
    "net/http"
    "strconv"
    "fmt"
)

type FilterFactory interface {
    CreateFilter(r *http.Request) (map[string]interface{}, map[string]string, error)
}

type DefaultFilterFactory struct{}

func NewDefaultFilterFactory() FilterFactory {
    return &DefaultFilterFactory{}
}

func (f *DefaultFilterFactory) CreateFilter(r *http.Request) (map[string]interface{}, map[string]string, error) {
    filters := make(map[string]interface{})
    sortFields := make(map[string]string)

    for key, values := range r.URL.Query() {
        if len(values) > 0 {
            switch key {
            case "timestamp", "salary":
                sortFields[key] = values[0]
            case "salary[gte]":
                salaryGte, err := strconv.ParseFloat(values[0], 64)
                if err != nil {
                    return nil, nil, fmt.Errorf("invalid salary value: %v", err)
                }
                filters["base_salary"] = map[string]interface{}{
                    "gte": salaryGte,
                }
            case "age":
                age, err := strconv.Atoi(values[0])
                if err != nil {
                    return nil, nil, fmt.Errorf("invalid age value: %v", err)
                }
                filters["age"] = map[string]interface{}{
                    "gte": age,
                }
            case "industry", "job_title", "currency", "location":
                filters[key] = values[0]
            case "years_of_experience":
                years, err := strconv.ParseFloat(values[0], 64)
                if err != nil {
                    return nil, nil, fmt.Errorf("invalid years of experience value: %v", err)
                }
                filters["years_of_experience"] = map[string]interface{}{
                    "gte": years,
                }
            // case "timestamp":
            //     filters["timestamp"] = values[0]
            default:
                return nil, nil, fmt.Errorf("unsupported filter: %s", key)
            }
        }
    }

    return filters, sortFields, nil
}
