package main

import (
	"funtastix/backend/docs"
	"funtastix/backend/routers"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Funtastix
// @version         1.0.0
// @description     Funtastix backend-app.
// @host      			172.16.211.131:8888
// @BasePath  /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	route := gin.Default()

	// route.Use(middlewares.RateLimiter)
	route.Static("/profile/images", "uploads/profile")
	route.MaxMultipartMemory = 2 << 20
	docs.SwaggerInfo.BasePath = "/"
	route.GET("/docs/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	routers.Routers(route)
	// libs.GracefulShutdown(":8888", route)
	route.Run(":8888")
}
