package handlers

import (
	"github.com/go-chi/chi/v5"
	"my_life/internal/repository"
	"my_life/internal/services"
)

type TasksHandler interface {
	repository.Repository
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
		r.Get("/", t.getListsByUId)
	})

	r.Route("/users", func(r chi.Router) {
		r.Post("/", t.createUser)
		r.Get("/", t.getUserById)
	})

	return r
}
