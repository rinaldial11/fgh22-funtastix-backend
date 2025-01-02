package dto

import "time"

type Genre string
type Cast string

type MovieDTO struct {
	Title       string    `form:"title"`
	Image       string    `form:"-" swaggerignore:"true"`
	Banner      string    `form:"-" swaggerignore:"true"`
	ReleaseDate time.Time `form:"release_date"`
	Author      string    `form:"author"`
	Duration    string    `form:"duration"`
	Synopsis    string    `form:"synopsis"`
	Genre       []string  `form:"genre_name" binding:"required"`
	Cast        []string  `form:"cast_name" binding:"required"`
}

// Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3MzU3NTg3ODQsInVzZXJJZCI6MX0.7DjHkTFre17fIsVYNIQHjAe0jHPF5xZdYD_wZDkkiRo
