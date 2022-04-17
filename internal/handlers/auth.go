package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"my_life/internal/domain"
	"my_life/internal/services"
	"net/http"
)

type authHandler struct {
	services *services.Service
}

func newAuthHandler(services *services.Service) *authHandler {
	return &authHandler{
		services: services,
	}
}

type usernamePasswd struct {
	Name   string `json:"name"`
	Passwd string `json:"password"`
}

// todo нет никакой проверки того, что анмаршалится в структуру -
// если в json-е не будет передано какое-то поле,
// то соответствующее поле структуры просто получит дефолтное нулевое значение.
// Нужна валидация полей структуры.
func (h *authHandler) signUp(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		fmt.Fprintf(w, "error unmarshalling json: %v", err)
		return
	}
	defer func() { _ = r.Body.Close() }()

	userId, err := h.services.CreateUser(context.TODO(), &user)
	if err != nil {
		fmt.Fprintf(w, "error creating user: %v", err)
	} else {
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "User with id = %domain successfully created", userId)
	}
}

func (h *authHandler) signIn(w http.ResponseWriter, r *http.Request) {
	var user usernamePasswd
	var response Response

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errResponse(&response, http.StatusBadRequest, "error unmarshalling json")
		return
	}
	if user.Name == "" || user.Passwd == "" {
		w.WriteHeader(http.StatusBadRequest)
		errResponse(&response, http.StatusBadRequest, "Empty username or password")
		return
	}

	token, err := h.services.GenerateToken(context.TODO(), user.Name, user.Passwd)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errResponse(&response, http.StatusBadRequest, "error generating token")
		return
	}

	response.Body = json.RawMessage(token)
	resp, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "error marshalling response", http.StatusInternalServerError)
		return
	}
	if _, err := w.Write(resp); err != nil {
		http.Error(w, "error writing to body", http.StatusInternalServerError)
	}
}

func (h *authHandler) verifyToken(next http.Handler) http.Handler {
	return h.services.VerifyToken(next)
}
