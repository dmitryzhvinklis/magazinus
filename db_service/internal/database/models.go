package database

import (
	"time"

	"gorm.io/gorm"
)

// User представляет собой модель пользователя с основной информацией
type User struct {
	gorm.Model
	Username     string    `gorm:"size:100;unique;not null"` // Имя пользователя
	Email        string    `gorm:"size:100;unique;not null"` // Электронная почта
	PasswordHash string    `gorm:"size:255;not null"`        // Хеш пароля
	FirstName    string    `gorm:"size:100"`                 // Имя
	LastName     string    `gorm:"size:100"`                 // Фамилия
	CreatedAt    time.Time `gorm:"autoCreateTime"`           // Дата создания
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`           // Дата последнего обновления
}

// Migrate выполняет миграцию моделей в базе данных
func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}
