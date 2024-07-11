package models

import (
	"github.com/google/uuid"
)

// Game model info
// @Description Game information between two employees
type Game struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid" swaggerignore:"true"`
	Player1ID uuid.UUID `gorm:"type:uuid"`
	Player2ID uuid.UUID `gorm:"type:uuid"`
	Player1   Employee  `gorm:"foreignKey:Player1ID"`
	Player2   Employee  `gorm:"foreignKey:Player2ID"`
	Score1    int
	Score2    int
} //@name Game
