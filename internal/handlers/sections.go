package handlers

import (
	"encoding/json"
	"fmt"
	"my_life/internal/domain"
	"my_life/internal/services"
	"net/http"
)

type sectionHandler struct {
	services *services.Service
}

func newSectionHandler(services *services.Service) *sectionHandler {
	return &sectionHandler{
		services: services,
	}
}

func (h sectionHandler) createSection(w http.ResponseWriter, r *http.Request) {
	var section domain.TaskSection

	defer func() { _ = r.Body.Close() }()
	if err := json.NewDecoder(r.Body).Decode(&section); err != nil {
		http.Error(w, fmt.Sprintf("error unmarshalling json: %s", err), http.StatusBadRequest)
		return
	}
	if err := h.services.CreateSection(r.Context(), &section); err != nil {
		http.Error(w, fmt.Sprintf("error adding data to db: %s", err), http.StatusBadRequest)
		return
	}

	UId, ok := r.Context().Value("UId").(int32)
	if ok {
		fmt.Fprintf(w, "New section for user with id=%d added\n", UId)
	} else {
		http.Error(w, fmt.Sprintf("can't get UId from request or convert int to int32"), http.StatusBadRequest)
	}
}

func (h sectionHandler) getSectionsByUId(w http.ResponseWriter, r *http.Request) {
	UId, ok := r.Context().Value("UId").(int32)
	if !ok {
		http.Error(w, fmt.Sprintf("can't retreive user id from context"), http.StatusBadRequest)
	}

	sections, err := h.services.GetSectionsByUId(r.Context(), UId)
	if err != nil {
		http.Error(w, fmt.Sprintf("can't retreive sections from db: %s", err), http.StatusBadRequest)
	}
	for _, section := range sections {
		sectionJSON, err := json.Marshal(section)
		if err != nil {
			http.Error(w, fmt.Sprintf("can't marshall struct to JSON: %s", err), http.StatusInternalServerError)
			return
		}
		if _, err := w.Write(sectionJSON); err != nil {
			http.Error(w, fmt.Sprintf("can't write JSON to response body: %s", err), http.StatusInternalServerError)
			return
		}
		if _, err := w.Write([]byte("\n")); err != nil {
			http.Error(w, fmt.Sprintf("can't write \\n JSON to response body: %s", err), http.StatusInternalServerError)
			return
		}
	}
}
