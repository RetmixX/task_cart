package v1

import (
	"github.com/gin-gonic/gin"
	v1 "task_cart/internal/transport/rest/controllers/v1"
)

var statusURL = "/api/v1/statuses"

func RegisterStatusRoutes(
	engine *gin.Engine, statusCtrl *v1.StatusController) {

	statusGroup := engine.Group(statusURL)

	statusGroup.GET("", statusCtrl.GetAll)
}
