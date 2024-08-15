package kafka

import "github.com/segmentio/kafka-go"

func NewProducer() *kafka.Writer {
	return kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:29092"}, // Change this to your Kafka server
		Topic:    "topic_0",
		Balancer: &kafka.LeastBytes{},
	})
}
