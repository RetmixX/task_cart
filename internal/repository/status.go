package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"task_cart/internal/model/entity"
	"task_cart/pkg/db"
)

type StatusInterface interface {
	FindById(id uint) (*entity.Status, error)
	All() ([]entity.Status, error)
}

type StatusRepository struct {
	db *gorm.DB
}

func NewStatusRepository(conn *gorm.DB) *StatusRepository {
	return &StatusRepository{db: conn}
}

func (s *StatusRepository) FindById(id uint) (*entity.Status, error) {
	const op = "repository.status.FindById"
	var status entity.Status
	if err := s.db.First(&status, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, db.EntityNotFoundErr
		}

		return nil, fmt.Errorf("%s: can't get status by id: %w", op, err)
	}

	return &status, nil
}

func (s *StatusRepository) All() ([]entity.Status, error) {
	const op = "repository.status.All"
	var statuses []entity.Status

	if err := s.db.Find(&statuses).Error; err != nil {
		return nil, fmt.Errorf("%s: can't get all status data: %w", op, err)
	}

	return statuses, nil
}
