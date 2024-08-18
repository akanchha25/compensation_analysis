// handlers/compensation_handler.go
package handlers

import (
    "encoding/json"
    "net/http"
    "compensation-api/services"
    "strconv"
)

type CompensationHandler struct {
    service services.CompensationService
}

func NewCompensationHandler(service services.CompensationService) *CompensationHandler {
    return &CompensationHandler{service: service}
}

func (h *CompensationHandler) GetCompensations(w http.ResponseWriter, r *http.Request) {
    filters := map[string]interface{}{}
    sortFields := map[string]string{}

    // Extract query parameters for filtering and sorting
    for key, values := range r.URL.Query() {
        if len(values) > 0 {
            switch {
            case key == "timestamp" || key == "salary":
                sortFields[key] = values[0]
            case key == "salary[gte]":
                if salaryGte, err := strconv.ParseFloat(values[0], 64); err == nil {
                    filters["base_salary"] = map[string]interface{}{
                        "gte": salaryGte,
                    }
                } else {
                    http.Error(w, "Invalid salary value", http.StatusBadRequest)
                    return
                }
            case key == "age":
                if age, err := strconv.Atoi(values[0]); err == nil {
                    filters["age"] = map[string]interface{}{
                        "gte": age, // For example, using "gte" for age
                    }
                } else {
                    http.Error(w, "Invalid age value", http.StatusBadRequest)
                    return
                }
            case key == "industry" || key == "job_title" || key == "currency" || key == "location":
                filters[key] = values[0]
            case key == "years_of_experience":
                if years, err := strconv.ParseFloat(values[0], 64); err == nil {
                    filters["years_of_experience"] = map[string]interface{}{
                        "gte": years, // For example, using "gte" for years_of_experience
                    }
                } else {
                    http.Error(w, "Invalid years of experience value", http.StatusBadRequest)
                    return
                }
            case key == "timestamp":
                filters["timestamp"] = values[0]
            default:
                http.Error(w, "Unsupported filter", http.StatusBadRequest)
                return
            }
        }
    }

    compensations, err := h.service.GetCompensations(filters, sortFields)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(compensations)
}
