package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Product_name string  `gorm:"not null; type:varchar(255)"`
	Price        float64 `gorm:"not null; type:decimal"`
	Stock        int     `gorm:"not null; type:integer; default:0"`
}
