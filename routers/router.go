package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tharinduyasantha/go-postgresql-rest-api/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("/", controllers.CreateUser)
		userRoutes.GET("all", controllers.GetUsers)
		userRoutes.GET("/:id", controllers.GetUser)
		userRoutes.DELETE("/:id", controllers.DeleteUser)
		userRoutes.PUT("/:id", controllers.UpdateUser)
	}
	return r
}
