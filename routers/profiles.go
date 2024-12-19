package routers

import (
	"funtastix/backend/middlewares"

	"github.com/gin-gonic/gin"
)

func ProfileRouter(router *gin.RouterGroup) {
	router.Use(middlewares.ValidateToken())

}
