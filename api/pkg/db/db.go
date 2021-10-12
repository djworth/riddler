package db

import (
	"github.com/djworth/riddler/pkg/db/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open("riddler.db"), &gorm.Config{})
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.Riddle{}, &models.AssignedRiddle{})
}
