package services_test

import (
	"testing"

	"github.com/Nivedithabp/receipt-processor-challenge/models"
	"github.com/Nivedithabp/receipt-processor-challenge/services"
)

func TestCalculatePoints(t *testing.T) {
	tests := []struct {
		name     string
		receipt  models.Receipt
		expected int
	}{
		{
			name: "Valid Receipt - All Rules Applied",
			receipt: models.Receipt{
				Retailer:     "Target",
				PurchaseDate: "2022-01-01",
				PurchaseTime: "13:01",
				Items: []models.Item{
					{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
					{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
					{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
					{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
					{ShortDescription: "   Klarbrunn 12-PK 12 FL OZ  ", Price: "12.00"},
				},
				Total: "35.35",
			},
			expected: 28,
		},
		{
			name: "Receipt with Round Dollar Total - 50 Points",
			receipt: models.Receipt{
				Retailer:     "Walmart",
				PurchaseDate: "2022-02-02",
				PurchaseTime: "14:30",
				Items: []models.Item{
					{ShortDescription: "Snickers Bar", Price: "1.00"},
				},
				Total: "1.00",
			},
			expected: 93, 
		},
		{
			name: "Total is a Multiple of 0.25 - 25 Points",
			receipt: models.Receipt{
				Retailer:     "Aldi",
				PurchaseDate: "2022-03-15",
				PurchaseTime: "15:01",
				Items: []models.Item{
					{ShortDescription: "Apple", Price: "5.00"},
				},
				Total: "5.00",
			},
			expected: 95,
		},
		{
			name: "Empty Items List - Minimal Points",
			receipt: models.Receipt{
				Retailer:     "Kroger",
				PurchaseDate: "2022-03-01",
				PurchaseTime: "10:00",
				Items:        []models.Item{},
				Total:        "0.00",
			},
			expected: 87,
		},
		{
			name: "Single Item with Description Length Multiple of 3",
			receipt: models.Receipt{
				Retailer:     "Meijer",
				PurchaseDate: "2022-03-19",
				PurchaseTime: "12:45",
				Items: []models.Item{
					{ShortDescription: "Banana", Price: "3.00"},
				},
				Total: "3.00",
			},
			expected: 88,
		},
		{
			name: "Odd Day Purchase - 6 Points",
			receipt: models.Receipt{
				Retailer:     "Store",
				PurchaseDate: "2022-03-05",
				PurchaseTime: "13:00",
				Items: []models.Item{
					{ShortDescription: "Water", Price: "1.00"},
				},
				Total: "1.00",
			},
			expected: 86,
		},
		{
			name: "Time Between 2PM and 4PM - 10 Points",
			receipt: models.Receipt{
				Retailer:     "Costco",
				PurchaseDate: "2022-04-10",
				PurchaseTime: "14:45",
				Items: []models.Item{
					{ShortDescription: "Milk", Price: "3.25"},
				},
				Total: "3.25",
			},
			expected: 41, 
		},
		{
			name: "No Points Condition Met",
			receipt: models.Receipt{
				Retailer:     "Gas Station",
				PurchaseDate: "2022-03-20",
				PurchaseTime: "05:00",
				Items: []models.Item{
					{ShortDescription: "Soda", Price: "1.00"},
				},
				Total: "1.00",
			},
			expected: 85, // Only retailer points + 1 item point + 6 odd day
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			points := services.CalculatePoints(test.receipt)
			if points != test.expected {
				t.Errorf("Expected %d points but got %d", test.expected, points)
			}
		})
	}
}

// Test ProcessReceipt and GetPoints together
func TestProcessAndGetPoints(t *testing.T) {
	receipt := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "14:30",
		Items: []models.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
		},
		Total: "6.49",
	}

	// Test ProcessReceipt to generate ID
	id := services.ProcessReceipt(receipt)
	if id == "" {
		t.Error("Expected a valid receipt ID but got an empty string")
	}

	// Test GetPoints to retrieve points
	points, exists := services.GetPoints(id)
	if !exists {
		t.Error("Expected points to exist but got false")
	}

	// Validate points
	expectedPoints := 22 
	if points != expectedPoints {
		t.Errorf("Expected %d points but got %d", expectedPoints, points)
	}
}

// Test GetPoints with invalid receipt ID
func TestGetPoints_InvalidID(t *testing.T) {
	points, exists := services.GetPoints("invalid-id")
	if exists {
		t.Error("Expected no points to exist for invalid ID, but got true")
	}
	if points != 0 {
		t.Errorf("Expected 0 points for invalid ID, but got %d", points)
	}
}
