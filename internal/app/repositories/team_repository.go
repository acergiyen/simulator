package repositories

import (
	"fmt"

	"github.com/acergiyen/simulator/infra/postgresql/models"
	"gorm.io/gorm"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . TeamRepositoryInterface

type TeamRepositoryInterface interface {
	GetAllTeams() ([]models.Team, error)
}

type TeamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) *TeamRepository {
	return &TeamRepository{db: db}
}

func (t *TeamRepository) GetAllTeams() ([]models.Team, error) {
	var teams []models.Team
	err := t.db.Preload("Players").Find(&teams).Error
	if err != nil {
		return nil, fmt.Errorf("database error::%v", err)
	}
	return teams, nil
}
