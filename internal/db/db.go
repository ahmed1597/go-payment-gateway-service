package db

import (
	"context"
	"go-payment-gateway/configs"
	"log"

	"github.com/jackc/pgx/v4"
)

// ConnectPostgres sets up the PostgreSQL connection using the individual config values
func ConnectPostgres(cfg *configs.Config) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Printf("Unable to connect to database: %v\n", err)
		return nil, err
	}

	log.Println("Connected to the database successfully.")
	return conn, nil
}
