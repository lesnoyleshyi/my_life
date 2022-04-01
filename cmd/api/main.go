package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"my_life/internal/handlers"
	"net/http"

	"my_life/internal/repository"
	"my_life/internal/services"
	"my_life/pkg/postgres"
)

const connstr = "postgres://go_user:8246go@postgres_my_life:5432/taskstore"

func main() {
	dbPool, err := postgres.NewPool(connstr)
	if err != nil {
		log.Fatalf("Error connecting database: %v\n", err)
	}
	defer dbPool.Close()

	repo := repository.NewRepository(dbPool)

	service := services.NewTaskService(repo)

	handler := handlers.NewTasksHandler(service)
	r := chi.NewRouter()

	r.Mount("/", handler.Routes())

	log.Fatal(http.ListenAndServe("api:8080", r))
}
