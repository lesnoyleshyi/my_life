package handlers

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"my_life/internal/services"
)

type handler struct {
	services       *services.Service
	sectionHandler sectionHandler
	taskHandler    taskHandler
}

func NewHandler(service *services.Service) *handler {
	return &handler{
		services:       service,
		sectionHandler: *newSectionHandler(service),
		taskHandler:    *newTaskHandler(service),
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

	r.With(h.verifyToken).Route("/", func(r chi.Router) {
		r.Post("/", h.taskHandler.createTask)
		r.Get("/", h.taskHandler.getTasksByUId)
	})
	//
	//r.With(h.verifyToken).Route("/", func(r chi.Router) {
	//	r.Post("/", h.createSubtask)
	//	r.Get("/", h.getSubTasksByUId)
	//})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.getUserById)
	})

	return r
}
