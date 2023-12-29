package main

import (
	"github.com/gin-gonic/gin"
	"github.com/malinatrash/insta-rest-api/routers"
)

func main() {
	r := gin.Default()

	// Инициализация роутера
	routers.SetupRoutes(r)

	if err := r.Run(":8000"); err != nil {
		panic(err)
	}
}
