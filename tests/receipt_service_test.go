package tests

import (
	"testing"

	"github.com/Nivedithabp/receipt-processor-challenge/models"
	"github.com/Nivedithabp/receipt-processor-challenge/services"
)

func TestCalculatePoints(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []models.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
		},
		Total: "35.35",
	}

	points := services.ProcessReceipt(receipt)
	if points == "" {
		t.Errorf("Expected valid receipt ID but got an empty string")
	}
}
