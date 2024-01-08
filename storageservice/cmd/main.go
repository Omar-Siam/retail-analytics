package main

import (
	"RetailAnalytics/storageservice/internal/clients"
	"RetailAnalytics/storageservice/internal/handlers"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	const port = "8083"

	r := mux.NewRouter()
	sqsClient := clients.NewSQSClient()

	r.HandleFunc("/ingest", handlers.PostToSQSHandler(sqsClient)).Methods("POST")

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Printf("Starting server on port %s", port)
	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Server failed: %v", err)
	}
}
