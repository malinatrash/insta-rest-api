package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/malinatrash/insta-rest-api/controllers"
)

func SetupUserRoutes(r *gin.Engine) {
	userRoutes := r.Group("/users")

	userRoutes.GET("/", controllers.GetAllUsers)
	userRoutes.GET("/login", controllers.GetUserByUsernameAndPassword)
	userRoutes.GET("/:id", controllers.GetUserByID)
	userRoutes.POST("/", controllers.CreateUser)
	userRoutes.PUT("/:id", controllers.UpdateUser)
	userRoutes.DELETE("/:id", controllers.DeleteUser)
	userRoutes.POST("/caption/:id", controllers.AddCaption)

}
