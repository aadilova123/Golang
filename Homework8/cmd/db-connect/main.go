package main

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
)

func main() {
	ctx := context.Background()
	// postgres://postgres:mypassword@localhost:5432/postgres
	urlExample := "postgres://postgres:postgres@localhost:5432/goods"
	conn, err := pgx.Connect(ctx, urlExample)
	if err != nil {
		panic(err)
	}
	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {

		}
	}(conn, context.Background())

	if err := conn.Ping(ctx); err != nil {
		panic(err)
	}
	log.Println("Pinged DB")
}
