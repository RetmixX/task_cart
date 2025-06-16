package entity

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	Name   string `gorm:"not null"`
	Orders []Order
}
