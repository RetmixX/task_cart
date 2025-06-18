package v1

import (
	"github.com/gin-gonic/gin"
	v1 "task_cart/internal/transport/rest/controllers/v1"
)

var productURL = "/v1/products"

func RegisterProductRoutes(
	engine *gin.Engine, productCtrl *v1.ProductController) {

	productGroup := engine.Group(productURL)

	productGroup.GET("", productCtrl.GetAll)
}
