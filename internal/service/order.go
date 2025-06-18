package service

import (
	"errors"
	"task_cart/internal/model/consts"
	"task_cart/internal/model/dto"
	"task_cart/internal/repository"
	"task_cart/pkg/db"
)

type OrderInterface interface {
	ViewOrders() ([]dto.OrderDTO, error)
	CreateOrder() (*dto.OrderDTO, error)
	ChangeStatus(orderId uint, statusId uint) (*dto.OrderDTO, error)
}

// NEED LOGGER!!!!!
type OrderService struct {
	orderRepo  repository.OrderInterface
	statusRepo repository.StatusInterface
	cartRepo   repository.CartInterface
}

func NewOrderService(orderRepo repository.OrderInterface,
	statusRepo repository.StatusInterface,
	cartRepo repository.CartInterface) *OrderService {
	return &OrderService{
		orderRepo:  orderRepo,
		statusRepo: statusRepo,
		cartRepo:   cartRepo,
	}
}

func (o *OrderService) ViewOrders() ([]dto.OrderDTO, error) {
	var ordersData []dto.OrderDTO
	orders, err := o.orderRepo.ViewOrders()

	if err != nil {
		return nil, consts.ServerErr
	}

	for _, v := range orders {
		ordersData = append(ordersData, *v.ToDTO())
	}

	return ordersData, nil

}

func (o *OrderService) CreateOrder() (*dto.OrderDTO, error) {
	cart, err := o.cartRepo.ViewCart()
	if err != nil {
		return nil, consts.ServerErr
	}

	if len(cart.Products) == 0 {
		return nil, consts.InvalidRequest
	}

	newOrder, err := o.orderRepo.CreateOrder()
	if err != nil {
		return nil, consts.ServerErr
	}

	return newOrder.ToDTO(), nil
}

func (o *OrderService) ChangeStatus(orderId uint, statusId uint) (*dto.OrderDTO, error) {
	status, err := o.statusRepo.FindById(statusId)

	if err != nil {
		if errors.Is(err, db.EntityNotFoundErr) {
			return nil, consts.NotFoundErr
		}

		return nil, consts.ServerErr
	}

	order, err := o.orderRepo.ChangeStatus(orderId, status)

	if err != nil {
		if errors.Is(err, db.EntityNotFoundErr) {
			return nil, consts.NotFoundErr
		}

		return nil, consts.ServerErr
	}

	return order.ToDTO(), nil
}
