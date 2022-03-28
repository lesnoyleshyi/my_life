package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"

	"my_life/internal/handlers"
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

	r.Get("/hello", func(writer http.ResponseWriter, request *http.Request) {
		if _, err := writer.Write([]byte("Sosi dick")); err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
		}
	})

	r.Mount("/", handler.Routes())

	log.Fatal(http.ListenAndServe(":80", r))
}
