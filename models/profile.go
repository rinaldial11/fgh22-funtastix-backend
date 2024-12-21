package models

import (
	"context"
	"fmt"
	"funtastix/backend/libs"

	"github.com/jackc/pgx/v5"
)

type Profile struct {
	Id          int    `json:"id"`
	FirstName   string `json:"firstName" form:"first_name"`
	LastName    string `json:"lastName" form:"last_name"`
	PhoneNumber string `json:"phoneNumber" form:"phone_number"`
	Point       string `json:"point" form:"point"`
	Picture     string `json:"picture" form:"picture"`
}

type ListProfiles []Profile

func GetAllProfiles() ListProfiles {
	conn := libs.DB()
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), `
		SELECT id, first_name, last_name, phone_number, point, picture 
		FROM profiles
	`)

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

func EditProfile(profileData Profile) Profile {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var editedProfile Profile
	conn.QueryRow(context.Background(), `
		UPDATE profiles SET first_name=$1, last_name=$2, phone_number=$3, point=$4, picture=$5
		WHERE id=$6
		RETURNING id, first_name, last_name, phone_number, point, picture
	`, profileData.FirstName, profileData.LastName, profileData.PhoneNumber, profileData.Point, profileData.Picture, profileData.Id).Scan(&editedProfile.Id, &editedProfile.FirstName, &editedProfile.LastName, &editedProfile.PhoneNumber, &editedProfile.Point, &editedProfile.Picture)
	return editedProfile
}

func SelectOneProfile(idProfile int) Profile {
	conn := libs.DB()
	defer conn.Close(context.Background())

	var profile Profile
	conn.QueryRow(context.Background(), `
		SELECT id, first_name, last_name, phone_number, point, picture
		FROM profiles
		WHERE
		id = $1
	`, idProfile).Scan(&profile.Id, &profile.FirstName, &profile.LastName, &profile.PhoneNumber, &profile.Point, &profile.Picture)
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
