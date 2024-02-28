package datamodel

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID   string `gorm:"unique"`
	Password string
	Email    string `gorm:"unique"`
	Cash     float64
}
