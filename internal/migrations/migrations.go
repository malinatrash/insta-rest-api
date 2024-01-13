package migrations

import (
	"fmt"
	"github.com/malinatrash/insta-rest-api/internal/database"
	"github.com/malinatrash/insta-rest-api/internal/models"
)

func MakeMigrations() {
	if err := migrateModels(&models.User{}); err != nil {
		fmt.Printf("Error migrating User model: %v\n", err)
		return
	}

	if err := migrateModels(&models.Session{}); err != nil {
		fmt.Printf("Error migrating Session model: %v\n", err)
		return
	}

	fmt.Println("Migrations completed successfully")
}

func migrateModels(models ...interface{}) error {
	return database.DB.AutoMigrate(models...)
}
