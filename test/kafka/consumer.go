package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	ID string

	ready chan bool
}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Printf("[%s] Message claimed: value = %s, timestamp = %v, topic= = %s\n", c.ID, string(msg.Value), msg.Timestamp, msg.Topic)
		session.MarkMessage(msg, "")
	}
	return nil
}
