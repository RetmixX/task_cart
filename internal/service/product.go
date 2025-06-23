package service

import (
	"log/slog"
	"task_cart/internal/model/consts"
	"task_cart/internal/model/dto"
	"task_cart/internal/repository"
	"task_cart/pkg/log/sl"
)

type ProductInterface interface {
	All() ([]dto.ProductWithCountDTO, error)
}

type ProductService struct {
	productRepo repository.ProductInterface
	logger      *slog.Logger
}

func NewProductService(productRepo repository.ProductInterface, log *slog.Logger) *ProductService {
	return &ProductService{productRepo: productRepo, logger: log}
}

func (p *ProductService) All() ([]dto.ProductWithCountDTO, error) {
	const op = "service.product.All"
	log := p.logger.With(slog.String("op", op))
	log.Info("Start All()")
	var productData []dto.ProductWithCountDTO
	products, err := p.productRepo.All()

	if err != nil {
		log.Error("fail get products ", sl.Err(err))
		return nil, consts.ServerErr
	}

	for _, v := range products {
		productData = append(productData, *v.ToWithCountDTO())
	}
	log.Info("End All()")
	return productData, nil
}
