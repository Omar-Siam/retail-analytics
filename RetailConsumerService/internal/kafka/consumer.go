package kafka

import (
	"context"
	"log"
	"sync"

	"github.com/IBM/sarama"
)

type Consumer struct {
	ready  chan bool
	client sarama.ConsumerGroup
	topic  string
}

func NewConsumer(brokers []string, groupID string, topic string) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_1_0
	config.Consumer.Return.Errors = true

	group, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, err
	}

	consumer := &Consumer{
		ready:  make(chan bool),
		client: group,
		topic:  topic,
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

	<-c.ready // Wait till the consumer has been set up
	log.Println("Consumer running...")

	wg.Wait()
}

// Implementing sarama.ConsumerGroupHandler interface methods

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Message claimed: value %s, timestamp %v, topic %s, partition %v\n", string(message.Value), message.Timestamp, message.Topic, message.Partition)
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
