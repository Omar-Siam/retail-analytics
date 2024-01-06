package main

import (
	"RetailAnalytics/consumerservice/internal/clients"
	"RetailAnalytics/consumerservice/internal/kafka"
	"log"
)

func main() {
	const brokerAddress = "localhost:9092"
	const consumerGroupID = "test-consumer-group"
	const topic = "test-topic-part"
	const bucketName = "test-s3-bucket"

	s3client := clients.NewS3Client()

	consumer, err := kafka.NewConsumer([]string{brokerAddress}, consumerGroupID, topic, s3client, bucketName)
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
