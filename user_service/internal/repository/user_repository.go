package repository

import (
	"context"
	"encoding/json"
	"time"
	model "user_service/internal/models"

	"github.com/go-redis/redis/v8"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

type UserRepository struct {
	rdb         *redis.Client
	kafkaWriter *kafka.Writer
	db          *gorm.DB
}

func NewUserRepository(rdb *redis.Client, kafkaWriter *kafka.Writer, db *gorm.DB) *UserRepository {
	return &UserRepository{rdb: rdb, kafkaWriter: kafkaWriter, db: db}
}

func (repo *UserRepository) Save(user *model.User) error {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// Сохранение в PostgreSQL
	err := repo.db.Create(user).Error
	if err != nil {
		return err
	}

	// Сохранение в Redis
	userData, err := json.Marshal(user)
	if err != nil {
		return err
	}
	err = repo.rdb.Set(context.Background(), user.Email, userData, 0).Err()
	if err != nil {
		return err
	}

	// Отправка в Kafka
	err = repo.kafkaWriter.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(user.Email),
		Value: userData,
	})
	return err
}
