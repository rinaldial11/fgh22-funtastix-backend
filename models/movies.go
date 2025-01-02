package models

import (
	"context"
	"fmt"
	"funtastix/backend/dto"
	"funtastix/backend/libs"
	"log"
	"time"

	"github.com/jackc/pgx/v5"
)

type Movie struct {
	Id          int       `json:"id" example:"1"`
	Title       string    `json:"title" example:"Avengers : Endgame"`
	Image       string    `json:"image" example:"example/avengers.jpg"`
	Banner      string    `json:"banner" example:"example/avengers.jpg"`
	ReleaseDate time.Time `json:"releaseDate" example:"2018-12-12"`
	Author      string    `json:"author" example:"Anthony Russo, Joe Russo"`
	Duration    string    `json:"duration" example:"03:02:00"`
	Synopsis    string    `json:"synopsis" example:"The Avengers assemble to reverse the damage caused by Thanos in Avengers: Infinity War."`
	Genre       any       `json:"genre"`
	Cast        any       `json:"casts"`
	UploadedBy  any       `json:"uploadedBy"`
}

type MovieHome struct {
	Id    int    `json:"id" form:"id" example:"1"`
	Title string `json:"title" form:"title" example:"Avengers : Endgame"`
	Image string `json:"image" form:"image" example:"example/avengers.jpg"`
	Genre any    `json:"genre"`
}

type ListMovieHome []MovieHome
type ListMovies []Movie

func GetAllMovies(page int, limit int, orderBy string, order string) ListMovieHome {
	conn := libs.DB()
	defer conn.Close(context.Background())

	offset := (page - 1) * limit
	query := fmt.Sprintf(`
  SELECT
    movies.id,
    movies.title,
    movies.image,
  ARRAY_AGG(
        DISTINCT movie_genre.genre_name
  ) AS genre
  FROM
    movies
  JOIN movie_genre ON movies.id = movie_genre.movie_id
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
	movies, err := pgx.CollectRows(rows, pgx.RowToStructByName[MovieHome])
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
  users.email as uploaded_by
  FROM
    movies
  JOIN movie_genre ON movies.id = movie_genre.movie_id
  JOIN movie_cast ON movies.id = movie_cast.movie_id
  JOIN users ON movies.uploaded_by = users.id
  WHERE movies.id = $1
  GROUP BY
    movies.id, users.email
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

func DropMovie(id int) Movie {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var deletedMovie Movie
	err := conn.QueryRow(context.Background(), `
    DELETE FROM movies 
    WHERE id = $1
    RETURNING id, title, image, banner, release_date, author, duration, synopsis, uploaded_by
  `, id).Scan(&deletedMovie.Id, &deletedMovie.Title, &deletedMovie.Image, &deletedMovie.Banner, &deletedMovie.ReleaseDate, &deletedMovie.Author, &deletedMovie.Duration, &deletedMovie.Synopsis, &deletedMovie.UploadedBy)

	if err != nil {
		fmt.Println(err)
	}
	return deletedMovie
}

func DropMovieGenre(idMovie int) int {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var movieId int
	err := conn.QueryRow(context.Background(), `
    DELETE FROM movie_genre
    WHERE movie_id = $1
    RETURNING movie_id
  `, idMovie).Scan(&movieId)
	if err != nil {
		fmt.Println(err)
	}
	return movieId
}

func DropMovieCast(idMovie int) int {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var movieId int
	err := conn.QueryRow(context.Background(), `
    DELETE FROM movie_cast
    WHERE movie_id = $1
    RETURNING movie_id
  `, idMovie).Scan(&movieId)
	if err != nil {
		fmt.Println(err)
	}
	return movieId
}

func AddMovie(movie dto.MovieDTO, userId int) int {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var movieId int
	err := conn.QueryRow(context.Background(), `
    INSERT INTO movies 
      (title, image, banner, release_date, author, duration, synopsis, uploaded_by)
    VALUES 
      ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING
      id
  `, movie.Title, movie.Image, movie.Banner, movie.ReleaseDate, movie.Author, movie.Duration, movie.Synopsis, userId).Scan(&movieId)
	if err != nil {
		fmt.Println(err)
	}
	return movieId
}

func AddMovieGenre(movie dto.MovieDTO, movieId int) dto.MovieDTO {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var movieGenre dto.MovieDTO

	bq := "INSERT INTO movie_genre (movie_id, genre_name) VALUES"

	var values []interface{}

	// if movie.Genre != nil {
	// 	bq += "($1, $2)"
	// 	values = append(values, movieId, movie.Genre)
	// }

	for i, v := range movie.Genre {
		if len(values) > 0 {
			bq += ","
		}
		j := i + 1
		bq += fmt.Sprintf("($%d,$%d)", 2*j-1, 2*j)
		values = append(values, movieId, v)
	}
	bq += "RETURNING genre_name"
	log.Println(bq)
	log.Println(values)
	err := conn.QueryRow(context.Background(), bq, values...).Scan(&movieGenre.Genre)
	if err != nil {
		fmt.Println(err)
	}
	return movieGenre
}

func AddMovieCast(movie dto.MovieDTO, movieId int) dto.MovieDTO {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var movieCast dto.MovieDTO

	bq := "INSERT INTO movie_cast (movie_id, cast_name) VALUES"

	var values []interface{}

	// if movie.Cast != nil {
	// 	bq += "($1, $2)"
	// 	values = append(values, movieId, movie.Cast)
	// }

	for i, v := range movie.Cast {
		if len(values) > 0 {
			bq += ","
		}
		j := i + 1
		bq += fmt.Sprintf("($%d,$%d)", 2*j-1, 2*j)
		values = append(values, movieId, v)
	}

	bq += "RETURNING cast_name"

	err := conn.QueryRow(context.Background(), bq, values...).Scan(&movieCast.Cast)
	if err != nil {
		fmt.Println(err)
	}
	return movieCast
}
