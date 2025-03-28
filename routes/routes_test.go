package routes_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/Nivedithabp/receipt-processor-challenge/models"
	"github.com/Nivedithabp/receipt-processor-challenge/routes"
	"github.com/Nivedithabp/receipt-processor-challenge/services"
)

func TestProcessReceiptHandler(t *testing.T) {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	t.Run("Valid Receipt - Returns 200", func(t *testing.T) {
		receipt := models.Receipt{
			Retailer:     "Target",
			PurchaseDate: "2022-01-01",
			PurchaseTime: "13:01",
			Items: []models.Item{
				{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			},
			Total: "6.49",
		}
		payload, _ := json.Marshal(receipt)
		req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer(payload))
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		if status := recorder.Code; status != http.StatusOK {
			t.Errorf("Expected 200 but got %d", status)
		}
	})

	t.Run("Invalid Receipt - Missing Fields", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer([]byte(`{}`)))
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		if status := recorder.Code; status != http.StatusBadRequest {
			t.Errorf("Expected 400 but got %d", status)
		}
	})

	t.Run("Invalid JSON - Returns 400", func(t *testing.T) {
		req, _ := http.NewRequest("POST", "/receipts/process", bytes.NewBuffer([]byte(`{invalid`)))
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		if status := recorder.Code; status != http.StatusBadRequest {
			t.Errorf("Expected 400 but got %d", status)
		}
	})
}

func TestGetPointsHandler(t *testing.T) {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	// Create a valid receipt and store the ID
	receipt := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []models.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
		},
		Total: "6.49",
	}
	receiptID := services.ProcessReceipt(receipt)

	t.Run("Valid Receipt ID - Returns Points", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/receipts/"+receiptID+"/points", nil)
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		if status := recorder.Code; status != http.StatusOK {
			t.Errorf("Expected 200 but got %d", status)
		}

		var response map[string]int
		json.Unmarshal(recorder.Body.Bytes(), &response)
		if response["points"] == 0 {
			t.Errorf("Expected non-zero points but got %d", response["points"])
		}
	})

	t.Run("Invalid Receipt ID - Returns 404", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/receipts/invalid-id/points", nil)
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		if status := recorder.Code; status != http.StatusNotFound {
			t.Errorf("Expected 404 but got %d", status)
		}
	})
}

func TestHealthCheckHandler(t *testing.T) {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	t.Run("Health Check - Returns 200", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/health", nil)
		recorder := httptest.NewRecorder()

		router.ServeHTTP(recorder, req)

		if status := recorder.Code; status != http.StatusOK {
			t.Errorf("Expected 200 but got %d", status)
		}

		var response map[string]string
		json.Unmarshal(recorder.Body.Bytes(), &response)

		if response["status"] != "up" {
			t.Errorf("Expected status 'up' but got %s", response["status"])
		}

		if _, err := time.ParseDuration(response["uptime"]); err != nil {
			t.Errorf("Expected valid uptime but got %s", response["uptime"])
		}
	})
}
