package models

import "github.com/google/uuid"

// Tournament model info
// @Description Tournament account information
type Tournament struct {
	ID    uuid.UUID `gorm:"primary_key; unique; type:uuid;" swaggerignore:"true"`
	Games []Game    `gorm:"foreignKey:ID;association_foreignkey:ID" swaggerignore:"true"`
} //@name Tournament
