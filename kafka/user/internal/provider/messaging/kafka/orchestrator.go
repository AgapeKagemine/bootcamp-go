package kafka

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"user/internal/domain"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

func NewOrchestrator() {
	producer := NewProducer()
	reader := NewConsumer()

	defer func() {
		producer.Close()
		reader.Close()
	}()

	fmt.Println("validateUser is waiting for messages...")

	// Continuously read messages from the topic
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("Error reading message")
			continue
		}

		log.Printf("RECEIVED MESSAGE : %s\n", string(m.Value))

		// Parse the incoming message
		var incoming domain.Message
		if err := json.Unmarshal(m.Value, &incoming); err != nil {
			log.Error().Err(err).Msg("Error parsing message")
			continue
		}

		userService := "http://127.0.0.1:8090/user/valid"
		userRequest := domain.UserRequest{
			UserId: incoming.UserId,
		}

		payload, err := json.Marshal(userRequest)
		if err != nil {
			log.Error().Err(err).Msg("Error marshalling user request")
			continue
		}

		res, err := http.Post(userService, "application/json", bytes.NewBuffer(payload))
		if err != nil {
			log.Error().Err(err).Msg("Error sending user request")
			continue
		}

		byteRes, err := io.ReadAll(res.Body)
		if err != nil {
			log.Error().Err(err).Msg("Error reading user response")
			continue
		}

		var userResponse domain.UserResponse
		err = json.Unmarshal(byteRes, &userResponse)
		if err != nil {
			log.Error().Err(err).Msg("Error decoding user response")
			continue
		}

		// Create a response message
		response := domain.Response{
			OrderType:     incoming.OrderType,
			OrderService:  "validateUser",
			TransactionId: incoming.TransactionId,
			UserId:        incoming.UserId,
			PackageId:     incoming.PackageId,
			RespCode:      http.StatusOK,
			RespStatus:    "Success",
			RespMessage:   "User is valid",
		}

		if userResponse.Message != "valid" {
			response.RespCode = http.StatusBadRequest
			response.RespStatus = "Failed"
			response.RespMessage = "User invalid"
		}

		// Marshal the response to JSON
		responseBytes, err := json.Marshal(response)
		if err != nil {
			log.Error().Err(err).Msg("Error marshalling response")
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
			log.Error().Err(err).Msg("Error writing message to topic_0")
			continue
		}

		log.Printf("RESPONSE SENT: %s\n", string(responseBytes))
	}
}
