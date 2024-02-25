package services

import (
	"log"

	"github.com/acergiyen/simulator/infra/postgresql/models"
	"github.com/acergiyen/simulator/internal/app/repositories"
)

type AssistService struct {
	lgr  *log.Logger
	repo repositories.AssistRepositoryInterface
}

func NewAssistService(lgr *log.Logger, repo repositories.AssistRepositoryInterface) *AssistService {
	return &AssistService{lgr: lgr, repo: repo}
}

func (s *AssistService) AddAssist(assist *models.Assist) error {
	err := s.repo.AddAssist(assist)
	if err != nil {
		s.lgr.Fatalf("AddScores error:%v", err)
		return err
	}
	return nil
}
