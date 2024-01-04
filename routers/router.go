package routers

import (
	"github.com/gin-gonic/gin"
	docs "github.com/malinatrash/insta-rest-api/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	SetupUserRoutes(r)
	SetupImageRoutes(r)

}
