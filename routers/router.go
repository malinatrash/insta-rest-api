package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	SetupSwaggerRoute(r)
	SetupUserRoutes(r)
	SetupImageRoutes(r)
}
