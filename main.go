package main

import (
	"funtastix/backend/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.Static("/profile/images", "uploads/profile")
	route.MaxMultipartMemory = 2 << 20

	routers.Routers(route)
	route.Run(":8888")
}
