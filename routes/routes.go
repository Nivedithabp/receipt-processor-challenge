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

// RegisterRoutes registers all routes
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/receipts/process", ProcessReceiptHandler).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", GetPointsHandler).Methods("GET")
	router.HandleFunc("/health", HealthCheckHandler).Methods("GET")
}
// @Summary Process a receipt and generate an ID
// @Description Submits a receipt and returns a unique ID
// @Tags Receipts
// @Accept json
// @Produce json
// @Param receipt body models.Receipt true "Receipt JSON"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /receipts/process [post]
func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil || !isValidReceipt(receipt) {
		http.Error(w, `{"error": "Invalid request payload"}`, http.StatusBadRequest)
		return
	}

	id := services.ProcessReceipt(receipt)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

// isValidReceipt validates required fields in the receipt
func isValidReceipt(receipt models.Receipt) bool {
	if receipt.Retailer == "" ||
		receipt.PurchaseDate == "" ||
		receipt.PurchaseTime == "" ||
		len(receipt.Items) == 0 ||
		receipt.Total == "" {
		return false
	}
	return true
}

// @Summary Get points for a receipt
// @Description Returns points for a given receipt ID
// @Tags Receipts
// @Produce json
// @Param id path string true "Receipt ID"
// @Success 200 {object} map[string]int
// @Failure 404 {object} map[string]string
// @Router /receipts/{id}/points [get]
func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	points, exists := services.GetPoints(id)
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]int{"points": points})
}

// @Summary Check health of the API
// @Description Returns application uptime and status
// @Tags Health
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	uptime := time.Since(appStartTime).Round(time.Second)
	response := HealthResponse{
		Status: "up",
		Uptime: uptime.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

