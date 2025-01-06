package models

import (
	"context"
	"fmt"
	"funtastix/backend/dto"
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
	Id        int    `json:"id" db:"id"`
	ProfileId int    `json:"profileId,omitempty" db:"profile_id"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email" form:"email" db:"email" binding:"required"`
	Password  string `json:"password,omitempty" form:"password" db:"password" binding:"required"`
	Role      string `json:"role"`
}

type ListUsers []User

func SelectOneUsers(idUser int) dto.ProfileDTO {
	conn := libs.DB()
	defer conn.Close(context.Background())
	var user dto.ProfileDTO

	err := conn.QueryRow(context.Background(), `
    SELECT users.email, users.password
		FROM users
    WHERE
    users.id = $1
  `, idUser).Scan(&user.Email, &user.Password)
	if err != nil {
		fmt.Println(err)
	}
	return user
}

func GetAllUsers(page int, limit int, orderBy string, order string) ListUsers {
	conn := libs.DB()
	defer conn.Close(context.Background())

	modifyQuery := fmt.Sprintf(` 
		SELECT users.id, users.profile_id, profiles.first_name, profiles.last_name, users.email, users.password, users.role
		FROM users
    JOIN profiles ON users.profile_id = profiles.id
		ORDER BY %s %s 
		OFFSET $1 LIMIT $2`, orderBy, order)
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
		SELECT id, profile_id, email, password, role
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
	err := conn.QueryRow(context.Background(), `
		SELECT id, email, password
		FROM users
		WHERE
		email = $1
	`, email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		fmt.Println(err)
	}
	return user
}

func AddUser(userData dto.RegisterDTO, profile_id int) User {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var user User
	conn.QueryRow(context.Background(), `
		INSERT INTO users (profile_id, email, password, role)
		values
		($1, $2, $3, '')
	`, profile_id, userData.Email, userData.Password).Scan(&user.ProfileId, &user.Email, &user.Password)
	return user
}

func UpdateUser(userData dto.ProfileDTO, userId int) User {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var updatedUser User
	err := conn.QueryRow(context.Background(), `
		UPDATE users SET email=$1, password=$2 WHERE id=$3
		RETURNING id, profile_id, email, password, role
	`, userData.Email, userData.Password, userId).Scan(&updatedUser.Id, &updatedUser.ProfileId, &updatedUser.Email, &updatedUser.Password, &updatedUser.Role)
	if err != nil {
		fmt.Println(err)
	}
	return updatedUser
}

func DropUser(id int) User {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var deletedUser User
	conn.QueryRow(context.Background(), `
		DELETE FROM users
		WHERE id = $1
		RETURNING id, profile_id, email, password, role
	`, id).Scan(&deletedUser.Id, &deletedUser.ProfileId, &deletedUser.Email, &deletedUser.Password, &deletedUser.Role)
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

// func Register(user User) User {
// 	conn := libs.DB()
// 	tx, err := conn.Begin(context.Background())
// 	if err != nil {
// 		defer tx.Rollback(context.Background())
// 	}

// 	var profileId int
// 	tx.QueryRow(context.Background(), `
// 		INSERT INTO profiles (first_name, last_name, phone_number, point, picture)
// 		VALUES ('', '', '', '', '')
// 		RETURNING id
// 	`).Scan(&profileId)

// 	var newUser User
// 	tx.QueryRow(context.Background(), `
// 		INSERT INTO users (profile_id, email, password, role)
// 		values
// 		($1, $2, $3, $4)
// 		RETURNING id, profile_id, first_name, last_name, email, password, role
// 	`, profileId, user.Email, user.Password, user.Role).Scan(&newUser.Id, &newUser.ProfileId, &newUser.FirstName, &newUser.LastName, &newUser.Email, &newUser.Password, &newUser.Role)
// 	fmt.Println(newUser)
// 	tx.Commit(context.Background())
// 	return newUser
// }
