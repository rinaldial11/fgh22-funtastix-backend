package routers

import (
	"funtastix/backend/controllers"
	"funtastix/backend/middlewares"

	"github.com/gin-gonic/gin"
)

func OrderRouter(router *gin.RouterGroup) {
	router.Use(middlewares.RateLimiter)
	router.Use(middlewares.ValidateToken())
	router.GET("", controllers.GetAllOrders)
	router.POST("", controllers.AddOrder)
	router.GET("/payment-methods", controllers.GetAllPaymentMethods)
	router.GET("/seats", controllers.GetAllSeats)
}
