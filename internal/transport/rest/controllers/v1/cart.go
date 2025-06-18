package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"task_cart/internal/helper"
	"task_cart/internal/model/consts"
	"task_cart/internal/model/dto"
	"task_cart/internal/service"
)

type CartController struct {
	cartService service.CartInterface
}

func NewCartController(cartService service.CartInterface) *CartController {
	return &CartController{cartService: cartService}
}

func (s *CartController) All(c *gin.Context) {
	result, err := s.cartService.SeeCart()

	if err != nil {
		helper.ServerErr(c)
		return
	}

	helper.OkResponse(c, result)
}

func (s *CartController) Add(c *gin.Context) {
	var body dto.AddProductCartDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		helper.ValidationErr(c, err)
		return
	}

	result, err := s.cartService.AddProduct(body.ProductId, body.Quantity)

	if err != nil {
		if errors.Is(err, consts.NotFoundErr) {
			helper.NotFoundResponse(c)
			return
		}

		helper.ServerErr(c)
		return
	}

	helper.CreatedResponse(c, result)
}

func (s *CartController) DeleteProduct(c *gin.Context) {
	var body dto.DeleteProductDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		helper.ValidationErr(c, err)
		return
	}

	result, err := s.cartService.DeleteProduct(body.ProductId)

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
