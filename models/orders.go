package models

import (
	"context"
	"fmt"
	"funtastix/backend/libs"
	"time"

	"github.com/jackc/pgx/v5"
)

type OrderInput struct {
	Id              int
	UserId          int `form:"user_id"`
	MovieId         int `form:"movie_id"`
	DateId          int `form:"date_id"`
	TimeId          int `form:"time_id"`
	LocationId      int `form:"location_id"`
	CinemaId        int `form:"cinema_id"`
	SeatId          int `form:"seat_id"`
	PaymentMethodId int `form:"payment_method_id"`
}

type OrderDetails struct {
	Id            int         `json:"id"`
	MovieName     []string    `json:"movie" db:"title"`
	Email         any         `json:"email" db:"email"`
	Date          []time.Time `json:"date"`
	Time          any         `json:"time"`
	Location      any         `json:"location"`
	Cinema        any         `json:"cinema"`
	Seat          any         `json:"seat"`
	Price         int         `json:"price"`
	PaymentMethod any         `json:"paymentMethod" db:"method"`
}

type Orders []OrderDetails

func AllOrdersDetail() Orders {
	conn := libs.DB()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), `
    SELECT 
      orders.id, ARRAY_AGG(DISTINCT movies.title) as title, ARRAY_AGG(DISTINCT users.email) as email, ARRAY_AGG(DISTINCT show_dates.date) as date, ARRAY_AGG(DISTINCT show_times.time) as time, ARRAY_AGG(DISTINCT show_locations.location) as location, ARRAY_AGG(DISTINCT show_cinemas.cinema) as cinema, ARRAY_AGG(DISTINCT seats.seat) as seat, SUM(seats.price) as price, ARRAY_AGG(DISTINCT payment_methods.method) as method
    FROM
      users
    JOIN
      orders ON users.id = orders.user_id
    JOIN
      movies ON orders.movie_id = movies.id
		JOIN 
			show_dates ON orders.date_id = show_dates.id
		JOIN 
			show_times ON orders.time_id = show_times.id
		JOIN 
			show_locations ON orders.location_id = show_locations.id
		JOIN 
			show_cinemas ON orders.cinema_id = show_cinemas.id
		JOIN 
			seats ON orders.seat_id = seats.id
		JOIN 
			payment_methods ON orders.payment_method_id = payment_methods.id
		GROUP BY
			orders.id
  `)
	if err != nil {
		fmt.Println(err)
	}
	orders, err := pgx.CollectRows(rows, pgx.RowToStructByName[OrderDetails])
	if err != nil {
		fmt.Println(err)
	}
	return orders
}

func AddOrder(formOrder OrderInput) OrderInput {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var order OrderInput
	conn.QueryRow(context.Background(), `
    INSERT INTO orders (user_id, movie_id, date_id, time_id, location_id, cinema_id, seat_id, payment_method_id)
    VALUES
      ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING id
  `, formOrder.UserId, formOrder.MovieId, formOrder.DateId, formOrder.TimeId, formOrder.LocationId, formOrder.CinemaId, formOrder.SeatId, formOrder.PaymentMethodId).Scan(&order.Id)
	return order
}

func SelectOneOrder(orderId int) OrderDetails {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var order OrderDetails
	conn.QueryRow(context.Background(), `
    SELECT 
      orders.id, ARRAY_AGG(DISTINCT movies.title) as title, ARRAY_AGG(DISTINCT users.email) as email, ARRAY_AGG(DISTINCT show_dates.date) as date, ARRAY_AGG(DISTINCT show_times.time) as time, ARRAY_AGG(DISTINCT show_locations.location) as location, ARRAY_AGG(DISTINCT show_cinemas.cinema) as cinema, ARRAY_AGG(DISTINCT seats.seat) as seat, SUM(seats.price) as price, ARRAY_AGG(DISTINCT payment_methods.method) as method
    FROM
      users
    JOIN
      orders ON users.id = orders.user_id
    JOIN
      movies ON orders.movie_id = movies.id
		JOIN 
			show_dates ON orders.date_id = show_dates.id
		JOIN 
			show_times ON orders.time_id = show_times.id
		JOIN 
			show_locations ON orders.location_id = show_locations.id
		JOIN 
			show_cinemas ON orders.cinema_id = show_cinemas.id
		JOIN 
			seats ON orders.seat_id = seats.id
		JOIN 
			payment_methods ON orders.payment_method_id = payment_methods.id
		WHERE 
			orders.id = $1
		GROUP BY
			orders.id
  `, orderId).Scan(&order.Id, &order.MovieName, &order.Email, &order.Date, &order.Time, &order.Location, &order.Cinema, &order.Seat, &order.Price, &order.PaymentMethod)
	return order
}
