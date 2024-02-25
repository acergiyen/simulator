package repositories

import (
	"fmt"

	"github.com/acergiyen/simulator/infra/postgresql/models"
	"gorm.io/gorm"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . MatchRepositoryInterface
type MatchRepositoryInterface interface {
	AddMatch(match *models.Match) error
	GetAllMatches() ([]models.Match, error)
	GetAllMatchesWithStatistics() ([]models.Match, error)
}

type MatchRepository struct {
	db *gorm.DB
}

func NewMatchRepository(db *gorm.DB) *MatchRepository {
	return &MatchRepository{db: db}
}

func (m *MatchRepository) AddMatch(match *models.Match) error {
	err := m.db.Create(match).Error
	if err != nil {
		return fmt.Errorf("database error:%v", err)
	}
	return nil
}

func (m *MatchRepository) GetAllMatches() ([]models.Match, error) {
	var matches []models.Match
	err := m.db.Preload("HomeTeam.Players").Preload("AwayTeam.Players").Find(&matches).Error
	if err != nil {
		fmt.Println("Error during Preload:", err)
		return nil, fmt.Errorf("database error::%v", err)
	}
	return matches, nil
}

func (m *MatchRepository) GetAllMatchesWithStatistics() ([]models.Match, error) {
	var matches []models.Match
	err := m.db.Preload("Scores").Preload("Assists").Find(&matches).Error
	if err != nil {
		fmt.Println("Error during Preload:", err)
		return nil, fmt.Errorf("database error::%v", err)
	}
	return matches, nil
}
