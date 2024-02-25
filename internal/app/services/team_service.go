package services

import (
	"log"

	"github.com/acergiyen/simulator/infra/postgresql/models"
	"github.com/acergiyen/simulator/internal/app/repositories"
)

type TeamService struct {
	lgr  *log.Logger
	repo repositories.TeamRepositoryInterface
}

func NewTeamService(lgr *log.Logger, repo repositories.TeamRepositoryInterface) *TeamService {
	return &TeamService{lgr: lgr, repo: repo}
}

func (t *TeamService) GetAllTeams() ([]models.Team, error) {
	teams, err := t.repo.GetAllTeams()
	if err != nil {
		t.lgr.Printf("GetAllTeams error :%v", err)
		return nil, err
	}
	return teams, nil

}
