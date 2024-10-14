package gateway

import (
	"fmt"
)

// NewGateway dynamically creates a gateway based on its name
func NewGateway(name string, priority int, id int) (Gateway, error) {
	switch name {
	case "GatewayA":
		return NewGatewayA(priority, id), nil
	case "GatewayB":
		return NewGatewayB(priority, id), nil
	default:
		return nil, fmt.Errorf("unsupported gateway: %s", name)
	}
}
