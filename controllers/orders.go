package controllers

import (
	"encoding/json"
	"fmt"
	"funtastix/backend/dto"
	"funtastix/backend/libs"
	"funtastix/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllOrders(ctx *gin.Context) {
	orders := models.AllOrdersDetail()

	if len(orders) == 1 {
		ctx.JSON(http.StatusOK, models.Response{
			Succsess: true,
			Message:  "all orders list",
			Results:  orders[0],
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "all orders list",
		Results:  orders,
	})
}

// AddOrder godoc
// @Schemes
// @Summary Add order
// @Description Add orders
// @Tags orders
// @Accept x-www-form-urlencoded
// @Produce json
// @Param formMovie formData dto.OrderTempDTO false "add order"
// @Param seat_id[] formData array false "add seat order"
// @Success 200 {object} models.Response{results=models.OrderDetails}
// @Security ApiKeyAuth
// @Router /orders [post]
func AddOrder(ctx *gin.Context) {
	var order dto.OrderDTO
	claims, _ := ctx.Get("claims")
	claimsJson, err := json.Marshal(claims)
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
	if err := ctx.ShouldBind(&order); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, models.Response{
			Succsess: false,
			Message:  "Invalid input data",
		})
		return
	}

	order.UserId = claimsStruct.UserID
	orderID := models.AddOrder(order)
	models.AddSeatOrder(order.SeatId, orderID.Id)
	newOrder := models.SelectOneOrderSeat(orderID.Id)

	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "Your order details, please pay your order",
		Results:  newOrder,
	})
}

func GetAllPaymentMethods(ctx *gin.Context) {
	paymentMethods := models.GetAllPaymentMethods()

	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "List all payment methods",
		Results:  paymentMethods,
	})
}

func GetAllSeats(ctx *gin.Context) {
	seats := models.GetAllSeats()

	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "Seats layout info",
		Results:  seats,
	})
}

func AddSeatOrder(ctx *gin.Context) {
	var seatForm dto.OrderSeatDTO
	ctx.ShouldBind(&seatForm)
}
