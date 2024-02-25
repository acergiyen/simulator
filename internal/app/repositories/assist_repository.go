package repositories

import (
	"fmt"

	"github.com/acergiyen/simulator/infra/postgresql/models"
	"gorm.io/gorm"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . AssistRepositoryInterface

type AssistRepositoryInterface interface {
	AddAssist(assist *models.Assist) error
}

type AssistRepository struct {
	db *gorm.DB
}

func NewAssistRepository(db *gorm.DB) *AssistRepository {
	return &AssistRepository{db: db}
}

func (s *AssistRepository) AddAssist(assist *models.Assist) error {
	err := s.db.Create(assist).Error
	if err != nil {
		return fmt.Errorf("database error:%v", err)
	}
	return nil
}
