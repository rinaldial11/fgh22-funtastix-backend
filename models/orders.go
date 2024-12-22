package models

import (
	"context"
	"fmt"
	"funtastix/backend/libs"
	"time"

	"github.com/jackc/pgx/v5"
)

type Order struct {
	Id            int       `json:"id"`
	ProfileId     int       `json:"profileId" form:"profile_id"`
	MovieId       int       `json:"movieId" form:"movie_id"`
	MovieName     string    `json:"movie" db:"title"`
	FirstName     string    `json:"name" db:"first_name"`
	Date          time.Time `json:"date" form:"date"`
	Time          string    `json:"time" form:"time"`
	Location      string    `json:"location" form:"location"`
	Cinema        string    `json:"cinema" form:"cinema"`
	Seat          string    `json:"seat" form:"seat"`
	PaymentMethod string    `json:"paymentMethod" form:"payment_method"`
}

type Orders []Order

func AllOrdersDetail() Orders {
	conn := libs.DB()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), `
    SELECT 
      orders.id, orders.profile_id, orders.movie_id, movies.title, profiles.first_name, orders.date, orders.time, orders.location, orders.cinema, orders.seat, orders.payment_method
    FROM
      profiles
    JOIN
      orders ON profiles.id = orders.profile_id
    JOIN
      movies ON orders.movie_id = movies.id
  `)
	if err != nil {
		fmt.Println(err)
	}
	orders, err := pgx.CollectRows(rows, pgx.RowToStructByName[Order])
	if err != nil {
		fmt.Println(err)
	}
	return orders
}

func AddOrder(formOrder Order) Order {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var order Order
	conn.QueryRow(context.Background(), `
    INSERT INTO orders (profile_id, movie_id, date, time, location, cinema, seat, payment_method)
    VALUES
      ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING id
  `, formOrder.ProfileId, formOrder.MovieId, formOrder.Date, formOrder.Time, formOrder.Location, formOrder.Cinema, formOrder.Seat, formOrder.PaymentMethod).Scan(&order.Id)
	return order
}

func SelectOneOrder(orderId int) Order {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var order Order
	conn.QueryRow(context.Background(), `
    SELECT 
      orders.id, orders.profile_id, orders.movie_id, movies.title, profiles.first_name, orders.date, orders.time, orders.location, orders.cinema, orders.seat, orders.payment_method
    FROM
      profiles
    JOIN
      orders ON profiles.id = orders.profile_id
    JOIN
      movies ON orders.movie_id = movies.id
    WHERE orders.id = $1
  `, orderId).Scan(&order.Id, &order.ProfileId, &order.MovieId, &order.MovieName, &order.FirstName, &order.Date, &order.Time, &order.Location, &order.Cinema, &order.Seat, &order.PaymentMethod)
	return order
}
