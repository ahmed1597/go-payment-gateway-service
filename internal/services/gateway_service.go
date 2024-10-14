package services

import (
	"context"
	"fmt"
	"go-payment-gateway/pkg/gateway"
	"log"

	"github.com/jackc/pgx/v4"
)

var availableGateways []gateway.Gateway

func LoadGateways(conn *pgx.Conn) error {
	rows, err := conn.Query(context.Background(), "SELECT id, name, priority FROM gateways")
	if err != nil {
		log.Printf("Error loading gateways: %v", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var priority int

		err := rows.Scan(&id, &name, &priority)
		if err != nil {
			return err
		}

		gateway, err := gateway.NewGateway(name, priority, id)
		if err != nil {
			return err
		}

		availableGateways = append(availableGateways, gateway)
	}

	return nil
}

// SelectBestGateway selects the gateway with the highest priority
func SelectBestGateway() (gateway.Gateway, error) {
	if len(availableGateways) == 0 {
		return nil, fmt.Errorf("no available gateways")
	}

	bestGateway := availableGateways[0]
	for _, gw := range availableGateways {
		if gw.GetPriority() < bestGateway.GetPriority() {
			bestGateway = gw
		}
	}

	return bestGateway, nil
}
