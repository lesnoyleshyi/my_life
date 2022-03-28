package handlers

import (
	"context"
	"fmt"
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

	r.Route("/lists", func(r chi.Router) {
		r.Post("/", t.createList)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			_, err := fmt.Fprintf(w, "test ok\n")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
		})
	})

	return r
}

func (t tasks) CreateList(ctx context.Context, list domain.TaskList) error {
	return t.service.CreateList(ctx, list)
}

func (t tasks) createList(w http.ResponseWriter, r *http.Request) {

}
