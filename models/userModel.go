package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"unique;not null"`
	Username string `gorm:"not null"`
	Password string `gorm:"not null"`
}
