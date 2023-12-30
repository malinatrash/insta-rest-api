package main

import (
	"github.com/gin-gonic/gin"
	"github.com/malinatrash/insta-rest-api/config"
	"github.com/malinatrash/insta-rest-api/database"
	"github.com/malinatrash/insta-rest-api/models"
	"github.com/malinatrash/insta-rest-api/routers"
)

func main() {
	db, err := database.ConnectDB(config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_NAME)
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return
	}

	router := gin.Default()
	routers.SetupRoutes(router)
	if err := router.Run("0.0.0.0:8000"); err != nil {
		panic(err)
	}
}
