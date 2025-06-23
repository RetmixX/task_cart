package service

import (
	"log/slog"
	"task_cart/internal/model/consts"
	"task_cart/internal/model/dto"
	"task_cart/internal/repository"
	"task_cart/pkg/log/sl"
)

type StatusInterface interface {
	All() ([]dto.StatusDTO, error)
}

type StatusService struct {
	statusRepo repository.StatusInterface
	logger     *slog.Logger
}

func NewStatusService(statusRepo repository.StatusInterface, log *slog.Logger) *StatusService {
	return &StatusService{statusRepo: statusRepo, logger: log}
}

func (s *StatusService) All() ([]dto.StatusDTO, error) {
	const op = "service.status.All"
	log := s.logger.With(slog.String("op", op))
	log.Info("Start All()")
	var statusData []dto.StatusDTO

	statuses, err := s.statusRepo.All()

	if err != nil {
		log.Error("fail get statuses ", sl.Err(err))
		return nil, consts.ServerErr
	}

	for _, v := range statuses {
		statusData = append(statusData, *v.ToDTO())
	}
	log.Info("End All()")
	return statusData, nil
}
