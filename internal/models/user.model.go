package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	ID       uint
	Username string `gorm:"not null"`
	Email    string `gorm:"not null;unique_index"`
	Password string `gorm:"not null"`
	Avatar   string
	Role     string `gorm:"check:role IN ('user', 'admin')"`
	Links    []Link
}
