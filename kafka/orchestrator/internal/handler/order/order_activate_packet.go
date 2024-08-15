package order

import (
	"context"
	"log"
	"net/http"
	"orchestrator/internal/domain"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
)

func ActivatePackage(c *gin.Context) {
	var orderReq domain.OrderRequest
	if err := c.ShouldBindJSON(&orderReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Create a new writer for Kafka
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{"localhost:29092"},
		Topic:    "topic_0",
		Balancer: &kafka.LeastBytes{}, // Use the least bytes balancer for efficiency
	})
	defer writer.Close()

	// Create a message to send to Kafka
	message := kafka.Message{
		Key:   []byte(orderReq.TransactionID), // Use Transaction ID as key
		Value: []byte(`{"orderType": "` + orderReq.OrderType + `", "transactionId": "` + orderReq.TransactionID + `", "userId": "` + orderReq.UserId + `", "packageId": "` + orderReq.PackageId + `"}`),
	}

	// Send the message to Kafka
	err := writer.WriteMessages(context.Background(), message)
	if err != nil {
		log.Printf("Failed to produce message: %s\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to produce Kafka message"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Order placed successfully", "order": orderReq})
}
