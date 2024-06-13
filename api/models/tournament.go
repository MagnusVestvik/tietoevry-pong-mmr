package models

import "github.com/google/uuid"

// Tournament model info
// @Description Tournament account information
type Tournament struct {
	ID    uuid.UUID `gorm:"primary_key; unique; type:uuid;" swaggerignore:"true"`
	Games []Game    `gorm:"foreignKey:Tournament1ID;association_foreignkey:ID" swaggerignore:"true"` // Assuming Tournament1ID is the foreign key
} //@name Tournament
