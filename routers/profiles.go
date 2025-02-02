package routers

import (
	"funtastix/backend/controllers"
	"funtastix/backend/middlewares"

	"github.com/gin-gonic/gin"
)

func ProfileRouter(router *gin.RouterGroup) {
	router.Use(middlewares.ValidateToken())
	router.GET("", controllers.GetCurrentProfile)
	router.GET("/:id", controllers.GetProfileById)
	router.PATCH("", controllers.EditProfile)
}
