package entity

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	IsOrdered bool      `gorm:"not null"`
	Products  []Product `gorm:"many2many:cart_products"`
}
