package models

import "time"

// Transaction represents a financial transaction
type Transaction struct {
	TransactionID string    `json:"transaction_id"`
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	CustomerID    string    `json:"customer_id"`
	Status        string    `json:"status"`
	GatewayID     int       `json:"gateway_id"`
	Type          string    `json:"type"` // Type of transaction: deposit, withdrawal, refund
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
