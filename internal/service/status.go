package service

import (
	"task_cart/internal/model/consts"
	"task_cart/internal/model/dto"
	"task_cart/internal/repository"
)

type StatusInterface interface {
	All() ([]dto.StatusDTO, error)
}

type StatusService struct {
	statusRepo repository.StatusInterface
}

func NewStatusService(statusRepo repository.StatusInterface) *StatusService {
	return &StatusService{statusRepo: statusRepo}
}

func (s *StatusService) All() ([]dto.StatusDTO, error) {
	var statusData []dto.StatusDTO

	statuses, err := s.statusRepo.All()

	if err != nil {
		return nil, consts.ServerErr
	}

	for _, v := range statuses {
		statusData = append(statusData, *v.ToDTO())
	}

	return statusData, nil
}
