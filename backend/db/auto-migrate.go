package db

import (
	"github.com/Lydoww/react-native-go-fastlane/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error {
	return db.AutoMigrate(&models.Event{}, &models.Ticket{}, &models.User{})
}