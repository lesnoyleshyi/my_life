package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"my_life/internal/domain"
	"my_life/internal/services"
	"net/http"
	"strconv"
)

type TasksHandler interface {
	CreateList(ctx context.Context, list *domain.TaskList) error
	GetListsByUId(ctx context.Context, UId int) ([]domain.TaskList, error)
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

	return r
}

func (t tasks) createList(w http.ResponseWriter, r *http.Request) {
	var list domain.TaskList

	defer func() { _ = r.Body.Close() }()
	if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
		http.Error(w, fmt.Sprintf("error unmarshalling JSON: %s", err), http.StatusBadRequest)
		return
	}
	if err := t.service.CreateList(context.Background(), &list); err != nil {
		http.Error(w, fmt.Sprintf("error adding data to db: %s", err), http.StatusBadRequest)
		return
	}
}

func (t tasks) getListsByUId(w http.ResponseWriter, r *http.Request) {
	sUId := r.URL.Query().Get("UId")
	defer func() { _ = r.Body.Close() }()

	UId, err := strconv.Atoi(sUId)
	if sUId == "" || err != nil {
		http.Error(w, "No user with such id! Provide it via url parameter 'UId'", http.StatusBadRequest)
		return
	}
	if lists, err := t.service.GetListsByUId(context.TODO(), UId); err == nil {
		if lists == nil {
			http.Error(w, "No user with such id!", http.StatusBadRequest)
			return
		}
		for i, list := range lists {
			fmt.Fprintf(w, "%d list: %v\n", i, list)
		}
	} else {
		http.Error(w, fmt.Sprintf("error receiving lists by id: %s", err), http.StatusInternalServerError)
	}
}
