package kafka

import (
	"github.com/IBM/sarama"
	"github.com/aws/aws-sdk-go/service/sqs"
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
func (p *Producer) PostTransaction(msg *sqs.Message) error {
	msgkafka := &sarama.ProducerMessage{
		Topic: TopicName,
		Value: sarama.StringEncoder(*msg.Body),
	}

	partition, offset, err := p.SendMessage(msgkafka)
	log.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", TopicName, partition, offset)
	return err
}
