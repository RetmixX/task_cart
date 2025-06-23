package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"strconv"
	"task_cart/internal/helper"
	"task_cart/internal/model/consts"
	"task_cart/internal/model/dto"
	"task_cart/internal/service"
	"task_cart/pkg/log/sl"
)

type OrderController struct {
	orderService service.OrderInterface
	log          *slog.Logger
}

func NewOrderController(orderService service.OrderInterface, log *slog.Logger) *OrderController {
	return &OrderController{orderService: orderService, log: log}
}

// All Orders godoc
//
//	@Summary	Просмотр всех заказов
//	@Tags		orders
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto.OrderDTO
//	@Router		/orders [get]
func (o *OrderController) GetAll(c *gin.Context) {
	const op = "v1.controllers.order.GetAll"
	log := o.log.With(slog.String("op", op))
	result, err := o.orderService.ViewOrders()

	if err != nil {
		log.Error("fail get orders: ", sl.Err(err))
		helper.ServerErr(c)
		return
	}

	helper.OkResponse(c, result)
}

// Create order godoc
//
//	@Summary	Создание заказа
//	@Tags		orders
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto.OrderDTO
//	@Router		/orders [post]
func (o *OrderController) Create(c *gin.Context) {
	const op = "v1.controllers.order.Create"
	log := o.log.With(slog.String("op", op))
	result, err := o.orderService.CreateOrder()

	if err != nil {
		if errors.Is(err, consts.InvalidRequest) {
			log.Warn("cart is empty")
			helper.BadRequest(c, "Cart is empty")
			return
		}
		log.Error("fail create new order: ", sl.Err(err))
		helper.ServerErr(c)
		return
	}

	helper.CreatedResponse(c, result)
}

// Create order godoc
//
//	@Summary	Изменение статуса заказа
//	@Tags		orders
//	@Accept		json
//	@Produce	json
//	@Param		id				path		int					true	"Ид статуса"
//	@Param		UpdateOrderDTO	body		dto.UpdateOrderDTO	true	"Обновление статуса"
//	@Success	200				{object}	dto.OrderDTO
//	@Router		/orders [put]
func (o *OrderController) Update(c *gin.Context) {
	const op = "v1.controllers.order.Update"
	log := o.log.With(slog.String("op", op))
	id := c.Param("id")
	orderId, err := strconv.Atoi(id)

	if err != nil {
		log.Warn("invalid url param: ", id)
		helper.UrlParamErr(c)
		return
	}

	var body dto.UpdateOrderDTO

	if err = c.ShouldBindJSON(&body); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			log.Warn("validation error")
			helper.ValidationErr(c, validationErrors)
			return
		}
		log.Warn("invalid body")
		helper.BadRequest(c, "invalid body")
		return
	}

	result, err := o.orderService.ChangeStatus(uint(orderId), body.StatusId)

	if err != nil {
		if errors.Is(err, consts.NotFoundErr) {
			log.Warn("not found status, idOrder: ", orderId, "statusId: ", body.StatusId)
			helper.NotFoundResponse(c)
			return
		}
		log.Error("fail change status: ", sl.Err(err))
		helper.ServerErr(c)
		return
	}

	helper.OkResponse(c, result)

}
