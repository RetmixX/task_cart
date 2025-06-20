package entity

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	CartID   uint    `gorm:"index"`
	StatusID uint    `gorm:"index"`
	Amount   float32 `gorm:"not null"`
}
