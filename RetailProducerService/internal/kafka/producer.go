package kafka

import (
	"RetailAnalytics/RetailProducerService/internal/models"
	"encoding/json"
	"github.com/IBM/sarama"
	"log"
)

const TopicName = "test-topic-part"

type Producer struct {
	sarama.SyncProducer
}

// NewProducer creates a new Kafka producer.
func NewProducer(brokers []string, config *sarama.Config) (*Producer, error) {
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}
	return &Producer{producer}, nil
}

// PostTransaction sends a transaction to Kafka.
func (p *Producer) PostTransaction(transaction models.Transaction) error {
	body, err := json.Marshal(transaction)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: TopicName,
		Value: sarama.ByteEncoder(body),
	}

	partition, offset, err := p.SendMessage(msg)
	log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", TopicName, partition, offset)
	return err
}
