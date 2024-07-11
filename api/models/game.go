package models

import (
	"github.com/google/uuid"
)

// Game model info
// @Description Game information between an employee and an user
type Game struct {
	ID           uuid.UUID `gorm:"primary_key;unique;type:uuid;" swaggerignore:"true"`
	Player1ID    uuid.UUID
	Player2ID    uuid.UUID
	Player1Score Score `gorm:"foreignKey:MatchID,PlayerID;references:ID,Player1ID" `
	Player2Score Score `gorm:"foreignKey:MatchID,PlayerID;references:ID,Player2ID" `
} //@name Game

// Score model info
// @Description Score information between an employee and an user
type Score struct {
	PlayerID uuid.UUID `gorm:"primary_key" swaggerignore:"true"`
	MatchID  uuid.UUID `gorm:"primary_key" swaggerignore:"true"`
	Score    int
} //@name Score
