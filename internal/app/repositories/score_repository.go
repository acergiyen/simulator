package repositories

import (
	"fmt"

	"github.com/acergiyen/simulator/infra/postgresql/models"
	"gorm.io/gorm"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . ScoreRepositoryInterface

type ScoreRepositoryInterface interface {
	AddScore(score *models.Score) error
}

type ScoreRepository struct {
	db *gorm.DB
}

func NewScoreRepository(db *gorm.DB) *ScoreRepository {
	return &ScoreRepository{db: db}
}

func (s *ScoreRepository) AddScore(score *models.Score) error {
	err := s.db.Create(score).Error
	if err != nil {
		return fmt.Errorf("database error:%v", err)
	}
	return nil
}
