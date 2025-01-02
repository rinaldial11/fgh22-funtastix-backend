package controllers

import (
	"encoding/json"
	"fmt"
	"funtastix/backend/dto"
	"funtastix/backend/libs"
	"funtastix/backend/models"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

// Edit Profile godoc
// @Schemes
// @Summary Edit profile
// @Description Edit current logged in profile
// @Tags profiles
// @Accept mpfd
// @Produce json
// @Param foundUser formData dto.ProfileDTO true "profile user"
// @Param picture formData file false "profile user"
// @Success 200 {object} models.Response{results=models.Profile}
// @Security ApiKeyAuth
// @Router /profiles [patch]
func EditProfile(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	claimsJson, err := json.Marshal(claims)
	picture, _ := ctx.FormFile("picture")
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Succsess: false,
			Message:  "Unexpected error",
		})
	}
	var claimsStruct libs.ClaimsWithPayload
	err = json.Unmarshal(claimsJson, &claimsStruct)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Succsess: false,
			Message:  "Unexpected error",
		})
	}
	if claimsStruct.UserID == 0 {
		ctx.JSON(http.StatusForbidden, models.Response{
			Succsess: false,
			Message:  "Invalid token",
		})
	}
	foundUser := models.SelectOneUsers(claimsStruct.UserID)
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
	models.UpdateUser(foundUser, claimsStruct.UserID)
	foundUser = models.SelectOneProfile(claimsStruct.UserID)

	if err = ctx.ShouldBind(&foundUser); err != nil {
		log.Println("---------------------", err)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Succsess: false,
			Message:  "Invalid input data",
		})
		return
	}

	if picture != nil {
		filename := uuid.New().String()
		splitedfilename := strings.Split(picture.Filename, ".")
		ext := splitedfilename[len(splitedfilename)-1]
		if ext != "jpg" && ext != "png" && ext != "jpeg" {
			ctx.JSON(http.StatusBadRequest, models.Response{
				Succsess: false,
				Message:  "wrong file format",
			})
			return
		}
		storedFile := fmt.Sprintf("%s.%s", filename, ext)
		ctx.SaveUploadedFile(picture, fmt.Sprintf("uploads/profile/%s", storedFile))
		foundUser.Picture = storedFile
	}

	if foundUser.Point == "" {
		foundUser.Point = "0"
	}

	updatedProfile := models.EditProfile(foundUser, claimsStruct.UserID)
	if updatedProfile == (models.Profile{}) {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Succsess: false,
			Message:  "Failed to update profile",
		})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "profile updated",
		Results:  updatedProfile,
	})
}

func GetProfileById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	foundProfile := models.SelectOneProfile(id)

	if foundProfile == (dto.ProfileDTO{}) {
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

// Profile godoc
// @Summary Profile Info
// @Description Get current logged in profile info
// @Tags profiles
// @Accept json
// @Produce json
// @Success 200 {object} models.Response{results=models.Profile}
// @Security ApiKeyAuth
// @Router /profiles [get]
func GetCurrentProfile(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	claimsJson, err := json.Marshal(claims)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Succsess: false,
			Message:  "Unexpected error",
		})
	}

	var claimsStruct libs.ClaimsWithPayload
	err = json.Unmarshal(claimsJson, &claimsStruct)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Succsess: false,
			Message:  "Unexpected error",
		})
	}
	if claimsStruct.UserID == 0 {
		ctx.JSON(http.StatusForbidden, models.Response{
			Succsess: false,
			Message:  "Invalid token",
		})
	}
	profile := models.SelectCurrentProfile(claimsStruct.UserID)
	if profile.Point == "" {
		profile.Point = "0"
	}
	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "Profile info",
		Results:  profile,
	})
}
