package models

import "gorm.io/gorm"

type Score struct {
	gorm.Model
	MatchID  uint `json:"match_id"`
	PlayerID uint `json:"player_id"`
	Score    int  `json:"score"`
}
