package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model

	Title  string `gorm:"not null;unique"`
	Link   string `gorm:"not null;"`
	Done   bool   `gorm:"default:false"`
	UserID uint
	User   User `gorm:"foreignKey:UserID"`
}
