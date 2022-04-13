package handlers

import (
	"encoding/json"
	"fmt"
	"my_life/internal/domain"
	"my_life/internal/services"
	"net/http"
)

type taskHandler struct {
	services *services.Service
}

func newTaskHandler(services *services.Service) *taskHandler {
	return &taskHandler{
		services: services,
	}
}

func (h taskHandler) createTask(w http.ResponseWriter, r *http.Request) {
	var task domain.Task

	defer func() { _ = r.Body.Close() }()

	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, fmt.Sprintf("error unmarshalling json: %s", err), http.StatusBadRequest)
		return
	}

	if err := h.services.CreateTask(r.Context(), &task); err != nil {
		http.Error(w, fmt.Sprintf("error adding data to db: %s", err), http.StatusBadRequest)
		return
	}

	UId, ok := r.Context().Value("UId").(int32)
	if ok {
		fmt.Fprintf(w, "New task for user with id=%d added\n", UId)
	} else {
		http.Error(w, fmt.Sprintf("can't get UId from request or convert int to int32"), http.StatusBadRequest)
	}
}

func (h taskHandler) getTasksByUId(w http.ResponseWriter, r *http.Request) {
	UId, ok := r.Context().Value("UId").(int32)
	if !ok {
		http.Error(w, fmt.Sprintf("can't retreive user id from context"), http.StatusBadRequest)
		return
	}

	tasks, err := h.services.GetTasksByUId(r.Context(), UId)
	if err != nil {
		http.Error(w, fmt.Sprintf("can't retreive tasks from db: %s", err), http.StatusBadRequest)
	}
	for _, task := range tasks {
		taskJSON, err := json.Marshal(task)
		if err != nil {
			http.Error(w, fmt.Sprintf("can't marshall struct to JSON: %s", err), http.StatusInternalServerError)
			return
		}
		if _, err := fmt.Fprintln(w, taskJSON); err != nil {
			http.Error(w, fmt.Sprintf("can't write JSON to response body: %s", err), http.StatusInternalServerError)
			return
		}
	}

}