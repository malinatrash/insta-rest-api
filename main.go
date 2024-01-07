package main

import (
	"github.com/gin-gonic/gin"
	"github.com/malinatrash/insta-rest-api/config"
	"github.com/malinatrash/insta-rest-api/database"
	_ "github.com/malinatrash/insta-rest-api/docs"
	"github.com/malinatrash/insta-rest-api/models"
	"github.com/malinatrash/insta-rest-api/routers"
)

// @title           insta REST API
// @version         0.0.1
// @description     Веб сервер для нашего клона инсты
// @host      92.51.45.202:8000
// @BasePath  /

func main() {
	db, err := database.ConnectDB(config.DB_HOST, config.DB_USER, config.DB_PASSWORD, config.DB_NAME)
	if err != nil {
		panic("Failed to connect to database!")
	}

	err = db.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&models.Session{})
	if err != nil {
		return
	}

	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		}

		c.Next()
	})

	routers.SetupRoutes(router)
	if err := router.Run("0.0.0.0:8000"); err != nil {
		panic(err)
	}
}
