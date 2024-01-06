package main

import (
	"RetailAnalytics/RetailDataRetrievalService/internal/client"
	"RetailAnalytics/RetailDataRetrievalService/internal/handlers"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	const port = "8081"

	r := mux.NewRouter()
	s3Client := client.NewS3Client()

	r.HandleFunc("/digest", handlers.RetrieveDataHandler(s3Client)).Methods("GET")

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
