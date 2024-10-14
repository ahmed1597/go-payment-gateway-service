package gateway

import "go-payment-gateway/internal/models"

// Gateway defines the interface that all gateways must implement
type Gateway interface {
	ProcessDeposit(transaction models.Transaction) (string, error)
	ProcessWithdrawal(transaction models.Transaction) (string, error)
	GetTransactionStatus(transactionID string) (models.Transaction, error)
	GetPriority() int
	GetID() int
}
