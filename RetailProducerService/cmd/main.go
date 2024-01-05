package main

import (
	"RetailAnalytics/RetailProducerService/internal/handlers"
	"RetailAnalytics/RetailProducerService/internal/kafka"
	"errors"
	"github.com/IBM/sarama"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	const port = "8080"
	const brokerAddress = "localhost:9092"

	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.Partitioner = sarama.NewRoundRobinPartitioner

	producer, err := kafka.NewProducer([]string{brokerAddress}, config)
	if err != nil {
		log.Fatal("Failed to start Kafka producer:", err)
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Printf("Failed to close Kafka producer: %v", err)
		}
	}()

	r := mux.NewRouter()
	r.HandleFunc("/ingest", handlers.PostTransaction(producer)).Methods("POST")

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
