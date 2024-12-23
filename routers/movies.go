package routers

import (
	"funtastix/backend/controllers"
	"funtastix/backend/middlewares"

	"github.com/gin-gonic/gin"
)

func MovieRouter(router *gin.RouterGroup) {
	router.Use(middlewares.ValidateToken())
	router.GET("", controllers.GetAllMovies)
	router.GET("/:id", controllers.GetMovieById)
}
