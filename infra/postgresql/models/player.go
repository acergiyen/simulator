package models

import "gorm.io/gorm"

type Player struct {
	gorm.Model
	PlayerName string   `json:"player_name"`
	TeamID     uint     `json:"team_id"`
	Scores     []Score  `gorm:"foreignKey:PlayerID"`
	Assists    []Assist `gorm:"foreignKey:PlayerID"`
}
