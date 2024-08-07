package handlers

import (
	"encoding/json"
	"net/http"
	"user_service/internal/kafka"
	"user_service/internal/models"
	"user_service/internal/security"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func RegisterUser(db *gorm.DB, producer *kafka.Producer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Generate UUID for user ID
		user.ID = uuid.New().String()

		// Hash the password
		hashedPassword, err := security.HashPassword(user.PasswordHash)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}
		user.PasswordHash = hashedPassword

		// Save user to Redis (simulate by directly saving to DB for simplicity)
		if err := db.Create(&user).Error; err != nil {
			http.Error(w, "Failed to save user", http.StatusInternalServerError)
			return
		}

		// Produce user data to Kafka
		if err := producer.Produce(user); err != nil {
			http.Error(w, "Failed to send user data to Kafka", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}
