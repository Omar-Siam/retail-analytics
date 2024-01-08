package models

// Confirmation represents the data structure for a response to a transaction posting.
type Confirmation struct {
	Status        string `json:"confirmationID"`
	TransactionID string `json:"transactionID"`
	ItemID        string `json:"itemID"`
}
