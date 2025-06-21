package v1

import (
	"github.com/gin-gonic/gin"
	"task_cart/internal/helper"
	"task_cart/internal/service"
)

type StatusController struct {
	statusService service.StatusInterface
}

func NewStatusController(statusService service.StatusInterface) *StatusController {
	return &StatusController{statusService: statusService}
}

// All Statuses godoc
//
//	@Summary	Просмотр всех статусов
//	@Tags		status
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto.StatusDTO
//	@Router		/statuses [get]
func (s *StatusController) GetAll(c *gin.Context) {
	result, err := s.statusService.All()

	if err != nil {
		helper.ServerErr(c)
		return
	}

	helper.OkResponse(c, result)
}
