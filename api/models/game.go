package models

import (
	"github.com/google/uuid"
)

// Game model info
// @Description Game information between an employee and an user
type Game struct {
	ID             uuid.UUID `gorm:"primary_key; unique; type:uuid;" swaggerignore:"true"`
	Employee1Score int       `validate:"required"`
	Employee2Score int       `validate:"required"`

	Employees    []Employee `gorm:"many2many:employee_games;"`
	TournamentID uuid.UUID
} //@name Game
