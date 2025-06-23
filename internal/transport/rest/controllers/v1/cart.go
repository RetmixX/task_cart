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

type CartController struct {
	cartService service.CartInterface
	log         *slog.Logger
}

func NewCartController(cartService service.CartInterface, log *slog.Logger) *CartController {
	return &CartController{cartService: cartService, log: log}
}

// All View cart godoc
//
//	@Summary	Просмотр товаров в корзине
//	@Tags		carts
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto.CartDTO
//	@Router		/cart [get]
func (s *CartController) All(c *gin.Context) {
	const op = "v1.controllers.cart.All"
	log := s.log.With(slog.String("op", op))
	result, err := s.cartService.SeeCart()

	if err != nil {
		log.Error("fail get cart data ", sl.Err(err))
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
//	@Router		/cart [post]
func (s *CartController) Add(c *gin.Context) {
	const op = "v1.controllers.cart.Add"
	log := s.log.With(slog.String("op", op))
	var body dto.AddProductCartDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		var validError validator.ValidationErrors
		if errors.As(err, &validError) {
			log.Warn("validation error")
			helper.ValidationErr(c, validError)
			return
		}
		log.Warn("sended json invalid")
		helper.BadRequest(c, "invalid body")
		return
	}

	result, err := s.cartService.AddProduct(body.ProductId, body.Quantity)

	if err != nil {
		if errors.Is(err, consts.NotFoundErr) {
			log.Warn("product not found by id: ", body.ProductId)
			helper.NotFoundResponse(c)
			return
		}

		if errors.Is(err, consts.InvalidRequest) {
			log.Warn("invalid count property: ", body.Quantity)
			helper.BadRequest(c, "invalid count quantity")
			return
		}
		log.Error("fail add product to cart, error: ", sl.Err(err))
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
//	@Router		/cart/product/{id} [delete]
func (s *CartController) DeleteProduct(c *gin.Context) {
	const op = "v1.controllers.cart.DeleteProduct"
	log := s.log.With(slog.String("op", op))

	id := c.Param("id")
	idProduct, err := strconv.Atoi(id)
	if err != nil {
		log.Warn("invalid url param: ", id)
		helper.UrlParamErr(c)
		return
	}
	result, err := s.cartService.DeleteProduct(uint(idProduct))

	if err != nil {
		if errors.Is(err, consts.NotFoundErr) {
			log.Warn("not found product with id: ", idProduct)
			helper.NotFoundResponse(c)
			return
		}
		log.Error("fail delete product from cart: ", idProduct, sl.Err(err))
		helper.ServerErr(c)
		return
	}

	helper.OkResponse(c, result)
}
