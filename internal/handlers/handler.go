package handlers

import (
	"github.com/go-chi/chi/v5"
	"my_life/internal/services"
)

type handler struct {
	services *services.Service
}

func NewHandler(service *services.Service) *handler {
	return &handler{service}
}

func (h *handler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Route("/", func(r chi.Router) {
		r.Post("/sign-up", h.signUp)
		r.Post("/sign-in", h.signIn)
	})

	r.With(h.verifyToken).Route("/lists", func(r chi.Router) {
		r.Post("/", h.createList)
		r.Get("/", h.getListsByUId)
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.getUserById)
	})

	return r
}
