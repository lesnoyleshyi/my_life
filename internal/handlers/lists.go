package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"my_life/internal/domain"
	"net/http"
	"strconv"
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
	sUId := r.URL.Query().Get("UId")
	defer func() { _ = r.Body.Close() }()

	UId, err := strconv.Atoi(sUId)
	if sUId == "" || err != nil {
		http.Error(w, "No user with such id! Provide it via url parameter 'UId'", http.StatusBadRequest)
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
