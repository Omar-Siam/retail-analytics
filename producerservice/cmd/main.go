package main

import (
	"RetailAnalytics/producerservice/internal/clients"
	"RetailAnalytics/producerservice/internal/kafka"
	"RetailAnalytics/producerservice/internal/repository"
	"github.com/IBM/sarama"
	"log"
	"time"
)

func main() {
	const brokerAddress = "localhost:9092"

	config := sarama.NewConfig()
	sqsClient := clients.NewSQSClient()

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

	for {
		msgs, err := repository.PollSQS(sqsClient)
		if err != nil {
			log.Printf("Failed to poll from SQS: %v", err)
			continue
		}

		for _, msg := range msgs {
			if err := producer.PostTransaction(msg); err != nil {
				log.Printf("Failed to send to Kafka: %v", err)
			} else {
				// Delete message from SQS queue after successful send to Kafka
				repository.DeleteMessageFromSQS(sqsClient, msg)
			}
		}
		time.Sleep(1 * time.Second) // Poll interval
	}

}
