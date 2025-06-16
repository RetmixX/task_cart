package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title string  `gorm:"size:255,not null"`
	Price float32 `gorm:"not null"`
	Count int     `gorm:"not null"`
}
