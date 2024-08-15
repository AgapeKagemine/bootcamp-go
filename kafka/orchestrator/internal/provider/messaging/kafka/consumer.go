package kafka

import "github.com/segmentio/kafka-go"

func NewConsumer() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:29092"}, // Your Kafka broker address
		Topic:       "topic_0",                   // The topic to consume from
		GroupID:     "my-consumer-group",         // Consumer group ID
		MinBytes:    10e3,                        // Minimum bytes to read
		MaxBytes:    10e6,                        // Maximum bytes to read
		StartOffset: kafka.LastOffset,            // Start consuming from the end (new messages)
	})
}
