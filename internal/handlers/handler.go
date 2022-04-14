package handlers

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"my_life/internal/services"
)

type handler struct {
	userHandler    *userHandler
	sectionHandler *sectionHandler
	taskHandler    *taskHandler
	subtaskHandler *subtaskHandler
}

func NewHandler(service *services.Service) *handler {
	return &handler{
		userHandler:    newUserHandler(service),
		sectionHandler: newSectionHandler(service),
		taskHandler:    newTaskHandler(service),
		subtaskHandler: newSubtaskHandler(service),
	}
}

func (h *handler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	r.Route("/", func(r chi.Router) {
		r.Post("/sign-up", h.signUp)
		r.Post("/sign-in", h.signIn)
	})

	r.With(h.verifyToken).Route("/lists", func(r chi.Router) {
		r.Post("/", h.createList)
		r.Get("/", h.getListsByUId)
	})

	r.With(h.verifyToken).Route("/sections", func(r chi.Router) {
		r.Post("/", h.sectionHandler.createSection)
		r.Get("/", h.sectionHandler.getSectionsByUId)
	})

	r.With(h.verifyToken).Route("/tasks", func(r chi.Router) {
		r.Post("/", h.taskHandler.createTask)
		r.Get("/", h.taskHandler.getTasksByUId)
	})

	r.With(h.verifyToken).Route("/subtasks", func(r chi.Router) {
		r.Post("/", h.subtaskHandler.createSubtask)
		r.Get("/", h.subtaskHandler.getSubTasksByUId)
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.userHandler.getUserById)
		r.Get("/full", h.userHandler.getFullUserInfo)
	})

	return r
}
