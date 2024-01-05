package models

import "time"

// Transaction represents the data structure for a retail transaction.
type Transaction struct {
	TransactionID string    `json:"transactionID"`
	RetailerID    string    `json:"retailerID"`
	RetailerName  string    `json:"retailerName"`
	TimeStamp     time.Time `json:"timestamp"`
	PaymentType   string    `json:"paymentType"`
	Item          `json:"item"`
}

// Item represents the data structure for an item in a transaction.
type Item struct {
	ItemID string  `json:"itemID"`
	Price  float64 `json:"price"`
}
