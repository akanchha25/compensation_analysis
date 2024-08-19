
package server

import (
    "log"
    "net/http"
    "compensation-api/handlers"
    "github.com/gorilla/mux"
)

func NewRouter(handler *handlers.CompensationHandler) *mux.Router {
    r := mux.NewRouter()
    
    r.HandleFunc("/compensation_data", handler.GetCompensations).Methods("GET")

    r.HandleFunc("/compensation_data/id", handler.GetCompensationByID).Methods("GET")

    // Health check endpoint
    r.HandleFunc("/health", HealthCheckHandler).Methods("GET")

    return r
}

// HealthCheckHandler is the handler for the health check route.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Server is healthy"))
}

func StartServer(handler *handlers.CompensationHandler, port string) {
    r := NewRouter(handler)

    // Log before starting the server
    log.Printf("Starting server on port %s...", port)

    // Start the server
    err := http.ListenAndServe(":"+port, r)
    if err != nil {
        log.Fatalf("Server failed to start: %s", err)
    }

    // Log that the server has started successfully
    log.Printf("Server started successfully on port %s", port)
}
