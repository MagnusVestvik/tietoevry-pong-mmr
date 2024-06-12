package models

import "github.com/google/uuid"

// Employee model info
// @Description Employee account information
type Employee struct {
	ID         uuid.UUID `gorm:"primary_key; unique; type:uuid;" swaggerignore:"true"`
	Name       string    `validate:"required"`
	Email      string    `validate:"required"`
	Password   string    `validate:"required"`
	Elo        int       `validate:"required"`
	Department string    `validate:"required"`
	Games      []Game    `gorm:"foreignKey:Employee1ID;association_foreignkey:ID" swaggerignore:"true"` // Assuming Employee1ID is the foreign key
} //@name Employee

// UpdateEmployee model info
// @Description UpdateEmployee account information
type UpdateEmployee struct {
	Name       *string
	Email      *string
	Password   *string
	Elo        *int
	Department *string
} //@name UpdateEmployee
