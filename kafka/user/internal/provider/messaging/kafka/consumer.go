package kafka

import "github.com/segmentio/kafka-go"

func NewConsumer() *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{"localhost:29092"}, // Change this to your Kafka server
		Topic:       "topic_validateUser",
		GroupID:     "my-consumer-group",
		StartOffset: kafka.LastOffset, // Start consuming from the end (new messages)
	})
}
