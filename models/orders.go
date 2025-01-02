package models

import (
	"context"
	"fmt"
	"funtastix/backend/dto"
	"funtastix/backend/libs"
	"time"

	"github.com/jackc/pgx/v5"
)

type OrderFirstDetails struct {
	Id         int    `json:"orderId"`
	UserId     int    `json:"userId"`
	MovieName  string `json:"movie"`
	MovieImage string `json:"image"`
	MovieGenre any    `json:"genre"`
	MovieTime  string `json:"time"`
}

type OrderPaymentDetails struct {
	Id          int       `json:"id"`
	Date        time.Time `json:"date"`
	Time        string    `json:"time"`
	MovieName   string    `json:"movie" db:"title"`
	Cinema      string    `json:"cinema"`
	TicketCount int       `json:"NumberOfTicket"`
	Price       int       `json:"price"`
	FullName    string    `json:"fullName"`
	Email       string    `json:"email" db:"email"`
	PhoneNumber string    `json:"phoneNumber"`
}

type OrderDetails struct {
	Id            int       `json:"id"`
	MovieName     string    `json:"movie" db:"title"`
	Email         string    `json:"email" db:"email"`
	Date          time.Time `json:"date"`
	Time          string    `json:"time"`
	Location      string    `json:"location"`
	Cinema        string    `json:"cinema"`
	Seat          any       `json:"seat"`
	Price         int       `json:"price"`
	PaymentMethod string    `json:"paymentMethod,omitempty" db:"method"`
}

type Orders []OrderDetails

func AllOrdersDetail() Orders {
	conn := libs.DB()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), `
    SELECT 
      orders.id, movies.title as title, users.email as email, show_dates.date as date, show_times.time as time, show_locations.location as location, show_cinemas.cinema as cinema, ARRAY_AGG(DISTINCT seats.seat) as seat, SUM(seats.price) as price, payment_methods.method as method
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
			orders.id, movies.title, users.email, show_dates.date, show_times.time, show_locations.location, show_cinemas.cinema, payment_methods.method
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

func AddOrder(formOrder dto.OrderDTO) int {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var orderId int
	conn.QueryRow(context.Background(), `
    INSERT INTO orders (user_id, movie_id, date_id, time_id, location_id, cinema_id)
    VALUES
      ($1, $2, $3, $4, $5, $6)
    RETURNING id
  `, formOrder.UserId, formOrder.MovieId, formOrder.DateId, formOrder.TimeId, formOrder.LocationId, formOrder.CinemaId).Scan(&orderId)
	return orderId
}

// func SelectOneOrderSeat(orderId int) OrderDetails {
// 	conn := libs.DB()
// 	defer conn.Close(context.Background())
// 	var order OrderDetails
// 	conn.QueryRow(context.Background(), `
//   SELECT
//       orders.id, movies.title as title, users.email as email, show_dates.date as date, show_times.time as time, show_locations.location as location, show_cinemas.cinema as cinema, ARRAY_AGG(DISTINCT seats.seat) as seat, SUM(seats.price) as price, payment_methods.method as method
//     FROM
//       users
//     JOIN
//       orders ON users.id = orders.user_id
//     JOIN
//       movies ON orders.movie_id = movies.id
// 		JOIN
// 			show_dates ON orders.date_id = show_dates.id
// 		JOIN
// 			show_times ON orders.time_id = show_times.id
// 		JOIN
// 			show_locations ON orders.location_id = show_locations.id
// 		JOIN
// 			show_cinemas ON orders.cinema_id = show_cinemas.id
// 		JOIN
// 			seats ON orders.seat_id = seats.id
// 		JOIN
// 			payment_methods ON orders.payment_method_id = payment_methods.id
// 		WHERE
// 			orders.id = $1
// 		GROUP BY
// 			orders.id, movies.title, users.email, show_dates.date, show_times.time, show_locations.location, show_cinemas.cinema, payment_methods.method
//   `, orderId).Scan(&order.Id, &order.MovieName, &order.Email, &order.Date, &order.Time, &order.Location, &order.Cinema, &order.Seat, &order.Price, &order.PaymentMethod)
// 	return order
// }

func SelectOneOrderFirst(orderId int) OrderFirstDetails {
	conn := libs.DB()
	defer conn.Close(context.Background())
	var order OrderFirstDetails
	err := conn.QueryRow(context.Background(), `
  SELECT 
      orders.id, orders.user_id, movies.title as title, movies.image as image, ARRAY_AGG(movie_genre.genre_name) as genre, show_times.time as time
    FROM
      movies
		JOIN
			movie_genre ON movies.id = movie_genre.movie_id
    JOIN
			orders ON movies.id = orders.movie_id
    JOIN
			show_times ON orders.time_id = show_times.id
		WHERE 
			orders.id = $1
		GROUP BY
			orders.id, orders.user_id, movies.title, movies.image, show_times.time
  `, orderId).Scan(&order.Id, &order.UserId, &order.MovieName, &order.MovieImage, &order.MovieGenre, &order.MovieTime)
	if err != nil {
		fmt.Println(err)
	}
	return order
}
