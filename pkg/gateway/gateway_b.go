package gateway

import (
	"encoding/xml"
	"fmt"
	"go-payment-gateway/internal/models"
	"time"
)

// GatewayB implements the XML-based payment gateway
type GatewayB struct {
	Name     string
	Priority int
	ID       int
}

// NewGatewayB creates a new instance of GatewayB
func NewGatewayB(priority int, id int) *GatewayB {
	return &GatewayB{
		Name:     "GatewayB",
		Priority: priority,
		ID:       id,
	}
}

// ProcessDeposit handles deposits for GatewayB
func (g *GatewayB) ProcessDeposit(transaction models.Transaction) (string, error) {
	requestBody, _ := xml.Marshal(transaction)
	fmt.Printf("Processing deposit for GatewayB (XML): %s\n", requestBody)
	return "Deposit processed successfully via GatewayB", nil
}

// ProcessWithdrawal handles withdrawals for GatewayB
func (g *GatewayB) ProcessWithdrawal(transaction models.Transaction) (string, error) {
	requestBody, _ := xml.Marshal(transaction)
	fmt.Printf("Processing withdrawal for GatewayB (XML): %s\n", requestBody)
	return "Withdrawal processed successfully via GatewayB", nil
}

// GetTransactionStatus simulates retrieving the transaction status from GatewayB
func (g *GatewayB) GetTransactionStatus(transactionID string) (models.Transaction, error) {
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
		Type:          "withdrawal",
		UpdatedAt:     time.Now(),
	}

	return transaction, nil
}

// GetPriority returns the priority of the gateway
func (g *GatewayB) GetPriority() int {
	return g.Priority
}

// GetID returns the ID of the gateway
func (g *GatewayB) GetID() int {
	return g.ID
}
