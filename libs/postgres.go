package libs

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

func DB() *pgx.Conn {
	godotenv.Load()
	config, _ := pgx.ParseConfig("")
	conn, err := pgx.Connect(context.Background(), config.ConnString())
	if err != err {
		fmt.Println(err)
		os.Exit(1)
	}
	return conn
}
