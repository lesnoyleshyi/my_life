package handlers

import (
	"encoding/json"
	"log"
	"my_life/internal/domain"
	"net/http"
)

//// Another variant of response struct
//type Response struct {
//	Success  bool                     `json:"success"`
//	Body     json.RawMessage          `json:"body,omitempty"`
//	ErrCode  int                      `json:"errCode,omitempty"`
//	ErrMsg   string                   `json:"errMsg,omitempty"`
//}

type Response struct {
	Success  bool                     `json:"success"`
	ErrCode  int                      `json:"errCode,omitempty"`
	ErrMsg   string                   `json:"errMsg,omitempty"`
	Lists    []domain.ReplTaskList    `json:"lists,omitempty"`
	Sections []domain.ReplTaskSection `json:"sections,omitempty"`
	Tasks    []domain.ReplTask        `json:"tasks,omitempty"`
	Subtasks []domain.ReplSubtask     `json:"subtasks,omitempty"`
}

func errResponse(errToLog error, errCode int, errMsg string, w http.ResponseWriter) {
	response := Response{
		Success: false,
		ErrCode: errCode,
		ErrMsg:  errMsg,
	}

	if errToLog == nil {
		log.Println(errMsg)
	} else {
		log.Println(errToLog)
	}
	w.WriteHeader(errCode)

	resp, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "unable to respond correctly", http.StatusInternalServerError)
	} else {
		if _, err := w.Write(resp); err != nil {
			http.Error(w, "unable to respond correctly", http.StatusInternalServerError)
		}
	}
}
