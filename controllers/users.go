package controllers

import (
	"funtastix/backend/libs"
	"funtastix/backend/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(ctx *gin.Context) {
	search := ctx.Query("search")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "5"))
	order := strings.ToLower(ctx.DefaultQuery("order", "ASC"))
	orderBy := ctx.DefaultQuery("sort_by", "id")
	count := models.CountUser(search)
	allUsers := models.GetAllUsers(page, limit, orderBy, order)

	if order != "ASC" {
		order = "DESC"
	}

	foundUser := models.SearchUserByEmail(search)
	if search != "" {
		if len(foundUser) == 1 {
			ctx.JSON(http.StatusOK, models.Response{
				Succsess: true,
				Message:  "list all users",
				PageInfo: models.PageInfo(libs.GetPageInfo(page, limit, count)),
				Results:  foundUser[0],
			})
			return
		}
		ctx.JSON(http.StatusOK, models.Response{
			Succsess: true,
			Message:  "list all users",
			PageInfo: models.PageInfo(libs.GetPageInfo(page, limit, count)),
			Results:  foundUser,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "list all users",
		PageInfo: models.PageInfo(libs.GetPageInfo(page, limit, count)),
		Results:  allUsers,
	})
}

func GetUserById(ctx *gin.Context) {
	idUser, _ := strconv.Atoi(ctx.Param("id"))
	foundUser := models.SelectOneUsers(idUser)

	if foundUser == (models.User{}) {
		ctx.JSON(http.StatusNotFound, models.Response{
			Succsess: false,
			Message:  "user not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "Details user",
		Results:  foundUser,
	})
}

func UpdateUser(ctx *gin.Context) {
	idUser, _ := strconv.Atoi(ctx.Param("id"))
	foundUser := models.SelectOneUsers(idUser)
	if foundUser == (models.User{}) {
		ctx.JSON(http.StatusNotFound, models.Response{
			Succsess: false,
			Message:  "user not found",
		})
		return
	}

	ctx.ShouldBind(&foundUser)
	if len(foundUser.Email) < 8 || !strings.Contains(foundUser.Email, "@") {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Succsess: false,
			Message:  "email must be 8 character and contains @",
		})
		return
	}
	if !strings.Contains(foundUser.Password, "$argon2i$v=19$m=65536,t=1") {
		if foundUser.Password != "" {
			if len(foundUser.Password) < 6 {
				ctx.JSON(http.StatusBadRequest, models.Response{
					Succsess: false,
					Message:  "password length at least 6 chatacter",
				})
				return
			}
			foundUser.Password = libs.CreateHash(foundUser.Password)
		}
	}
	updatedUser := models.UpdateUser(foundUser)
	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "user updated",
		Results:  updatedUser,
	})
}

func DeleteUser(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	profile := models.SelectOneProfile(id)

	if profile == (models.Profile{}) {
		ctx.JSON(http.StatusNotFound, models.Response{
			Succsess: false,
			Message:  "user not found",
		})
		return
	}
	deletedUser := models.DropUser(profile.Id)
	models.DropProfile(profile.Id)
	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "user deleted successfully",
		Results:  deletedUser,
	})
}

func CreateUser(ctx *gin.Context) {
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
	formUser.Password = hasher
	profileId := models.AddProfile()
	models.AddUser(formUser, profileId.Id)

	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "register success",
	})
}
