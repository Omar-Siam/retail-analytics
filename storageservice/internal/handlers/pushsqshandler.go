package handlers

import (
	"RetailAnalytics/storageservice/internal/models"
	"RetailAnalytics/storageservice/internal/repository"
	"encoding/json"
	"github.com/aws/aws-sdk-go/service/sqs"
	"log"
	"net/http"
)

func PostToSQSHandler(queue *sqs.SQS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transaction models.Transaction
		if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if err := repository.SendMessageToSQS(queue, transaction); err != nil {
			log.Printf("Failed to add to queue: %v", err)
			http.Error(w, "Failed to process", http.StatusInternalServerError)
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
}
