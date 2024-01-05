package main

import (
	"RetailAnalytics/RetailConsumerService/internal/kafka"
	"log"
)

func main() {
	const brokerAddress = "localhost:9092"
	const consumerGroupID = "test-consumer-group"
	const topic = "test-topic-part"

	consumer, err := kafka.NewConsumer([]string{brokerAddress}, consumerGroupID, topic)
	if err != nil {
		log.Fatalln("Failed to start consumer:", err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Println("Failed to close consumer:", err)
		}
	}()

	consumer.ConsumeMessages()
}
