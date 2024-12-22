package routers

import "github.com/gin-gonic/gin"

func Routers(router *gin.Engine) {
	UserRouter(router.Group("/users"))
	ProfileRouter(router.Group("/profiles"))
	MovieRouter(router.Group("/movies"))
	// MovieAdminRouter(router.Group("/movies"))
	AuthRouter(router.Group("/auth"))
	OrderRouter(router.Group("/orders"))
}
