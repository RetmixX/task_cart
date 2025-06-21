package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
	"task_cart/internal/helper"
	"task_cart/internal/model/consts"
	"task_cart/internal/model/dto"
	"task_cart/internal/service"
)

type OrderController struct {
	orderService service.OrderInterface
}

func NewOrderController(orderService service.OrderInterface) *OrderController {
	return &OrderController{orderService: orderService}
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
	result, err := o.orderService.ViewOrders()

	if err != nil {
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
	result, err := o.orderService.CreateOrder()

	if err != nil {
		if errors.Is(err, consts.InvalidRequest) {
			helper.BadRequest(c, "Cart is empty")
			return
		}

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
	id := c.Param("id")
	orderId, err := strconv.Atoi(id)

	if err != nil {
		helper.UrlParamErr(c)
		return
	}

	var body dto.UpdateOrderDTO

	if err = c.ShouldBindJSON(&body); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) {
			helper.ValidationErr(c, validationErrors)
			return
		}
		helper.BadRequest(c, "invalid body")
		return
	}

	result, err := o.orderService.ChangeStatus(uint(orderId), body.StatusId)

	if err != nil {
		if errors.Is(err, consts.NotFoundErr) {
			helper.NotFoundResponse(c)
			return
		}

		helper.ServerErr(c)
		return
	}

	helper.OkResponse(c, result)

}
