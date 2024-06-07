package models

import (
	"github.com/google/uuid"
)

// Game model info
// @Description Game information between an employee and an user
type Game struct {
	ID             uuid.UUID `gorm:"primary_key; unique; type:uuid;" swaggerignore:"true"`
	Employee1ID    uuid.UUID `validate:"required"`
	Employee2ID    uuid.UUID `validate:"required"`
	Employee1Score int       `validate:"required"`
	Employee2Score int       `validate:"required"`

	Employee1 Employee `gorm:"foreignKey:Employee1ID"`
	Employee2 Employee `gorm:"foreignKey:Employee2ID"`
} //@name Game
