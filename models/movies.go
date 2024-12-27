package models

import (
	"context"
	"fmt"
	"funtastix/backend/libs"
	"time"

	"github.com/jackc/pgx/v5"
)

type Movie struct {
	Id          int       `json:"id" form:"id"`
	Title       string    `json:"title" form:"title"`
	Image       string    `json:"image" form:"image"`
	Banner      string    `json:"banner" form:"banner"`
	ReleaseDate time.Time `json:"releaseDate" form:"release_date"`
	Author      string    `json:"author" form:"author"`
	Duration    string    `json:"duration" form:"duration"`
	Synopsis    string    `json:"synopsis" form:"synopsis"`
	Genre       any       `json:"genre"`
	Cast        any       `json:"casts"`
	UploadedBy  any       `json:"uploadedBy"`
}

type ListMovies []Movie

func GetAllMovies(page int, limit int, orderBy string, order string) ListMovies {
	conn := libs.DB()
	defer conn.Close(context.Background())

	offset := (page - 1) * limit
	query := fmt.Sprintf(`
  SELECT
    movies.id,
    movies.title,
    movies.image,
    movies.banner,
    movies.release_date,
    movies.author,
    movies.duration,
    movies.synopsis,
  ARRAY_AGG(
        DISTINCT movie_genre.genre_name
  ) AS genre,
  ARRAY_AGG(DISTINCT movie_cast.cast_name) AS cast,
  ARRAY_AGG(DISTINCT users.email) as uploaded_by
  FROM
    movies
  JOIN movie_genre ON movies.id = movie_genre.movie_id
  JOIN movie_cast ON movies.id = movie_cast.movie_id
  JOIN users ON movies.uploaded_by = users.id
  GROUP BY
    movies.id
  ORDER BY movies.%s %s
  OFFSET $1 LIMIT $2
  `, orderBy, order)
	rows, err := conn.Query(context.Background(), query, offset, limit)
	if err != nil {
		fmt.Println(err)
	}
	movies, err := pgx.CollectRows(rows, pgx.RowToStructByName[Movie])
	if err != nil {
		fmt.Println(err)
	}
	return movies
}

func SelectOneMovie(id int) Movie {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var movie Movie
	conn.QueryRow(context.Background(), `
  SELECT
    movies.id,
    movies.title,
    movies.image,
    movies.banner,
    movies.release_date,
    movies.author,
    movies.duration,
    movies.synopsis,
  ARRAY_AGG(
        DISTINCT movie_genre.genre_name
  ) AS genre,
  ARRAY_AGG(DISTINCT movie_cast.cast_name) AS cast, 
  ARRAY_AGG(DISTINCT users.email) as uploaded_by
  FROM
    movies
  JOIN movie_genre ON movies.id = movie_genre.movie_id
  JOIN movie_cast ON movies.id = movie_cast.movie_id
  JOIN users ON movies.uploaded_by = users.id
  WHERE movies.id = $1
  GROUP BY
    movies.id
  `, id).Scan(&movie.Id, &movie.Title, &movie.Image, &movie.Banner, &movie.ReleaseDate, &movie.Author, &movie.Duration, &movie.Synopsis, &movie.Genre, &movie.Cast, &movie.UploadedBy)
	return movie
}

func SearchMovieByTitle(title string, page int, limit int, orderBy string, order string) ListMovies {
	conn := libs.DB()
	defer conn.Close(context.Background())

	offset := (page - 1) * limit
	modifyQuery := fmt.Sprintf(`
    SELECT
    movies.id,
    movies.title,
    movies.image,
    movies.banner,
    movies.release_date,
    movies.author,
    movies.duration,
    movies.synopsis,
  ARRAY_AGG(
        DISTINCT movie_genre.genre_name
  ) AS genre,
  ARRAY_AGG(DISTINCT movie_cast.cast_name) AS cast, 
  ARRAY_AGG(DISTINCT users.email) as uploaded_by
  FROM
    movies
  JOIN movie_genre ON movies.id = movie_genre.movie_id
  JOIN movie_cast ON movies.id = movie_cast.movie_id
  JOIN users ON movies.uploaded_by = users.id
  WHERE 
    movies.title ILIKE $1
  GROUP BY
    movies.id
  ORDER BY %s %s
  OFFSET $2 LIMIT $3
  `, orderBy, order)
	tittleSubstring := "%" + title + "%"
	rows, err := conn.Query(context.Background(), modifyQuery, tittleSubstring, offset, limit)
	if err != nil {
		fmt.Println(err)
	}
	movies, err := pgx.CollectRows(rows, pgx.RowToStructByName[Movie])
	if err != nil {
		fmt.Println(err)
	}
	return movies
}

func CountMovie(search string) int {
	conn := libs.DB()
	defer conn.Close(context.Background())

	titleSubstring := "%" + search + "%"
	var total int
	conn.QueryRow(context.Background(), `
    SELECT COUNT(movies.id)
    FROM movies
    WHERE title ILIKE $1
  `, titleSubstring).Scan(&total)
	return total
}
