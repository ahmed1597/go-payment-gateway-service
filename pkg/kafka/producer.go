package kafka

import (
	"context"
	"go-payment-gateway/internal/models"
	"log"

	"github.com/segmentio/kafka-go"
)

// ProduceStatusUpdate sends a transaction status update to Kafka
func ProduceStatusUpdate(transaction models.Transaction) error {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "transaction-status-updates",
	})

	message := []byte("TransactionID: " + transaction.TransactionID + ", Status: " + transaction.Status)

	err := writer.WriteMessages(context.Background(), kafka.Message{
		Value: message,
	})
	if err != nil {
		log.Printf("Failed to send status update to Kafka: %v", err)
		return err
	}

	log.Printf("Sent transaction status update to Kafka: %s", message)
	return nil
}
