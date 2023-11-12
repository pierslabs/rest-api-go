package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"not null"`
	Email    string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}
