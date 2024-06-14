package config

import (
	"fmt"
	"os"

	"github.com/MagnusV9/tietoevry-pong-mmr/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := os.Getenv("DATABASE_SECRET")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	db.AutoMigrate(&models.Employee{}, &models.Tournament{}, &models.Game{})

	DB = db
}
