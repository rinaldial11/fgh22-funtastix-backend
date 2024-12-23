package routers

import (
	"funtastix/backend/controllers"
	"funtastix/backend/middlewares"

	"github.com/gin-gonic/gin"
)

func OrderRouter(router *gin.RouterGroup) {
	router.Use(middlewares.ValidateToken())
	router.GET("", controllers.GetAllOrders)
	router.POST("", controllers.AddOrder)
}
