package service

import (
	"errors"
	"log/slog"
	"task_cart/internal/model/consts"
	"task_cart/internal/model/dto"
	"task_cart/internal/repository"
	"task_cart/pkg/db"
	"task_cart/pkg/log/sl"
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
	logger     *slog.Logger
}

func NewOrderService(orderRepo repository.OrderInterface,
	statusRepo repository.StatusInterface,
	cartRepo repository.CartInterface, log *slog.Logger) *OrderService {
	return &OrderService{
		orderRepo:  orderRepo,
		statusRepo: statusRepo,
		cartRepo:   cartRepo,
		logger:     log,
	}
}

func (o *OrderService) ViewOrders() ([]dto.OrderDTO, error) {
	const op = "service.order.ViewOrders"
	log := o.logger.With(slog.String("op", op))
	log.Info("Start ViewOrders()")
	var ordersData []dto.OrderDTO
	orders, err := o.orderRepo.ViewOrders()

	if err != nil {
		log.Error("fail get orders ", sl.Err(err))
		return nil, consts.ServerErr
	}

	for _, v := range orders {
		ordersData = append(ordersData, *v.ToDTO())
	}
	log.Info("End ViewOrders()")
	return ordersData, nil

}

func (o *OrderService) CreateOrder() (*dto.OrderDTO, error) {
	const op = "service.order.CreateOrder"
	log := o.logger.With(slog.String("op", op))
	log.Info("Start CreateOrder()")
	cart, err := o.cartRepo.ViewCart()
	if err != nil {
		log.Error("fail get current cart ", sl.Err(err))
		return nil, consts.ServerErr
	}

	if len(cart.Products) == 0 {
		log.Warn("cart is empty")
		return nil, consts.InvalidRequest
	}

	newOrder, err := o.orderRepo.CreateOrder()
	if err != nil {
		log.Error("fail create order ", sl.Err(err))
		return nil, consts.ServerErr
	}
	log.Info("End CreateOrder()")
	return newOrder.ToDTO(), nil
}

func (o *OrderService) ChangeStatus(orderId uint, statusId uint) (*dto.OrderDTO, error) {
	const op = "service.order.ChangeStatus"
	log := o.logger.With(slog.String("op", op))
	log.Info("Start ChangeStatus()")

	status, err := o.statusRepo.FindById(statusId)
	if err != nil {
		if errors.Is(err, db.EntityNotFoundErr) {
			log.Warn("not found status by id: ", status)
			return nil, consts.NotFoundErr
		}
		log.Error("fail get status: ", sl.Err(err))
		return nil, consts.ServerErr
	}

	order, err := o.orderRepo.ChangeStatus(orderId, status)

	if err != nil {
		if errors.Is(err, db.EntityNotFoundErr) {
			log.Warn("not found order by id: ", orderId)
			return nil, consts.NotFoundErr
		}
		log.Error("fail get order", sl.Err(err))
		return nil, consts.ServerErr
	}
	log.Info("End ChangeStatus()")
	return order.ToDTO(), nil
}
