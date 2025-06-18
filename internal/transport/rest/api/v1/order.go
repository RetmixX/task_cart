package v1

import (
	"github.com/gin-gonic/gin"
	v1 "task_cart/internal/transport/rest/controllers/v1"
)

var orderURL = "/v1/orders"

func RegisterOrderRoutes(
	engine *gin.Engine, orderCtrl *v1.OrderController) {

	orderGroup := engine.Group(orderURL)

	orderGroup.GET("all", orderCtrl.GetAll)
	orderGroup.PUT(":id", orderCtrl.Update)
	orderGroup.POST("", orderCtrl.Create)
}
