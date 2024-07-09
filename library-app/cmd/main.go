package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"library-app/internal/handlers"
	"library-app/internal/repository/postgres"
	"library-app/internal/usecase"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()                                                   //
	connStr := "postgres://myuser:mypassword@localhost:5432/mydb?sslmode=disable" // DSN

	repo, err := postgres.NewPostgres(ctx, connStr) // Create a new Postgres repository
	if err != nil {
		log.Fatal(fmt.Errorf("postgres new: %w", err))
	}

	uc := usecase.NewUseCase(repo)

	r := mux.NewRouter() // Create a new router
	handlers.RegisterBookHandlers(r, uc)

	log.Fatal(http.ListenAndServe(":8000", r)) // Start the server
}
