package entity

import (
	"gorm.io/gorm"
	"task_cart/internal/model/dto"
)

type Product struct {
	gorm.Model
	Title string  `gorm:"size:255,not null"`
	Price float32 `gorm:"not null"`
	Count int     `gorm:"not null"`
}

func (p *Product) ToDTO() *dto.ProductDTO {
	return &dto.ProductDTO{
		Id:    p.ID,
		Title: p.Title,
		Price: p.Price,
	}
}

func (p *Product) ToWithCountDTO() *dto.ProductWithCountDTO {
	return &dto.ProductWithCountDTO{
		ProductDTO: *p.ToDTO(),
		Count:      p.Count,
	}
}
