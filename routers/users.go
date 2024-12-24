package routers

import (
	"funtastix/backend/controllers"
	"funtastix/backend/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {
	router.Use(middlewares.ValidateToken())
	router.GET("", controllers.GetCurrentUser)
	router.GET("/:id", controllers.GetUserById)
	router.PATCH("", controllers.UpdateUser)
	router.DELETE("/:id", controllers.DeleteUser)
	router.POST("", controllers.CreateUser)
}
