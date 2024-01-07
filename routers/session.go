package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/malinatrash/insta-rest-api/controllers"
)

func SetupSessionRoutes(r *gin.Engine) {
	r.GET("/session", controllers.GetSessionByUsernameAndPassword)
}
