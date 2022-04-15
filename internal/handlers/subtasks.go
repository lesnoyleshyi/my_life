package handlers

import (
	"encoding/json"
	"fmt"
	"my_life/internal/domain"
	"my_life/internal/services"
	"net/http"
)

type subtaskHandler struct {
	services *services.Service
}

func newSubtaskHandler(services *services.Service) *subtaskHandler {
	return &subtaskHandler{
		services: services,
	}
}

func (h subtaskHandler) createSubtask(w http.ResponseWriter, r *http.Request) {
	var subtask domain.Subtask

	if err := json.NewDecoder(r.Body).Decode(&subtask); err != nil {
		http.Error(w, fmt.Sprintf("error unmarshalling json: %s", err), http.StatusBadRequest)
		return
	}

	if err := h.services.CreateSubTask(r.Context(), &subtask); err != nil {
		http.Error(w, fmt.Sprintf("error adding data to db: %s", err), http.StatusBadRequest)
		return
	}

	UId, ok := r.Context().Value("UId").(int32)
	if ok {
		fmt.Fprintf(w, "New subtask for user with id=%domain added\n", UId)
	} else {
		http.Error(w, fmt.Sprintf("can't get UId from request or convert it int32"), http.StatusBadRequest)
	}
}

func (h subtaskHandler) getSubTasksByUId(w http.ResponseWriter, r *http.Request) {
	UId, ok := r.Context().Value("UId").(int32)
	if !ok {
		http.Error(w, fmt.Sprintf("can't retreive user id from context"), http.StatusBadRequest)
		return
	}

	subtasks, err := h.services.GetSubTasksByUId(r.Context(), UId)
	if err != nil {
		http.Error(w, fmt.Sprintf("can't retreive tasks from db: %s", err), http.StatusBadRequest)
		return
	}
	for _, subtask := range subtasks {
		subtaskJson, err := json.Marshal(subtask)
		if err != nil {
			http.Error(w, fmt.Sprintf("can't marshall struct to JSON: %s", err), http.StatusInternalServerError)
			return
		}
		if _, err := fmt.Fprintln(w, string(subtaskJson)); err != nil {
			http.Error(w, fmt.Sprintf("can't write JSON to response body: %s", err), http.StatusInternalServerError)
			return
		}
	}
}
