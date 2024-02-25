package services

import (
	"log"

	"github.com/acergiyen/simulator/infra/postgresql/models"
	"github.com/acergiyen/simulator/internal/app/repositories"
)

type MatchService struct {
	lgr  *log.Logger
	repo repositories.MatchRepositoryInterface
}

func NewMatchService(lgr *log.Logger, repo repositories.MatchRepositoryInterface) *MatchService {
	return &MatchService{lgr: lgr, repo: repo}
}

func (m *MatchService) AddMatches(matches []models.Match) error {
	for i := 0; i < len(matches); i++ {

		err := m.repo.AddMatch(&matches[i])
		if err != nil {
			m.lgr.Printf("AddMatches error:%v", err)
			return err
		}
	}

	return nil
}

func (m *MatchService) GetAllMatches() ([]models.Match, error) {
	matches, err := m.repo.GetAllMatches()
	if err != nil {
		m.lgr.Printf("GetAllMatches error:%v", err)
		return nil, err
	}
	return matches, nil
}

func (m *MatchService) GetAllMatchesWithStatistics() ([]models.Match, error) {
	matches, err := m.repo.GetAllMatchesWithStatistics()
	if err != nil {
		m.lgr.Printf("GetAllMatches error:%v", err)
		return nil, err
	}
	return matches, nil
}
