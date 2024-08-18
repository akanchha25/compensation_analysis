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
            if key == "timestamp" || key == "salary" {
                sortFields[key] = values[0]
            } else if key == "salary[gte]" {
                if salaryGte, err := strconv.ParseFloat(values[0], 64); err == nil {
                    filters["base_salary"] = map[string]interface{}{
                        "gte": salaryGte,
                    }
                } else {
                    http.Error(w, "Invalid salary value", http.StatusBadRequest)
                    return
                }
            } else {
                filters[key] = values[0]
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
