package v1

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"task_cart/internal/helper"
	"task_cart/internal/service"
	"task_cart/pkg/log/sl"
)

type ProductController struct {
	productService service.ProductInterface
	log            *slog.Logger
}

func NewProductController(productService service.ProductInterface, log *slog.Logger) *ProductController {
	return &ProductController{productService: productService, log: log}
}

// All Products godoc
//
//	@Summary	Просмотр всех товаров
//	@Tags		products
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto.ProductWithCountDTO
//	@Router		/products [get]
func (p *ProductController) GetAll(c *gin.Context) {
	const op = "v1.controllers.product.GetAll"
	log := p.log.With(slog.String("op", op))
	result, err := p.productService.All()

	if err != nil {
		log.Error("fail get products: ", sl.Err(err))
		helper.ServerErr(c)
		return
	}

	helper.OkResponse(c, result)
}
