package routers

import (
	"funtastix/backend/controllers"

	"github.com/gin-gonic/gin"
)

func ProfileRouter(router *gin.RouterGroup) {
	// router.Use(middlewares.ValidateToken())
	router.GET("", controllers.GetAllProfiles)
	router.GET("/:id", controllers.GetProfileById)
	router.PATCH("/:id", controllers.EditProfile)
}
