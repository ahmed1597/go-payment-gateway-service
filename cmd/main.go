package main

import (
	"context"
	"go-payment-gateway/configs"
	"go-payment-gateway/internal/db"
	"go-payment-gateway/internal/services"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration from .env
	cfg, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to PostgreSQL
	conn, err := db.ConnectPostgres(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer conn.Close(context.Background())

	// Load available gateways from the database
	err = services.LoadGateways(conn)
	if err != nil {
		log.Fatalf("Failed to load gateways: %v", err)
	}

	// Set up HTTP router
	router := mux.NewRouter()
	router.HandleFunc("/transaction/deposit", services.HandleDepositHTTP(conn)).Methods("POST")
	router.HandleFunc("/transaction/withdraw", services.HandleWithdrawalHTTP(conn)).Methods("POST")

	// Start the server
	log.Printf("Server is running on port %s", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, router))
}
