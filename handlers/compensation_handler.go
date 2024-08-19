// handlers/compensation_handler.go
package handlers

import (
	"compensation-api/services"
	"encoding/json"
	"net/http"
	"strings"
)

type CompensationHandler struct {
    service       services.CompensationService
    filterFactory FilterFactory
}

func NewCompensationHandler(service services.CompensationService, filterFactory FilterFactory) *CompensationHandler {
    return &CompensationHandler{service: service, filterFactory: filterFactory}
}

func (h *CompensationHandler) GetCompensations(w http.ResponseWriter, r *http.Request) {
    filters, sortFields, err := h.filterFactory.CreateFilter(r)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    compensations, err := h.service.GetCompensations(filters, sortFields)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(compensations)
}

// handlers/compensation_handler.go
func (h *CompensationHandler) GetCompensationByID(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "Missing id parameter", http.StatusBadRequest)
        return
    }

    fields := r.URL.Query().Get("fields")
    var fieldList []string
    if fields != "" {
        fieldList = strings.Split(fields, ",")
    }

    compensation, err := h.service.GetCompensationByID(id, fieldList)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(compensation)
}
