-- Create the gateways table
CREATE TABLE IF NOT EXISTS gateways (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    priority INTEGER NOT NULL,
    data_format VARCHAR NOT NULL,
    protocol VARCHAR NOT NULL
);

-- Create the transactions table
CREATE TABLE IF NOT EXISTS transactions (
    transaction_id VARCHAR PRIMARY KEY,
    amount NUMERIC NOT NULL,
    currency VARCHAR NOT NULL,
    customer_id VARCHAR NOT NULL,
    status VARCHAR NOT NULL,
    gateway_id INTEGER REFERENCES gateways(id),
    type VARCHAR NOT NULL,  -- Transaction type: deposit, withdrawal, refund
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Insert initial data for Gateways
INSERT INTO gateways (id, name, priority, data_format, protocol) 
VALUES (1, 'GatewayA', 1, 'JSON', 'HTTP') 
ON CONFLICT (id) DO NOTHING;

INSERT INTO gateways (id, name, priority, data_format, protocol) 
VALUES (2, 'GatewayB', 2, 'XML', 'HTTP') 
ON CONFLICT (id) DO NOTHING;


