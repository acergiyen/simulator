package models

import "gorm.io/gorm"

type Team struct {
	gorm.Model
	TeamName string   `json:"team_name"`
	Players  []Player `gorm:"foreignKey:TeamID"`
}
