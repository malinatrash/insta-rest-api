package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/malinatrash/insta-rest-api/docs"
	"github.com/malinatrash/insta-rest-api/internal/config"
	"github.com/malinatrash/insta-rest-api/internal/database"
	"github.com/malinatrash/insta-rest-api/internal/routers"
	"github.com/malinatrash/insta-rest-api/internal/utils"
)

// @title           insta REST API
// @version         0.0.1
// @description     Веб сервер для нашего клона инсты
// @host      92.51.45.202:8000
// @BasePath  /

func main() {
	_, err := database.ConnectDB(config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_NAME)
	if err != nil {
		panic("Failed to connect to database!")
	}

	utils.MakeMigrations()

	router := gin.Default()

	config.SetupCORS(router)

	routers.SetupRoutes(router)
	if err := router.Run("0.0.0.0:8000"); err != nil {
		panic(err)
	}
}
