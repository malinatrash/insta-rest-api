package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/malinatrash/insta-rest-api/internal/controllers"
)

func SetupImageRoutes(r *gin.Engine) {
	imageRoutes := r.Group("/image")
	imageRoutes.GET("/", controllers.GetImageForUser)
	imageRoutes.POST("/", controllers.UploadImageForUser)
}
