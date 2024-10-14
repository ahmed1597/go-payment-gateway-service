package services

import (
	"context"
	"encoding/json"
	"go-payment-gateway/internal/models"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v4"
)

// InsertTransaction inserts a new transaction into the database, including type
func InsertTransaction(conn *pgx.Conn, transaction models.Transaction) error {
	_, err := conn.Exec(context.Background(),
		"INSERT INTO transactions (transaction_id, amount, currency, customer_id, status, gateway_id, type, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $8)",
		transaction.TransactionID, transaction.Amount, transaction.Currency, transaction.CustomerID, transaction.Status, transaction.GatewayID, transaction.Type, time.Now(),
	)
	if err != nil {
		log.Printf("Failed to insert transaction: %v", err)
	}
	return err
}

// HandleDeposit processes a deposit request, assigning "deposit" as the type
func HandleDeposit(conn *pgx.Conn, transaction models.Transaction) (string, error) {
	// Select the best gateway from gateway service
	selectedGateway, err := SelectBestGateway()
	if err != nil {
		return "", err
	}

	// Assign the correct gateway ID and type
	transaction.GatewayID = selectedGateway.GetID()
	transaction.Type = "deposit"
	transaction.Status = "pending"

	// Insert transaction into database
	err = InsertTransaction(conn, transaction)
	if err != nil {
		return "", err
	}

	// Process the deposit via the selected gateway
	response, err := selectedGateway.ProcessDeposit(transaction)
	if err != nil {
		log.Printf("Failed to process deposit: %v", err)
		UpdateTransactionStatus(conn, transaction.TransactionID, "failed")
		return "", err
	}

	// Update transaction status to successful
	err = UpdateTransactionStatus(conn, transaction.TransactionID, "successful")
	if err != nil {
		return "", err
	}

	return response, nil
}

// HandleDepositHTTP is the HTTP handler for deposits
func HandleDepositHTTP(conn *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transaction models.Transaction
		// Decode the JSON request body
		if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		// Call the HandleDeposit function
		response, err := HandleDeposit(conn, transaction)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

// HandleWithdrawalHTTP is the HTTP handler for withdrawals
func HandleWithdrawalHTTP(conn *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transaction models.Transaction
		// Decode the JSON request body
		if err := json.NewDecoder(r.Body).Decode(&transaction); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}
		// Call the HandleWithdrawal function
		response, err := HandleWithdrawal(conn, transaction)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
}

// HandleWithdrawal processes a withdrawal request, assigning "withdrawal" as the type
func HandleWithdrawal(conn *pgx.Conn, transaction models.Transaction) (string, error) {
	// Select the best gateway from gateway service
	selectedGateway, err := SelectBestGateway()
	if err != nil {
		return "", err
	}

	// Assign the correct gateway ID and type
	transaction.GatewayID = selectedGateway.GetID()
	transaction.Type = "withdrawal"
	transaction.Status = "pending"

	// Insert transaction into database
	err = InsertTransaction(conn, transaction)
	if err != nil {
		return "", err
	}

	// Process the withdrawal via the selected gateway
	response, err := selectedGateway.ProcessWithdrawal(transaction)
	if err != nil {
		log.Printf("Failed to process withdrawal: %v", err)
		UpdateTransactionStatus(conn, transaction.TransactionID, "failed")
		return "", err
	}

	// Update transaction status to successful
	err = UpdateTransactionStatus(conn, transaction.TransactionID, "successful")
	if err != nil {
		return "", err
	}

	return response, nil
}

// UpdateTransactionStatus updates only the status and updated_at of a transaction in the database
func UpdateTransactionStatus(conn *pgx.Conn, transactionID, status string) error {
	_, err := conn.Exec(context.Background(),
		"UPDATE transactions SET status=$1, updated_at=$2 WHERE transaction_id=$3",
		status, time.Now(), transactionID,
	)
	if err != nil {
		log.Printf("Failed to update transaction status: %v", err)
		return err
	}
	return nil
}
