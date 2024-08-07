package main

import (
	"log"
	"net/http"
	"user_service/internal/config"
	"user_service/internal/database"
	"user_service/internal/handlers"
	"user_service/internal/kafka"

	"github.com/gorilla/mux"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Connect to database
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Initialize Kafka producer
	producer, err := kafka.NewProducer(cfg.Kafka.Broker, cfg.Kafka.Topic)
	if err != nil {
		log.Fatalf("Error initializing Kafka producer: %v", err)
	}
	defer producer.Close()

	// Setup router
	router := mux.NewRouter()
	router.HandleFunc("/register", handlers.RegisterUser(db, producer)).Methods("POST")

	// Start server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
