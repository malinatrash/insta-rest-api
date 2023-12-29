package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/malinatrash/insta-rest-api/handlers"
)

func SetupRoutes(r *gin.Engine) {
	// Пример маршрута для пользователя
	r.GET("/user", handlers.GetUser)
	// ... другие маршруты ...
}
