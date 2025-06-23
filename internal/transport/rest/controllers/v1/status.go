package v1

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"task_cart/internal/helper"
	"task_cart/internal/service"
	"task_cart/pkg/log/sl"
)

type StatusController struct {
	statusService service.StatusInterface
	log           *slog.Logger
}

func NewStatusController(statusService service.StatusInterface, log *slog.Logger) *StatusController {
	return &StatusController{statusService: statusService, log: log}
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
	const op = "v1.controllers.status.GetAll"
	log := s.log.With(slog.String("op", op))
	result, err := s.statusService.All()

	if err != nil {
		log.Error("fail get statuses: ", sl.Err(err))
		helper.ServerErr(c)
		return
	}

	helper.OkResponse(c, result)
}
