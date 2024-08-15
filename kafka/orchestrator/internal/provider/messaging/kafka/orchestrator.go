package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"orchestrator/internal/domain"
	"orchestrator/internal/provider/database"
	"orchestrator/internal/repository"

	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
)

// StartConsumer starts consuming messages from the Kafka topic
func NewOrchestrator() {
	reader := NewConsumer()

	db, err := database.NewDB()
	if err != nil {
		log.Error().Err(err).Msg(fmt.Sprintf("Error connecting to the database: %v\n", err))
	}

	repo := repository.NewOrchestratorConfig(db)

	defer reader.Close()

	log.Info().Msg("Listening for messages...")

	for {
		// Read messages from the topic
		message, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Error().Err(err).Msg(fmt.Sprintf("Error while reading message: %v\n", err))
		}

		// Print the received message value
		log.Info().Msg(fmt.Sprintf("Received message: %s\n", string(message.Value)))

		// Parse the incoming message
		var incoming domain.Message
		if err := json.Unmarshal(message.Value, &incoming); err != nil {
			log.Error().Err(err).Msg(fmt.Sprintf("Error parsing message: %v\n", err))
			continue // Skip this message if parsing fails
		}

		ct := context.WithValue(context.Background(), domain.Key("type"), incoming.OrderType)
		ctx := context.WithValue(ct, domain.Key("service"), incoming.OrderService)

		orchestrate, err := repo.GetConfig(ctx)
		if err != nil {
			log.Error().Err(err).Msg(fmt.Sprintf("Error getting orchestration config: %v\n", err))
			continue // Skip this message if parsing fails
		}

		writer := NewProducer(orchestrate.TargetTopic)

		// Determine where to send or how to process the message
		if orchestrate.OrderService == "" {
			// If the message matches the first format, send to topic_validateUser
			responseBytes, _ := json.Marshal(incoming) // Ignore error for simplicity
			err = writer.WriteMessages(context.Background(),
				kafka.Message{
					Key:   []byte(incoming.TransactionId),
					Value: responseBytes,
				},
			)
			if err != nil {
				log.Error().Err(err).Msg(fmt.Sprintf("Error writing message to topic_%s: %v\n", orchestrate.TargetTopic, err))
				writer.Close()
				continue
			}
			log.Info().Msg(fmt.Sprintf("Message sent to topic_%s: %s\n", orchestrate.TargetTopic, string(responseBytes)))
		} else {
			// Log the completion message from the third format
			if incoming.RespCode == http.StatusOK && orchestrate.TargetTopic == "FINISH" {
				log.Printf("===============================================================================================")
				log.Printf("Transaction ID %s for order type '%s' is COMPLETED\n", incoming.TransactionId, incoming.OrderType)
				log.Printf("===============================================================================================")
				writer.Close()
				continue
			}
			if orchestrate.OrderService == incoming.OrderService {
				// If the message matches the second format, send to topic_activatePackage
				responseBytes, _ := json.Marshal(incoming) // Ignore error for simplicity
				if incoming.RespCode != http.StatusOK {
					writer.Close()
					continue
				}
				err = writer.WriteMessages(context.Background(),
					kafka.Message{
						Key:   []byte(incoming.TransactionId),
						Value: responseBytes,
					},
				)
				if err != nil {
					log.Error().Err(err).Msg(fmt.Sprintf("Error writing message to topic_%s: %v\n", orchestrate.TargetTopic, err))
					writer.Close()
					continue
				}
				log.Info().Msg(fmt.Sprintf("Message sent to topic_%s: %s\n", orchestrate.TargetTopic, string(responseBytes)))
			}
			writer.Close()
		}
	}
}
