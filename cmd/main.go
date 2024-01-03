package main

import (
	"RetailAnalytics/internal/handlers"
	"errors"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	const port = "8080"

	r := mux.NewRouter()
	r.HandleFunc("/ingest", handlers.PostTransaction).Methods("POST")

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
