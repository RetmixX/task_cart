package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"task_cart/internal/model/entity"
	"task_cart/pkg/db"
)

type OrderInterface interface {
	ViewOrders() ([]entity.OrderWithCart, error)
	CreateOrder() (*entity.OrderWithCart, error)
	ChangeStatus(orderId uint, newStatus *entity.Status) (*entity.OrderWithCart, error)
}

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(conn *gorm.DB) *OrderRepository {
	return &OrderRepository{db: conn}
}

func (o *OrderRepository) ViewOrders() ([]entity.OrderWithCart, error) {
	const op = "repository.order.ViewOrders"
	var orders []entity.Order
	if err := o.db.Find(&orders).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	var orderData []entity.OrderWithCart

	for _, v := range orders {
		var result entity.OrderWithCart
		var status entity.Status
		if err := o.db.First(&status, v.StatusID).Error; err != nil {
			return nil, fmt.Errorf("%s: can't get status: %w", op, err)
		}
		result.Order = v
		result.Status = status
		err := o.db.Table("cart_products").Select("products.*, cart_products.quantity").
			Joins("join products on products.id = cart_products.product_id").
			Where("cart_products.cart_id = ?", v.CartID).Scan(&result.Products).Error
		if err != nil {
			return nil, fmt.Errorf("%s: error sql!: %w", op, err)
		}

		orderData = append(orderData, result)
	}

	return orderData, nil
}

func (o *OrderRepository) CreateOrder() (*entity.OrderWithCart, error) {
	const op = "repository.order.CreateOrder"

	var result entity.OrderWithCart
	var currentCart entity.Cart

	if err := o.db.Where("is_ordered = ?", false).First(&currentCart).Error; err != nil {
		return nil, fmt.Errorf("%s: can't get current cart: %w", op, err)
	}
	var issuedStatus entity.Status
	if err := o.db.Where("name = ?", "Issued").First(&issuedStatus).Error; err != nil {
		return nil, fmt.Errorf("%s: can't get issued status: %w", op, err)
	}
	newOrder := entity.Order{
		StatusID: issuedStatus.ID,
		CartID:   currentCart.ID,
	}

	if err := o.db.Create(&newOrder).Error; err != nil {
		return nil, fmt.Errorf("%s: can't create order: %w", op, err)
	}

	if err := o.db.Create(&entity.Cart{}).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	if err := o.db.Model(&currentCart).Update("is_ordered", true).Error; err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	err := o.db.Table("cart_products").Select("products.*, cart_products.quantity").
		Joins("join products on products.id = cart_products.product_id").
		Where("cart_products.cart_id = ?", currentCart.ID).Scan(&result.Products).Error

	if err != nil {
		return nil, fmt.Errorf("%s: can't get cart info: %w", op, err)
	}
	result.Order = newOrder
	result.Status = issuedStatus

	return &result, nil

}

func (o *OrderRepository) ChangeStatus(orderId uint, newStatus *entity.Status) (*entity.OrderWithCart, error) {
	const op = "repository.order.ChangeStatus"
	var result entity.OrderWithCart
	var order entity.Order
	if err := o.db.First(&order, orderId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, db.EntityNotFoundErr
		}

		return nil, fmt.Errorf("%s: can't get order by id: %w", op, err)
	}

	if err := o.db.Model(&order).Update("status_id", newStatus.ID).Error; err != nil {
		return nil, fmt.Errorf("%s: can't update status order: %w", op, err)
	}

	err := o.db.Table("cart_products").Select("products.*, cart_products.quantity").
		Joins("join products on products.id = cart_products.product_id").
		Where("cart_products.cart_id = ?", order.CartID).Scan(&result.Products).Error

	if err != nil {
		return nil, fmt.Errorf("%s: can't get cart info: %w", op, err)
	}
	result.Order = order
	result.Status = *newStatus

	return &result, nil
}
