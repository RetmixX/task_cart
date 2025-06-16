package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"task_cart/internal/model/entity"
	"task_cart/pkg/db"
)

type ProductInterface interface {
	All() ([]entity.Product, error)
	ById(id uint) (*entity.Product, error)
	Create(data *entity.Product) (*entity.Product, error)
	Update(id uint, data *entity.Product) (*entity.Product, error)
	Delete(id uint) error
}

type ProductService struct {
	db *gorm.DB
}

func (p *ProductService) All() ([]entity.Product, error) {
	const op = "repository.product.All"
	var products []entity.Product

	if err := p.db.Find(&products).Error; err != nil {
		return nil, fmt.Errorf("%s: can't get products: %w", op, err)
	}

	return products, nil
}

func (p *ProductService) ById(id uint) (*entity.Product, error) {
	const op = "repository.product.ById"
	var product entity.Product

	if err := p.db.First(&product, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, db.EntityNotFoundErr
		}
		return nil, fmt.Errorf("%s: can't get product by id: %w", op, err)
	}

	return &product, nil
}

func (p *ProductService) Create(data *entity.Product) (*entity.Product, error) {
	const op = "repository.product.Create"
	if err := p.db.Create(data).Error; err != nil {
		return nil, fmt.Errorf("%s: can't create product: %w", op, err)
	}

	return data, nil
}

func (p *ProductService) Update(id uint, data *entity.Product) (*entity.Product, error) {
	const op = "repository.product.Update"

	var updated entity.Product

	if err := p.db.First(&updated, id).Updates(data).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, db.EntityNotFoundErr
		}

		return nil, fmt.Errorf("%s: can't update product id: %d, err: %w", op, id, err)
	}

	return &updated, nil
}

func (p *ProductService) Delete(id uint) error {
	const op = "repository.product.Delete"
	var deleted entity.Product
	if err := p.db.First(&deleted, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return db.EntityNotFoundErr
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := p.db.Delete(&deleted); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func NewProductService(conn *gorm.DB) *ProductService {
	return &ProductService{db: conn}
}
