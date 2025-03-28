package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/Nivedithabp/receipt-processor-challenge/models"
	"github.com/Nivedithabp/receipt-processor-challenge/services"
)

// App start time for uptime calculation
var appStartTime = time.Now()

// HealthResponse defines the response structure for /health
type HealthResponse struct {
	Status string `json:"status"`
	Uptime string `json:"uptime"`
}

// HealthCheckHandler returns application health status
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(appStartTime).Round(time.Second)
	response := HealthResponse{
		Status: "up",
		Uptime: uptime.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// RegisterRoutes registers all routes
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/receipts/process", ProcessReceiptHandler).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", GetPointsHandler).Methods("GET")
	router.HandleFunc("/health", HealthCheckHandler).Methods("GET")
}
