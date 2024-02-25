package services

import (
	"log"

	"github.com/acergiyen/simulator/infra/postgresql/models"
	"github.com/acergiyen/simulator/internal/app/repositories"
)

type ScoreService struct {
	lgr  *log.Logger
	repo repositories.ScoreRepositoryInterface
}

func NewScoreService(lgr *log.Logger, repo repositories.ScoreRepositoryInterface) *ScoreService {
	return &ScoreService{lgr: lgr, repo: repo}
}

func (s *ScoreService) AddScore(score *models.Score) error {
	err := s.repo.AddScore(score)
	if err != nil {
		s.lgr.Printf("AddScores error:%v", err)
		return err
	}
	return nil
}
