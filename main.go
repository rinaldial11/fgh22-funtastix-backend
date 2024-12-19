package main

import (
	"funtastix/backend/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()

	route.Static("/movies/images", "uploads/movies")
	route.MaxMultipartMemory = 2 << 20

	routers.Routers(route)
	route.Run(":8888")
}
