package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"package/internal/domain"

	"github.com/segmentio/kafka-go"
)

func NewOrchestrator() {
	producer := NewProducer()
	reader := NewConsumer()

	defer func() {
		producer.Close()
		reader.Close()
	}()

	fmt.Println("activatePackage is waiting for messages...")

	// Continuously read messages from the topic
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Error reading message: %s\n", err)
			continue
		}

		log.Printf("RECEIVED MESSAGE : %s\n", string(m.Value))

		// Parse the incoming message
		var incoming domain.Message
		if err := json.Unmarshal(m.Value, &incoming); err != nil {
			log.Printf("Error parsing message: %s\n", err)
			continue
		}

		// Create a response message
		response := domain.Response{
			OrderType:     incoming.OrderType,
			OrderService:  "activatePackage",
			TransactionId: incoming.TransactionId,
			UserId:        incoming.UserId,
			PackageId:     incoming.PackageId,
			RespCode:      http.StatusOK,
			RespStatus:    "Success",
			RespMessage:   "Package is Successfully Activated",
		}

		// Marshal the response to JSON
		responseBytes, err := json.Marshal(response)
		if err != nil {
			log.Printf("Error marshalling response: %s\n", err)
			continue
		}

		// Produce the response message to topic_0
		err = producer.WriteMessages(context.Background(),
			kafka.Message{
				Key:   []byte(incoming.TransactionId), // Optional: set a key if needed
				Value: responseBytes,
			},
		)

		if err != nil {
			log.Printf("Error writing message to topic_0: %s\n", err)
			continue
		}

		log.Printf("RESPONSE SENT: %s\n", string(responseBytes))
	}
}
