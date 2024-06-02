package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

var Conn *pgx.Conn

func init() {
	var err error
	ctx := context.Background()
	connStr := "postgres://jamiu:jamiu@localhost:5432/checklist?sslmode=disable"
	Conn, err = pgx.Connect(ctx, connStr)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
}
