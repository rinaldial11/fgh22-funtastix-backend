package middlewares

import (
	"fmt"
	"funtastix/backend/libs"
	"funtastix/backend/models"
	"net/http"

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
		err := libs.ValidateToken(head)
		fmt.Println(err)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, models.Response{
				Succsess: false,
				Message:  "Unauthorized",
			})
			ctx.Abort()
		}

		ctx.Next()
	}
}
