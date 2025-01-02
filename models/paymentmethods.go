package models

import (
	"context"
	"fmt"
	"funtastix/backend/libs"

	"github.com/jackc/pgx/v5"
)

type PaymentMethod struct {
	Id     int    `json:"id"`
	Method string `json:"paymentMethod" db:"method"`
}

type ListPaymentMethods []PaymentMethod

func GetAllPaymentMethods() ListPaymentMethods {
	conn := libs.DB()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), `
    SELECT id, method FROM payment_methods
  `)
	if err != nil {
		fmt.Println(err)
	}
	methods, err := pgx.CollectRows(rows, pgx.RowToStructByName[PaymentMethod])
	if err != nil {
		fmt.Println(err)
	}
	return methods
}
