package handlers

import (
	"context"
	"github.com/go-chi/chi/v5"
	"my_life/internal/domain"
	"my_life/internal/services"
	"net/http"
)

type TasksHandler interface {
	CreateList(ctx context.Context, list domain.TaskList) error
}

type tasks struct {
	service TasksHandler
}

func NewTasksHandler(service *services.TaskService) *tasks {
	return &tasks{service: service}
}

func (t *tasks) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/hello", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Sosi dick"))
	}))

	r.Route("/lists", func(r chi.Router) {
		r.Post("/", t.createList)
	})

	return r
}

func (t tasks) CreateList(ctx context.Context, list domain.TaskList) error {
	return t.service.CreateList(ctx, list)
}

func (t tasks) createList(w http.ResponseWriter, r *http.Request) {

}
