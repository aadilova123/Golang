package main

import (
	"context"
	"hw6/project/internal/http"
	"hw6/project/internal/store/inmemory"
	"log"
)

func main() {
	store := inmemory.NewDB()
	srv := http.NewServer(context.Background(), ":8080", store)
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
	srv.WaitForGracefulTermination()
}