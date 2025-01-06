package routers

import (
	"funtastix/backend/controllers"
	"funtastix/backend/middlewares"

	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.RouterGroup) {
	router.Use(middlewares.RateLimiter)
	router.POST("/login", controllers.Login)
	router.POST("/register", controllers.Register)
}
