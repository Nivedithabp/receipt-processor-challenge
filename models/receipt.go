package models

// Receipt defines the receipt structure
type Receipt struct {
	Retailer     string `json:"retailer"`
	PurchaseDate string `json:"purchaseDate"`
	PurchaseTime string `json:"purchaseTime"`
	Items        []Item `json:"items"`
	Total        string `json:"total"`
}

// Item defines the item structure
type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}
