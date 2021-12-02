package main

import (
	"context"
	"hw8/internal/http"
	"hw8/internal/store/postgres"
)

func main() {
	urlExample := "postgres://postgres:postgres@localhost:5432/goods"
	store := postgres.NewDB()
	if err := store.Connect(urlExample); err != nil {
		panic(err)
	}
	defer store.Close()

	srv := http.NewServer(context.Background(), ":8080", store)
	if err := srv.Run(); err != nil {
		panic(err)
	}

	srv.WaitForGracefulTermination()
}
