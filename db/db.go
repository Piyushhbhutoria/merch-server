package db

import (
	"context"
	"fmt"
	"os"

	"github.com/Piyushhbhutoria/merch-server/config"
	"github.com/jackc/pgx/v4"
)

var db *pgx.Conn

func Init() {
	c := config.GetConfig()
	conn, err := pgx.Connect(context.Background(), c.GetString("db.endpoint"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	db = conn
	// defer db.Close(context.Background())
}

func GetDB() *pgx.Conn {
	return db
}
