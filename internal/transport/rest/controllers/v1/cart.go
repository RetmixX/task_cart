package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
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

// All View cart godoc
//
//	@Summary	Просмотр товаров в корзине
//	@Tags		carts
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto.CartDTO
//	@Router		/api/v1/cart [get]
func (s *CartController) All(c *gin.Context) {
	result, err := s.cartService.SeeCart()

	if err != nil {
		helper.ServerErr(c)
		return
	}

	helper.OkResponse(c, result)
}

// Add product to cart
//
//	@Summary	Добавление товара в корзину
//	@Tags		carts
//	@Param		AddProductCartDTO	body	dto.AddProductCartDTO	true	"Добавить товар в козрину"
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto.CartDTO
//	@Router		/api/v1/cart [post]
func (s *CartController) Add(c *gin.Context) {
	var body dto.AddProductCartDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		var validError validator.ValidationErrors
		if errors.As(err, &validError) {
			helper.ValidationErr(c, validError)
			return
		}
		helper.BadRequest(c, "invalid body")
		return
	}

	result, err := s.cartService.AddProduct(body.ProductId, body.Quantity)

	if err != nil {
		if errors.Is(err, consts.NotFoundErr) {
			helper.NotFoundResponse(c)
			return
		}

		if errors.Is(err, consts.InvalidRequest) {
			helper.BadRequest(c, "invalid count quantity")
			return
		}

		helper.ServerErr(c)
		return
	}

	helper.CreatedResponse(c, result)
}

// Delete product to cart
//
//	@Summary	Удаление товара из корзины
//	@Tags		carts
//	@Param		id	path	int	true	"Ид товара"
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto.CartDTO
//	@Router		/api/v1/cart/product/{id} [delete]
func (s *CartController) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	idProduct, err := strconv.Atoi(id)
	if err != nil {
		helper.UrlParamErr(c)
		return
	}
	result, err := s.cartService.DeleteProduct(uint(idProduct))

	if err != nil {
		if errors.Is(err, consts.NotFoundErr) {
			helper.NotFoundResponse(c)
			return
		}
		fmt.Println(err)
		helper.ServerErr(c)
		return
	}

	helper.OkResponse(c, result)
}
