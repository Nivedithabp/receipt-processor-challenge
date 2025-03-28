package services

import (
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Nivedithabp/receipt-processor-challenge/models"
	"github.com/google/uuid"
)

var receipts = make(map[string]int)

// ProcessReceipt processes the receipt and generates an ID
func ProcessReceipt(receipt models.Receipt) string {
	id := uuid.New().String()
	points := calculatePoints(receipt)
	receipts[id] = points
	return id
}

// GetPoints returns points for a given receipt ID
func GetPoints(id string) (int, bool) {
	points, exists := receipts[id]
	return points, exists
}

// calculatePoints calculates the points based on rules
func calculatePoints(receipt models.Receipt) int {
	points := 0

	// 1. One point per alphanumeric character in retailer name
	re := regexp.MustCompile(`[a-zA-Z0-9]`)
	points += len(re.FindAllString(receipt.Retailer, -1))

	// 2. 50 points if total is a round dollar amount
	if strings.HasSuffix(receipt.Total, ".00") {
		points += 50
	}

	// 3. 25 points if total is a multiple of 0.25
	total, _ := strconv.ParseFloat(receipt.Total, 64)
	if math.Mod(total, 0.25) == 0 {
		points += 25
	}

	// 4. 5 points for every two items
	points += (len(receipt.Items) / 2) * 5

	// 5. Description length multiple of 3 - 20% price points
	for _, item := range receipt.Items {
		desc := strings.TrimSpace(item.ShortDescription)
		if len(desc)%3 == 0 {
			price, _ := strconv.ParseFloat(item.Price, 64)
			points += int(math.Ceil(price * 0.2))
		}
	}

	// 6. 6 points if purchase day is odd
	purchaseDate, _ := time.Parse("2006-01-02", receipt.PurchaseDate)
	if purchaseDate.Day()%2 != 0 {
		points += 6
	}

	// 7. 10 points for purchase between 2:00 PM and 4:00 PM
	purchaseTime, _ := time.Parse("15:04", receipt.PurchaseTime)
	if purchaseTime.Hour() >= 14 && purchaseTime.Hour() < 16 {
		points += 10
	}

	return points
}
