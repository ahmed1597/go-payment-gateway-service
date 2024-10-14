package gateway

import (
	"encoding/json"
	"fmt"
	"go-payment-gateway/internal/models"
	"time"
)

// GatewayA implements the JSON-based payment gateway
type GatewayA struct {
	Name     string
	Priority int
	ID       int
}

// NewGatewayA creates a new instance of GatewayA
func NewGatewayA(priority int, id int) *GatewayA {
	return &GatewayA{
		Name:     "GatewayA",
		Priority: priority,
		ID:       id,
	}
}

// ProcessDeposit handles deposits for GatewayA
func (g *GatewayA) ProcessDeposit(transaction models.Transaction) (string, error) {
	requestBody, _ := json.Marshal(transaction)
	fmt.Printf("Processing deposit for GatewayA (JSON): %s\n", requestBody)
	return "Deposit processed successfully via GatewayA", nil
}

// ProcessWithdrawal handles withdrawals for GatewayA
func (g *GatewayA) ProcessWithdrawal(transaction models.Transaction) (string, error) {
	requestBody, _ := json.Marshal(transaction)
	fmt.Printf("Processing withdrawal for GatewayA (JSON): %s\n", requestBody)
	return "Withdrawal processed successfully via GatewayA", nil
}

// GetTransactionStatus simulates retrieving the transaction status from GatewayA
func (g *GatewayA) GetTransactionStatus(transactionID string) (models.Transaction, error) {
	if transactionID == "" {
		return models.Transaction{}, fmt.Errorf("invalid transaction ID")
	}

	// Simulated response based on transactionID
	status := "pending"
	if len(transactionID) > 5 {
		status = "completed"
	} else if len(transactionID) == 5 {
		status = "failed"
	}

	// Simulate the retrieved transaction details
	transaction := models.Transaction{
		TransactionID: transactionID,
		Status:        status,
		Type:          "deposit",
		UpdatedAt:     time.Now(),
	}

	return transaction, nil
}

// GetPriority returns the priority of the gateway
func (g *GatewayA) GetPriority() int {
	return g.Priority
}

// GetID returns the ID of the gateway
func (g *GatewayA) GetID() int {
	return g.ID
}
