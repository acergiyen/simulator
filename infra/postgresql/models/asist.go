package models

import "gorm.io/gorm"

type Assist struct {
	gorm.Model
	MatchID     uint `json:"match_id"`
	PlayerID    uint `json:"player_id"`
	AssistCount int  `json:"assist_count"`
}
