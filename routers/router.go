package routers

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	SetupSwaggerRouter(r)
	SetupUserRoutes(r)
	SetupImageRoutes(r)
	SetupSessionRoutes(r)
}
