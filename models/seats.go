package models

import (
	"context"
	"fmt"
	"funtastix/backend/dto"
	"funtastix/backend/libs"

	"github.com/jackc/pgx/v5"
)

type Seat struct {
	Id       int    `json:"id"`
	SeatName string `json:"seat" db:"seat"`
	Price    int    `json:"price" db:"price"`
}

type Seats []Seat

func GetAllSeats() Seats {
	conn := libs.DB()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), `
    SELECT id, seat, price FROM seats
  `)
	if err != nil {
		fmt.Println(err)
	}
	seats, err := pgx.CollectRows(rows, pgx.RowToStructByName[Seat])
	if err != nil {
		fmt.Println(err)
	}
	return seats
}

func AddOrderSeats(seat dto.OrderSeatDTO) dto.OrderSeatDTO {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var bookedSeat dto.OrderSeatDTO
	if err := conn.QueryRow(context.Background(), `
		INSERT INTO 
			seats_order(user_id, order_id, seat_id)
		VALUES
			($1, $2, $3)
		RETURNING id, user_id, order_id, seat_id
	`, seat.UserId, seat.OrderId, seat.SeatId).Scan(&bookedSeat.Id, &bookedSeat.OrderId, &bookedSeat.UserId, &bookedSeat.SeatId); err != nil {
		fmt.Println(err)
	}
	return bookedSeat
}
