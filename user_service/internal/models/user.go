package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           string    `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Username     string    `gorm:"size:100;unique;not null"`
	Email        string    `gorm:"size:100;unique;not null"`
	PasswordHash string    `gorm:"size:255;not null"`
	FirstName    string    `gorm:"size:100"`
	LastName     string    `gorm:"size:100"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
