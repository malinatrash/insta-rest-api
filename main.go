package main

import (
	"github.com/gin-gonic/gin"
	"github.com/malinatrash/insta-rest-api/config"
	"github.com/malinatrash/insta-rest-api/database"
	_ "github.com/malinatrash/insta-rest-api/docs"
	"github.com/malinatrash/insta-rest-api/models"
	"github.com/malinatrash/insta-rest-api/routers"
)

// @title           Swagger Example API
// @version         1.0
// @description     Insta rest api
// @termsOfService  http://swagger.io/terms/

// @host      http://92.51.45.202:8000/
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
