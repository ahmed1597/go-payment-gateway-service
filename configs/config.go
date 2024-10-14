package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost       string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	KafkaBrokers string
	ServerPort   string
	DatabaseURL  string
}

// LoadConfig loads the configuration from environment variables and `.env` file
func LoadConfig() (*Config, error) {
	// Load environment variables from the `.env` file
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found.")
	}

	cfg := &Config{
		DBHost:       getEnv("DB_HOST", "localhost"),
		DBPort:       getEnv("DB_PORT", "5432"),
		DBUser:       getEnv("DB_USER", "postgres"),
		DBPassword:   getEnv("DB_PASSWORD", "password"),
		DBName:       getEnv("DB_NAME", "payment_db"),
		KafkaBrokers: getEnv("KAFKA_BROKERS", "kafka:9092"),
		ServerPort:   getEnv("SERVER_PORT", "8080"),
	}

	// Construct the Database URL
	cfg.DatabaseURL = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	log.Printf("Loaded configuration: %+v\n", cfg)
	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
