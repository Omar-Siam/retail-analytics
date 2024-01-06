package kafka

import (
	"RetailAnalytics/consumerservice/internal/repository"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"sync"
	"time"

	"github.com/IBM/sarama"
)

type Consumer struct {
	ready      chan bool
	client     sarama.ConsumerGroup
	topic      string
	s3Client   *s3.S3
	bucketName string
}

func NewConsumer(brokers []string, groupID string, topic string, s3Client *s3.S3, bucketName string) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_1_0
	config.Consumer.Return.Errors = true

	group, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, err
	}

	consumer := &Consumer{
		ready:      make(chan bool),
		client:     group,
		topic:      topic,
		s3Client:   s3Client,
		bucketName: bucketName,
	}

	return consumer, nil
}

func (c *Consumer) ConsumeMessages() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			if err := c.client.Consume(ctx, []string{c.topic}, c); err != nil {
				log.Panicf("Error from consumer: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
			c.ready = make(chan bool)
		}
	}()

	<-c.ready
	log.Println("Consumer running...")

	wg.Wait()
}

// Below are sarama.ConsumerGroupHandler interface methods. Required.

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Message claimed: value %s, timestamp %v, topic %s\n", string(message.Value), message.Timestamp, message.Topic)

		var jsonObject map[string]any
		err := json.Unmarshal(message.Value, &jsonObject)
		if err != nil {
			log.Printf("Error unmarshalling JSON: %s", err)
			continue
		}

		objectKey := fmt.Sprintf("%s/%d/%s.json", message.Topic, message.Partition, message.Timestamp.Format(time.RFC3339))
		err = repository.PutObjectToS3(c.s3Client, c.bucketName, objectKey, jsonObject)
		if err != nil {
			log.Printf("Error uploading to S3: %s", err)
			continue
		}

		session.MarkMessage(message, "")
	}
	return nil
}

func (c *Consumer) Close() error {
	if err := c.client.Close(); err != nil {
		return err
	}
	return nil
}
