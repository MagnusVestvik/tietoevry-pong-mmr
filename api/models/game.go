package models

import (
	"github.com/google/uuid"
)

// Game model info
// @Description Game information between an employee and an user
type Game struct {
	ID           uuid.UUID   `gorm:"primary_key;unique;type:uuid;" swaggerignore:"true" json:"id"`
	Employee1ID  uuid.UUID   `gorm:"type:uuid;not null" json:"employee1_id"`
	Employee2ID  uuid.UUID   `gorm:"type:uuid;not null" json:"employee2_id"`
	Employee1    PlayerScore `gorm:"foreignKey:Employee1ID" json:"employee1"`
	Employee2    PlayerScore `gorm:"foreignKey:Employee2ID" json:"employee2"`
	TournamentID *uuid.UUID  `json:"tournament_id"`
}

//@name Game
