package middlewares

import (
	"funtastix/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

var limiter = rate.NewLimiter(1, 1)

func RateLimiter(c *gin.Context) {
	if !limiter.Allow() {
		c.JSON(http.StatusTooManyRequests, models.Response{
			Succsess: false,
			Message:  "Too many requests",
		})
		c.Abort()
		return
	}
	c.Next()
}
