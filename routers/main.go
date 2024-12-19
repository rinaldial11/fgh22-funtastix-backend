package routers

import "github.com/gin-gonic/gin"

func Routers(router *gin.Engine) {
	UserRouter(router.Group("/users"))
	// MovieRouter(router.Group("/movies"))
	// MovieAdminRouter(router.Group("/movies"))
	AuthRouter(router.Group("/auth"))
}
