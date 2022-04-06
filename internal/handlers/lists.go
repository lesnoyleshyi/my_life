package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"my_life/internal/domain"
	"net/http"
)

func (h handler) createList(w http.ResponseWriter, r *http.Request) {
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

func (h handler) getListsByUId(w http.ResponseWriter, r *http.Request) {

	UId, ok := r.Context().Value("UId").(int)
	if !ok {
		http.Error(w, "something went wrong when parsing user id", http.StatusBadRequest)
		return
	}
	if lists, err := h.services.GetListsByUId(context.TODO(), int64(UId)); err == nil {
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
