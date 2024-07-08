package models

import (
	"github.com/google/uuid"
)

// Game model info
// @Description Game information between an employee and an user
type Game struct {
	ID           uuid.UUID  `gorm:"primary_key;unique;type:uuid;" swaggerignore:"true" json:"id"`
	Player1ID    uuid.UUID  `json:"p1_id"`
	Player2ID    uuid.UUID  `json:"p2_id"`
	Player1Score Score      `gorm:"foreignKey:MatchID,PlayerID;references:ID,Player1ID" json:"p1_score"`
	Player2Score Score      `gorm:"foreignKey:MatchID,PlayerID;references:ID,Player2ID" json:"p2_score"`
	TournamentID *uuid.UUID `json:"tournament_id"`
}

//@name Game

// Score model info
// @Description Score information between an employee and an user
type Score struct {
	PlayerID uuid.UUID `gorm:"primary_key" json:"player_id"`
	MatchID  uuid.UUID `gorm:"primary_key" json:"match_id"`
	Score    int       `json:"score"`
}

//@name Score
