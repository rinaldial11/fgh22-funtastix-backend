package models

import (
	"context"
	"fmt"
	"funtastix/backend/libs"

	"github.com/jackc/pgx/v5"
)

type PageInfo struct {
	CurrentPage int `json:"currentPage,omitempty"`
	NextPage    int `json:"nextPage,omitempty"`
	PrevPage    int `json:"prevPage,omitempty"`
	TotalPage   int `json:"totalPage,omitempty"`
	TotalData   int `json:"totalData,omitempty"`
}

type Response struct {
	Succsess bool   `json:"success"`
	Message  string `json:"message"`
	PageInfo any    `json:"pageInfo,omitempty"`
	Results  any    `json:"results,omitempty"`
}

type User struct {
	Id       int    `json:"id"`
	Fullname string `json:"fullname" form:"fullname"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type ListUsers []User

func SelectOneUsers(idUser int) User {
	conn := libs.DB()
	defer conn.Close(context.Background())
	var user User

	conn.QueryRow(context.Background(), `
    SELECT id, email, password
    FROM users
    WHERE
    id = $1
  `, idUser).Scan(&user.Id, &user.Email, &user.Password)
	return user
}

func GetAllUsers(page int, limit int, orderBy string, order string) ListUsers {
	conn := libs.DB()
	defer conn.Close(context.Background())

	modifyQuery := fmt.Sprintf("SELECT id, '' as fullname, email, password FROM users ORDER BY %s %s OFFSET $1 LIMIT $2", orderBy, order)
	offset := (page - 1) * limit
	rows, err := conn.Query(context.Background(), modifyQuery, offset, limit)
	if err != nil {
		fmt.Println(err)
	}
	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[User])
	if err != nil {
		fmt.Println(err)
	}
	return users
}

func SearchUserByEmail(email string) ListUsers {
	conn := libs.DB()
	defer conn.Close(context.Background())

	emailSubstring := "%" + email + "%"
	rows, err := conn.Query(context.Background(), `
		SELECT users.id, '' as fullname, users.email, users.password
		FROM users
		WHERE 
		email ILIKE $1
	`, emailSubstring)
	if err != nil {
		fmt.Println(err)
	}
	users, err := pgx.CollectRows(rows, pgx.RowToStructByName[User])
	if err != nil {
		fmt.Println(err)
	}
	return users
}

func FindUserByEmail(email string) User {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var user User
	conn.QueryRow(context.Background(), `
		SELECT id, email, password
		FROM users
		WHERE
		email = $1
	`, email).Scan(&user.Id, &user.Email, &user.Password)
	return user
}

func AddUser(userData User) User {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var user User
	conn.QueryRow(context.Background(), `
		INSERT INTO users (email, password)
		values
		($1, $2)
	`, userData.Email, userData.Password).Scan(&user.Email, &user.Password)
	return user
}

func UpdateUser(userData User) User {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var updatedUser User
	conn.QueryRow(context.Background(), `
		UPDATE users SET email=$1, password=$2 WHERE id=$3
		RETURNING id, email, password
	`, userData.Email, userData.Password, userData.Id).Scan(&updatedUser.Id, &updatedUser.Email, &updatedUser.Password)
	return updatedUser
}

func DropUser(id int) User {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var deletedUser User
	conn.QueryRow(context.Background(), `
		DELETE FROM users
		WHERE id = $1
		RETURNING id, email, password
	`, id).Scan(&deletedUser.Id, &deletedUser.Email, &deletedUser.Password)
	return deletedUser
}

func CountUser(search string) int {
	conn := libs.DB()
	defer conn.Close(context.Background())

	titleSubstring := "%" + search + "%"
	var total int
	conn.QueryRow(context.Background(), `
		SELECT COUNT(users.id) 
		FROM users
		WHERE email ILIKE $1
	`, titleSubstring).Scan(&total)
	return total
}
