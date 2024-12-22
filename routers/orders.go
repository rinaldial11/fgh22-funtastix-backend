package routers

import (
	"funtastix/backend/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRouter(router *gin.RouterGroup) {
	router.GET("", controllers.GetAllOrders)
	router.POST("", controllers.AddOrder)
}
