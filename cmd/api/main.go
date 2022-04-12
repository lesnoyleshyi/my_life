package main

import (
	"github.com/go-chi/chi/v5"
	"log"
	"my_life/internal/handlers"
	"net/http"

	_ "my_life/docs"
	"my_life/internal/repository"
	"my_life/internal/services"
	"my_life/pkg/postgres"
)

const connstr = "postgres://go_user:8246go@postgres_my_life:5432/taskstore"

// @title			API for My Life application
// @version			0.1
// @description		It's made for testing purposes only

// @host			132.226.200.167:8080
// @BasePAth		/

func main() {
	dbPool, err := postgres.NewPool(connstr)
	if err != nil {
		log.Fatalf("Error connecting database: %v\n", err)
	}
	defer dbPool.Close()

	repo := repository.NewRepository(dbPool)

	service := services.NewService(repo)

	handler := handlers.NewHandler(service)
	r := chi.NewRouter()
	r.Mount("/", handler.Routes())

	log.Fatal(http.ListenAndServe("api:8080", r))
}
