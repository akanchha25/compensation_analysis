// handlers/compensation_handler.go
package handlers

import (
    "encoding/json"
    "net/http"
    "compensation-api/services"
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
