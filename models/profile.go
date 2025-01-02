package models

import (
	"context"
	"fmt"
	"funtastix/backend/dto"
	"funtastix/backend/libs"

	"github.com/jackc/pgx/v5"
)

type Profile struct {
	Id          int    `json:"id" example:"1"`
	FirstName   string `json:"firstName" example:"Budiono"`
	LastName    string `json:"lastName" example:"Siregar"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber" example:"08516839587"`
	Point       string `json:"point" example:"0"`
	Picture     string `json:"picture" example:"03f91853-f686-4190-a854-06f32dc17da7.jpeg"`
}

type ListProfiles []Profile

func GetAllProfiles(page int, limit int, ordrerBy string, order string) ListProfiles {
	conn := libs.DB()
	defer conn.Close(context.Background())

	modifyQuery := fmt.Sprintf("SELECT id, first_name, last_name, phone_number, point, picture FROM profiles ORDER BY %s %s OFFSET $1 LIMIT $2", ordrerBy, order)
	offset := (page - 1) * limit
	rows, err := conn.Query(context.Background(), modifyQuery, offset, limit)

	if err != nil {
		fmt.Println(err)
	}
	profiles, err := pgx.CollectRows(rows, pgx.RowToStructByName[Profile])
	if err != nil {
		fmt.Println(err)
	}
	return profiles
}

func AddProfile() Profile {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var profile Profile
	conn.QueryRow(context.Background(), `
		INSERT INTO profiles (first_name, last_name, phone_number, point, picture)
		VALUES ('', '', '', '', '')
		RETURNING id 
	`).Scan(&profile.Id)
	return profile
}

func EditProfile(profileData dto.ProfileDTO, userId int) Profile {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var editedProfile Profile
	conn.QueryRow(context.Background(), `
		UPDATE profiles SET first_name=$1, last_name=$2, phone_number=$3, point=$4, picture=$5
		WHERE id=$6
		RETURNING id, first_name, last_name, phone_number, point, picture
	`, profileData.FirstName, profileData.LastName, profileData.PhoneNumber, profileData.Point, profileData.Picture, userId).Scan(&editedProfile.Id, &editedProfile.FirstName, &editedProfile.LastName, &editedProfile.PhoneNumber, &editedProfile.Point, &editedProfile.Picture)
	return editedProfile
}

func SelectOneProfile(idProfile int) dto.ProfileDTO {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var profile dto.ProfileDTO
	err := conn.QueryRow(context.Background(), `
		SELECT first_name, last_name, users.email as email, phone_number, point, picture
		FROM profiles
		JOIN
			users ON profiles.id = users.profile_id
		WHERE
		profiles.id = $1
	`, idProfile).Scan(&profile.FirstName, &profile.LastName, &profile.Email, &profile.PhoneNumber, &profile.Point, &profile.Picture)
	if err != nil {
		fmt.Println(err)
	}
	return profile
}

func DropProfile(id int) Profile {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var deletedProfile Profile
	conn.QueryRow(context.Background(), `
		DELETE FROM profiles
		WHERE id = $1
		RETURNING id, first_name, last_name, phone_number, point, picture
	`, id).Scan(&deletedProfile.Id, &deletedProfile.FirstName, &deletedProfile.LastName, &deletedProfile.PhoneNumber, &deletedProfile.Point, &deletedProfile.Picture)
	return deletedProfile
}

func SearchProfileByName(name string) ListProfiles {
	conn := libs.DB()
	defer conn.Close(context.Background())

	nameSubstring := "%" + name + "%"
	rows, err := conn.Query(context.Background(), `
		SELECT id, first_name, last_name, phone_number, point, picture
		FROM profiles
		WHERE
		first_name ILIKE $1
	`, nameSubstring)
	if err != nil {
		fmt.Println(err)
	}
	profiles, err := pgx.CollectRows(rows, pgx.RowToStructByName[Profile])
	if err != nil {
		fmt.Println(err)
	}
	return profiles
}

func SelectCurrentProfile(idProfile int) Profile {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var profile Profile
	err := conn.QueryRow(context.Background(), `
		SELECT users.id, first_name, last_name, users.email as email, phone_number, point, picture
		FROM profiles
		JOIN
			users ON profiles.id = users.profile_id
		WHERE
		profiles.id = $1
	`, idProfile).Scan(&profile.Id, &profile.FirstName, &profile.LastName, &profile.Email, &profile.PhoneNumber, &profile.Point, &profile.Picture)
	if err != nil {
		fmt.Println(err)
	}
	return profile
}
