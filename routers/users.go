package routers

import (
	"funtastix/backend/controllers"
	"funtastix/backend/middlewares"

	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {
	router.Use(middlewares.ValidateToken())
	router.GET("", controllers.GetAllUsers)
	router.GET("/:id", controllers.GetUserById)
	router.PATCH("/:id", controllers.UpdateUser)
	router.DELETE("/:id", controllers.DeleteUser)
	router.POST("", controllers.CreateUser)
}
