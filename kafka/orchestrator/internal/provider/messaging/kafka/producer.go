package kafka

import (
	"fmt"

	"github.com/segmentio/kafka-go"
)

func NewProducer(topic string) *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:29092"},    // Your Kafka broker address
		Topic:   fmt.Sprintf("topic_%s", topic), // The topic to send messages for validation
	})
}