package kafka

import (
	"context"
	"go-payment-gateway/internal/services"
	"log"
	"strings"

	"github.com/jackc/pgx/v4"
	"github.com/segmentio/kafka-go"
)

// ConsumeStatusUpdates listens for status updates from Kafka and updates the transaction status in the database
func ConsumeStatusUpdates(conn *pgx.Conn) {
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "transaction-status-updates",
		GroupID: "transaction-status-group",
	})

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("Failed to read message from Kafka: %v", err)
			continue
		}

		log.Printf("Received Kafka message: %s", string(msg.Value))

		// Parse message and update transaction status in the database
		transactionID, status := parseKafkaMessage(string(msg.Value))
		err = services.UpdateTransactionStatus(conn, transactionID, status)
		if err != nil {
			log.Printf("Failed to update transaction status: %v", err)
		}
	}
}

func parseKafkaMessage(message string) (string, string) {
	parts := strings.Split(message, ", ")
	transactionID := strings.Replace(parts[0], "TransactionID: ", "", 1)
	status := strings.Replace(parts[1], "Status: ", "", 1)
	return transactionID, status
}
