package utils

import (
	"github.com/malinatrash/insta-rest-api/database"
	"github.com/malinatrash/insta-rest-api/models"
)

func MakeMigrations() {
	var err error
	err = database.DB.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
	err = database.DB.AutoMigrate(&models.Session{})
	if err != nil {
		return
	}
}
