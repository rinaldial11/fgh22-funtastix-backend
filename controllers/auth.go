package controllers

import (
	"fmt"
	"funtastix/backend/libs"
	"funtastix/backend/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var formUser models.User
	ctx.ShouldBind(&formUser)
	found := models.FindUserByEmail(strings.ToLower(formUser.Email))
	if found != (models.User{}) {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Succsess: false,
			Message:  "email not available",
		})
		return
	}
	if len(formUser.Email) < 8 || !strings.Contains(formUser.Email, "@") {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Succsess: false,
			Message:  "email must be 8 character and contains @",
		})
		return
	}
	if len(formUser.Password) < 6 {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Succsess: false,
			Message:  "password length at least 6 chatacter",
		})
		return
	}
	hasher := libs.CreateHash(formUser.Password)
	formUser.Email = strings.ToLower(formUser.Email)

	if strings.Contains(formUser.Email, "admin") {
		formUser.Role = "admin"
	} else {
		formUser.Role = "user"
	}
	formUser.Password = hasher

	profile := models.AddProfile()
	models.AddUser(formUser, profile.Id)
	// models.Register(formUser)

	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "register success",
	})
}

func Login(ctx *gin.Context) {
	var form models.User
	err := ctx.ShouldBind(&form)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Succsess: false,
			Message:  "Unexpected error",
		})
	}

	fmt.Println(form.Email)

	user := models.FindUserByEmail(form.Email)
	isValid := libs.HashValidator(form.Password, user.Password)
	if isValid {
		token := libs.GenerateToken(struct {
			UserID int `json:"userId"`
		}{
			UserID: user.Id,
		})
		ctx.JSON(http.StatusOK, models.Response{
			Succsess: true,
			Message:  "login success",
			Results:  token,
		})
		return
	}
	ctx.JSON(http.StatusUnauthorized, models.Response{
		Succsess: false,
		Message:  "wrong email or password",
	})
}
