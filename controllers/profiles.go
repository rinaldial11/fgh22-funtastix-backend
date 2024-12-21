package controllers

import (
	"fmt"
	"funtastix/backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllProfiles(ctx *gin.Context) {
	allProfiles := models.GetAllProfiles()

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
