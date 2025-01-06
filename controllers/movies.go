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

// ListMovies godoc
// @Summary      List Movies
// @Description  get all movies
// @Tags         movies
// @Accept       json
// @Produce      json
// @Param search query string false "Search"
// @Param page query string false "Page"
// @Param limit query string false "Limit"
// @Success      200  {object} models.Response{results=models.ListMovieHome}
// @
// @Router       /movies [get]
func GetAllMovies(ctx *gin.Context) {
	search := ctx.Query("search")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "5"))
	order := strings.ToLower(ctx.DefaultQuery("order", "ASC"))
	orderBy := ctx.DefaultQuery("sort_by", "id")
	// movies := models.GetAllMovies(page, limit, orderBy, order)
	// count := models.CountMovie(search)

	if order != "ASC" {
		order = "DESC"
	}

	var movies models.ListMovieHome
	var count int
	modifyRequestUri := fmt.Sprintf("count+%s", ctx.Request.RequestURI)

	get := libs.GetFromRedis(ctx.Request.RequestURI)
	getCount := libs.GetFromRedis(modifyRequestUri)

	if get.Val() != "" {
		raw := []byte(get.Val())
		if err := json.Unmarshal(raw, &movies); err != nil {
			fmt.Println(err)
		}
	} else {
		movies = models.GetAllMovies(page, limit, orderBy, order)
		encoded, _ := json.Marshal(movies)
		libs.SetToRedis(ctx.Request.RequestURI, encoded)
	}

	if getCount.Val() != "" {
		raw := []byte(getCount.Val())
		json.Unmarshal(raw, &count)
	} else {
		count = models.CountMovie(search)
		encoded, _ := json.Marshal(count)
		libs.SetToRedis(modifyRequestUri, encoded)
	}
	foundMovie := models.SearchMovieByTitle(search, page, limit, orderBy, order)
	if search != "" {
		if len(foundMovie) == 1 {
			ctx.JSON(http.StatusOK, models.Response{
				Succsess: true,
				Message:  "List all movies",
				PageInfo: models.PageInfo(libs.GetPageInfo(page, limit, count)),
				Results:  foundMovie[0],
			})
			return
		}
		ctx.JSON(http.StatusOK, models.Response{
			Succsess: true,
			Message:  "List all movies",
			PageInfo: models.PageInfo(libs.GetPageInfo(page, limit, count)),
			Results:  foundMovie,
		})
		return
	}
	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "List all movies",
		PageInfo: models.PageInfo(libs.GetPageInfo(page, limit, count)),
		Results:  movies,
	})
}

// MovieDetails godoc
// @Summary Movie Details
// @Description Get movie details
// @Tags movies
// @Accept json
// @Produce json
// @Param id path int false "movie id"
// @Success 200 {object} models.Response{results=models.Movie}
// @Router /movies/{id} [get]
func GetMovieById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	movie := models.SelectOneMovie(id)

	if movie == (models.Movie{}) {
		ctx.JSON(http.StatusNotFound, models.Response{
			Succsess: false,
			Message:  "movie not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "details movie",
		Results:  movie,
	})
}

// DeleteMovie godoc
// @Summary Delete movie
// @Description Delete movie
// @Tags movies
// @Accept json
// @Produce json
// @Param id path int false "movie id"
// @Success 200 {object} models.Response{results=models.Movie}
// @Security ApiKeyAuth
// @Router /movies/{id} [delete]
func DeleteMovie(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	movie := models.SelectOneMovie(id)

	if movie == (models.Movie{}) {
		ctx.JSON(http.StatusNotFound, models.Response{
			Succsess: false,
			Message:  "movie not found",
		})
		return
	}
	idDeletedGenre := models.DropMovieGenre(id)
	idDeletedCast := models.DropMovieCast(idDeletedGenre)
	deletedMovie := models.DropMovie(idDeletedCast)
	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "Movie deleted",
		Results:  deletedMovie,
	})
}

// AddMovie godoc
// @Schemes
// @Summary Add movie
// @Description Add movie
// @Tags movies
// @Accept mpfd
// @Produce json
// @Param formMovie formData dto.MovieDTO false "add movie"
// @Param image formData file false "add image"
// @Param banner formData file false "add banner"
// @Success 200 {object} models.Response{results=models.Movie}
// @Security ApiKeyAuth
// @Router /movies [post]
func AddMovie(ctx *gin.Context) {
	var formMovie dto.MovieDTO
	claims, _ := ctx.Get("claims")
	claimsJson, err := json.Marshal(claims)
	image, _ := ctx.FormFile("image")
	banner, _ := ctx.FormFile("banner")
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
	err = ctx.ShouldBind(&formMovie)
	if err != nil {
		fmt.Println(err)
	}
	log.Println(formMovie.Genre)
	if image != nil {
		filename := uuid.New().String()
		splitedfilename := strings.Split(image.Filename, ".")
		ext := splitedfilename[len(splitedfilename)-1]
		if ext != "jpg" && ext != "png" && ext != "jpeg" {
			ctx.JSON(http.StatusBadRequest, models.Response{
				Succsess: false,
				Message:  "wrong file format",
			})
			return
		}
		storedFile := fmt.Sprintf("%s.%s", filename, ext)
		ctx.SaveUploadedFile(image, fmt.Sprintf("uploads/profile/%s", storedFile))
		formMovie.Image = storedFile
	}
	if banner != nil {
		filename := uuid.New().String()
		splitedfilename := strings.Split(banner.Filename, ".")
		ext := splitedfilename[len(splitedfilename)-1]
		if ext != "jpg" && ext != "png" && ext != "jpeg" {
			ctx.JSON(http.StatusBadRequest, models.Response{
				Succsess: false,
				Message:  "wrong file format",
			})
			return
		}
		storedFile := fmt.Sprintf("%s.%s", filename, ext)
		ctx.SaveUploadedFile(banner, fmt.Sprintf("uploads/profile/%s", storedFile))
		formMovie.Banner = storedFile
	}

	addedMovieId := models.AddMovie(formMovie, claimsStruct.UserID)
	models.AddMovieGenre(formMovie, addedMovieId)
	models.AddMovieCast(formMovie, addedMovieId)
	addedMovie := models.SelectOneMovie(addedMovieId)
	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "Movie added successfully",
		Results:  addedMovie,
	})
}
