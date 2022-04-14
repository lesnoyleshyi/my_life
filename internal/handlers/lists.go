package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"my_life/internal/domain"
	"my_life/internal/services"
	"net/http"
)

type listHandler struct {
	services *services.Service
}

func newListHandler(services *services.Service) *listHandler {
	return &listHandler{
		services: services,
	}
}

func (h listHandler) createList(w http.ResponseWriter, r *http.Request) {
	var list domain.TaskList

	defer func() { _ = r.Body.Close() }()
	if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
		http.Error(w, fmt.Sprintf("error unmarshalling JSON: %s", err), http.StatusBadRequest)
		return
	}
	if err := h.services.CreateList(context.Background(), &list); err != nil {
		http.Error(w, fmt.Sprintf("error adding data to db: %s", err), http.StatusBadRequest)
		return
	}
}

func (h listHandler) getListsByUId(w http.ResponseWriter, r *http.Request) {

	UId, ok := r.Context().Value("UId").(int32)
	if !ok {
		http.Error(w, "something went wrong when parsing user id", http.StatusBadRequest)
		return
	}
	if lists, err := h.services.GetListsByUId(r.Context(), UId); err == nil {
		if lists == nil {
			http.Error(w, "No lists for user with such id!", http.StatusBadRequest)
			return
		}
		for i, list := range lists {
			fmt.Fprintf(w, "%d list: %v\n", i, list)
		}
	} else {
		http.Error(w, fmt.Sprintf("error receiving lists by id: %s", err), http.StatusInternalServerError)
	}
}
