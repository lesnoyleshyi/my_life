package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"my_life/internal/domain"
	"net/http"
	"strconv"
)

func (t tasks) createUser(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "error unmarshalling JSON to struct", http.StatusBadRequest)
		return
	}

	if err := t.service.CreateUser(context.TODO(), &user); err != nil {
		http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
	} else {
		fmt.Fprintf(w, "User with id=%d added to database", user.UId)
	}
}

func (t tasks) getUserById(w http.ResponseWriter, r *http.Request) {
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

	user, err := t.service.GetUserById(context.TODO(), uint64(Uid))
	if err != nil {
		http.Error(w, fmt.Sprintf("error recieving user by this UId: %s", err), http.StatusInternalServerError)
		return
	}
	if _, err := fmt.Fprintf(w, "There is user with UId=%d", user.UId); err != nil {
		http.Error(w, "Something went wrong on the server side", http.StatusInternalServerError)
	}
}
