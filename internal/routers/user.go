package routers

import (
	"github.com/gin-gonic/gin"
	controllers2 "github.com/malinatrash/insta-rest-api/internal/controllers"
)

func SetupUserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/users")
	userRoutes.GET("/", controllers2.GetAllUsers)
	userRoutes.GET("/login", controllers2.GetUserBySession)
	userRoutes.GET("/:id", controllers2.GetUserByID)
	userRoutes.POST("/", controllers2.CreateUser)
	userRoutes.PUT("/:id", controllers2.UpdateUser)
	userRoutes.DELETE("/:id", controllers2.DeleteUser)
	userRoutes.POST("/caption/:id", controllers2.AddCaption)

}
