package controllers

import (
	"fmt"
	"funtastix/backend/dto"
	"funtastix/backend/libs"
	"funtastix/backend/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Register godoc
// @Summary Register
// @Schemes
// @Description Register account
// @Tags auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param formUser formData dto.RegisterDTO true "form register"
// @Success 200 {object} models.Response
// @Router /auth/register [post]
func Register(ctx *gin.Context) {
	var formUser dto.RegisterDTO
	if err := ctx.ShouldBind(&formUser); err != nil {
		errMess := libs.RegisterErrHandler(err)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Succsess: false,
			Message:  errMess,
		})
		return
	}
	found := models.FindUserByEmail(strings.ToLower(formUser.Email))
	if found != (models.User{}) {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Succsess: false,
			Message:  "email not available",
		})
		return
	}

	isStrong := libs.StrongPasswordHandler(formUser)

	if isStrong != "" {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Succsess: false,
			Message:  isStrong,
		})
		return
	}

	hasher := libs.CreateHash(formUser.Password)
	formUser.Email = strings.ToLower(formUser.Email)

	formUser.Password = hasher

	profile := models.AddProfile()
	models.AddUser(formUser, profile.Id)

	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "register success",
	})
}

// Login godoc
// @Summary Login
// @Schemes
// @Description Login authentication
// @Tags auth
// @Accept x-www-form-urlencoded
// @Produce json
// @Param form formData dto.LoginDTO true "form login"
// @Success 200 {object} models.Response
// @Router /auth/login [post]
func Login(ctx *gin.Context) {
	var form dto.LoginDTO
	err := ctx.ShouldBind(&form)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Succsess: false,
			Message:  "Unexpected error",
		})
	}

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
