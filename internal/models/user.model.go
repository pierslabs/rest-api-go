package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	ID       uuid.UUID `gorm:"primaryKey"`
	Username string    `gorm:"not null"`
	Email    string    `gorm:"not null;unique"`
	Password string    `gorm:"not null"`
}
