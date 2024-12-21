package controllers

import (
	"fmt"
	"funtastix/backend/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllProfiles(ctx *gin.Context) {
	search := ctx.Query("search")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "5"))
	order := strings.ToLower(ctx.DefaultQuery("order", "ASC"))
	orderBy := ctx.DefaultQuery("sort_by", "id")
	allProfiles := models.GetAllProfiles(page, limit, orderBy, order)

	if order != "ASC" {
		order = "DESC"
	}

	foundProfile := models.SearchProfileByName(search)
	if search != "" {
		if len(foundProfile) == 1 {
			ctx.JSON(http.StatusOK, models.Response{
				Succsess: true,
				Message:  "list all users",
				Results:  foundProfile[0],
			})
			return
		}
		ctx.JSON(http.StatusOK, models.Response{
			Succsess: true,
			Message:  "list all users",
			Results:  foundProfile,
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "list all users",
		Results:  allProfiles,
	})
}

func EditProfile(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	foundProfile := models.SelectOneProfile(id)

	if foundProfile == (models.Profile{}) {
		ctx.JSON(http.StatusNotFound, models.Response{
			Succsess: false,
			Message:  "profile not found",
		})
		return
	}
	// ctx.ShouldBind(&foundProfile)
	if err := ctx.ShouldBind(&foundProfile); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Succsess: false,
			Message:  "Invalid input data",
		})
		return
	}

	updatedProfile := models.EditProfile(foundProfile)
	if updatedProfile == (models.Profile{}) {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Succsess: false,
			Message:  "Failed to update profile",
		})
		return
	}
	fmt.Println(updatedProfile)
	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "profile updated",
		Results:  updatedProfile,
	})
}

func GetProfileById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	foundProfile := models.SelectOneProfile(id)

	if foundProfile == (models.Profile{}) {
		ctx.JSON(http.StatusNotFound, models.Response{
			Succsess: false,
			Message:  "profile not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "Details user",
		Results:  foundProfile,
	})
}
