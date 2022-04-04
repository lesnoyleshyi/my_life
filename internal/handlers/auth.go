package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"my_life/internal/domain"
	"net/http"
)

func (h *handler) signUp(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Fprintf(w, "error inmarshalling json: %v", err)
		return
	}

	userId, err := h.services.CreateUser(context.TODO(), &user)
	if err != nil {
		fmt.Fprintf(w, "error creating user: %v", err)
	} else {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "User with id = %d successfully created", userId)
	}
}

func (h *handler) signIn(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "this functionality will be implemented very soon!")
}
