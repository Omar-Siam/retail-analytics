package handlers

import (
	"RetailAnalytics/internal/models"
	"encoding/json"
	"net/http"
)

// PostTransaction handles the POST request to create a new transaction.
func PostTransaction(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transaction
	if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response := models.Confirmation{
		Status:        "Created",
		TransactionID: transaction.TransactionID,
		ItemID:        transaction.ItemID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}
