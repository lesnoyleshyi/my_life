package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"my_life/internal/domain"
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
	response := domain.Response{Success: true}

	UId, ok := r.Context().Value("UId").(int32)
	if !ok {
		http.Error(w, "kaka", http.StatusUnauthorized)
		return
	}

	replLists, err := h.services.GetFullUserInfo(r.Context(), UId)
	if err != nil {
		http.Error(w, fmt.Sprintf("error receiveing tasks: %v", err), http.StatusInternalServerError)
	}

	byteLists, err := json.Marshal(replLists)
	if err != nil {
		response.Success = false
		response.ErrCode = http.StatusInternalServerError
		response.ErrMsg = "marshalling problems"
	}
	response.Body = string(byteLists)
	resp, err := json.Marshal(response)
	if err != nil {
		log.Printf("can't marshall task to JSON: %v", err)
		response.Success = false
		response.ErrCode = http.StatusInternalServerError
		response.ErrMsg = "marshalling problems"
	}
	if _, err := fmt.Fprintln(w, string(resp)); err != nil {
		http.Error(w, "can't write JSON to body", http.StatusInternalServerError)
	}
}
