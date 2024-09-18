package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

func Connect() error {
	dbURL := os.Getenv("DATABASE_URL")
	var err error

	Conn, err = pgx.Connect(context.Background(), dbURL)
	if err != nil {
		return fmt.Errorf("unable to connect to database: %v", err)
	}

	return nil
}

func Close() {
	Conn.Close(context.Background())
}
