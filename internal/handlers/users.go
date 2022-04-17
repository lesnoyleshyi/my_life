package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"my_life/internal/services"
	"net/http"
	"strconv"
)

type userHandler struct {
	services *services.Service
}

func newUserHandler(services *services.Service) *userHandler {
	return &userHandler{services: services}
}

func (h userHandler) getUserById(w http.ResponseWriter, r *http.Request) {
	sUId := r.URL.Query().Get("UId")

	if sUId == "" {
		http.Error(w, "Provide user id as url parameter 'UId'", http.StatusBadRequest)
		return
	}

	Uid, err := strconv.Atoi(sUId)
	if err != nil {
		http.Error(w, "unable to convert Uid to integer", http.StatusBadRequest)
		return
	}

	user, err := h.services.GetUserById(context.TODO(), int32(Uid))
	if err != nil {
		http.Error(w, fmt.Sprintf("error recieving user by this UId: %s", err), http.StatusInternalServerError)
		return
	}
	if _, err := fmt.Fprintf(w, "There is user with UId=%domain in database", user.UId); err != nil {
		http.Error(w, "Something went wrong on the server side", http.StatusInternalServerError)
	}
}

func (h userHandler) getFullUserInfo(w http.ResponseWriter, r *http.Request) {
	response := Response{Success: true}

	UId, ok := r.Context().Value("UId").(int32)
	if !ok {
		errResponse(nil, http.StatusUnauthorized, "missing or invalid token", w)
		return
	}

	replTaskLists, err := h.services.GetFullUserInfo(r.Context(), UId)
	if err != nil {
		errResponse(err, http.StatusInternalServerError, "can't retrieve data frm db", w)
		return
	}

	byteLists, err := json.Marshal(replTaskLists)
	if err != nil {
		errResponse(err, http.StatusInternalServerError, "marshalling problems", w)
		return
	}
	response.Body = json.RawMessage(byteLists)

	resp, err := json.Marshal(response)
	if err != nil {
		errResponse(err, http.StatusInternalServerError, "marshalling problems", w)
	}
	if _, err := w.Write(resp); err != nil {
		http.Error(w, "can't write JSON to body", http.StatusInternalServerError)
	}
}
