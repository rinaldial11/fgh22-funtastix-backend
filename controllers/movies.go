package controllers

import (
	"funtastix/backend/libs"
	"funtastix/backend/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllMovies(ctx *gin.Context) {
	search := ctx.Query("search")
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "5"))
	order := strings.ToLower(ctx.DefaultQuery("order", "ASC"))
	orderBy := ctx.DefaultQuery("sort_by", "id")
	movies := models.GetAllMovies(page, limit, orderBy, order)
	count := models.CountMovie(search)

	if order != "ASC" {
		order = "DESC"
	}

	foundMovie := models.SearchMovieByTitle(search, page, limit, orderBy, order)
	if search != "" {
		if len(foundMovie) == 1 {
			ctx.JSON(http.StatusOK, models.Response{
				Succsess: true,
				Message:  "list all users",
				PageInfo: models.PageInfo(libs.GetPageInfo(page, limit, count)),
				Results:  foundMovie[0],
			})
			return
		}
		ctx.JSON(http.StatusOK, models.Response{
			Succsess: true,
			Message:  "list all users",
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
