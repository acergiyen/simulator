package models

import (
	"time"

	"gorm.io/gorm"
)

type Match struct {
	gorm.Model
	HomeTeamID  uint          `json:"home_team_id"`
	AwayTeamID  uint          `json:"away_team_id"`
	HomeTeam    Team          `gorm:"foreignKey:HomeTeamID"`
	AwayTeam    Team          `gorm:"foreignKey:AwayTeamID"`
	StartTime   time.Time     `json:"start_time"`
	Duration    time.Duration `json:"duration"`
	CurrentTime time.Duration
	Scores      []Score  `gorm:"foreignKey:MatchID"`
	Assists     []Assist `gorm:"foreignKey:MatchID"`
}
