package controllers

import (
	"funtastix/backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllOrders(ctx *gin.Context) {
	orders := models.AllOrdersDetail()

	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "all orders list",
		Results:  orders,
	})
}

func AddOrder(ctx *gin.Context) {
	var order models.Order
	ctx.ShouldBind(&order)

	orderID := models.AddOrder(order)
	newOrder := models.SelectOneOrder(orderID.Id)

	ctx.JSON(http.StatusOK, models.Response{
		Succsess: true,
		Message:  "please pay your order as soon as posible",
		Results:  newOrder,
	})
}
