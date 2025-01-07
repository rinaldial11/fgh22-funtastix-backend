package middlewares

import (
	"fmt"
	"funtastix/backend/libs"
	"funtastix/backend/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func ValidateToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		head := ctx.GetHeader("Authorization")
		if head == "" {
			ctx.JSON(http.StatusUnauthorized, models.Response{
				Succsess: false,
				Message:  "Unauthorized",
			})
			ctx.Abort()
			return
		}
		bearer := strings.Split(head, " ")[0]
		// log.Println(bearer)
		if bearer != "Bearer" {
			ctx.JSON(http.StatusUnauthorized, models.Response{
				Succsess: false,
				Message:  "Invalid token",
			})
			ctx.Abort()
			return
		}
		claims, err := libs.ValidateToken(head)
		if err != nil {
			fmt.Println(err.Error())
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.Response{
				Succsess: false,
				Message:  "Unexpected error",
			})
		}
		ctx.Set("claims", claims)

		ctx.Next()
	}
}
