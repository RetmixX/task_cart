package service

import (
	"task_cart/internal/model/consts"
	"task_cart/internal/model/dto"
	"task_cart/internal/repository"
)

type ProductInterface interface {
	All() ([]dto.ProductWithCountDTO, error)
}

type ProductService struct {
	productRepo repository.ProductInterface
}

func NewProductService(productRepo repository.ProductInterface) *ProductService {
	return &ProductService{productRepo: productRepo}
}

func (p *ProductService) All() ([]dto.ProductWithCountDTO, error) {
	var productData []dto.ProductWithCountDTO
	products, err := p.productRepo.All()

	if err != nil {
		return nil, consts.ServerErr
	}

	for _, v := range products {
		productData = append(productData, *v.ToWithCountDTO())
	}

	return productData, nil
}
