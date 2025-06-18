package v1

import (
	"github.com/gin-gonic/gin"
	"task_cart/internal/helper"
	"task_cart/internal/service"
)

type ProductController struct {
	productService service.ProductInterface
}

func NewProductController(productService service.ProductInterface) *ProductController {
	return &ProductController{productService: productService}
}

func (p *ProductController) GetAll(c *gin.Context) {
	result, err := p.productService.All()

	if err != nil {
		helper.ServerErr(c)
		return
	}

	helper.OkResponse(c, result)
}
