package entity

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	Orders []Order
}
