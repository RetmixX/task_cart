package entity

import (
	"gorm.io/gorm"
	"task_cart/internal/model/dto"
)

type Status struct {
	gorm.Model
	Name   string `gorm:"not null"`
	Orders []Order
}

func (s Status) ToDTO() *dto.StatusDTO {
	return &dto.StatusDTO{
		Id:   s.ID,
		Name: s.Name,
	}
}
