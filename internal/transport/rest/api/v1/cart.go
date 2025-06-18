package v1

import (
	"github.com/gin-gonic/gin"
	v1 "task_cart/internal/transport/rest/controllers/v1"
)

var cartURL = "/v1/carts"

func RegisterCartRoutes(
	engine *gin.Engine, cartCtrl *v1.CartController) {

	cartGroup := engine.Group(cartURL)

	cartGroup.GET("", cartCtrl.All)
	cartGroup.DELETE(":id", cartCtrl.DeleteProduct)
	cartGroup.POST("", cartCtrl.Add)
}
