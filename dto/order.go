package dto

type OrderSeatDTO struct {
	Id      int   `swaggerignore:"true"`
	UserId  int   `form:"user_id" binding:"required"`
	OrderId int   `form:"order_id" binding:"required"`
	SeatId  []int `form:"seat_id[]" binding:"required"`
}

type PaymentMethodDTO struct {
	Id              int `swaggerignore:"true"`
	PaymentMethodId int `form:"method"`
}

type OrderDTO struct {
	Id              int   `swaggerignore:"true"`
	UserId          int   `swaggerignore:"true"`
	MovieId         int   `form:"movie_id"`
	DateId          int   `form:"date_id"`
	TimeId          int   `form:"time_id"`
	LocationId      int   `form:"location_id"`
	CinemaId        int   `form:"cinema_id"`
	SeatId          []int `form:"seat_id[]" swaggerignore:"true"`
	PaymentMethodId int   `form:"method"`
}
