// internal/config/config.go
package config

import (
	"os"
	"strconv"
)

type Config struct {
	Database struct {
		Host     string
		Port     int
		User     string
		Password string
		DBName   string
		SSLMode  string
	}
	Redis struct {
		Host string
		Port int
	}
	Kafka struct {
		Broker string
		Topic  string
	}
}

func LoadConfig() (*Config, error) {
	var cfg Config

	// Database configuration
	cfg.Database.Host = os.Getenv("DB_HOST")
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return nil, err
	}
	cfg.Database.Port = port
	cfg.Database.User = os.Getenv("DB_USER")
	cfg.Database.Password = os.Getenv("DB_PASSWORD")
	cfg.Database.DBName = os.Getenv("DB_NAME")
	cfg.Database.SSLMode = "disable" // Or "require" if SSL is configured

	// Redis configuration
	cfg.Redis.Host = os.Getenv("REDIS_HOST")
	redisPort, err := strconv.Atoi(os.Getenv("REDIS_PORT"))
	if err != nil {
		return nil, err
	}
	cfg.Redis.Port = redisPort

	// Kafka configuration
	cfg.Kafka.Broker = os.Getenv("KAFKA_BROKER")
	cfg.Kafka.Topic = os.Getenv("KAFKA_TOPIC")

	return &cfg, nil
}
