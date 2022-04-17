package handlers

import (
	"encoding/json"
)

type Response struct {
	Success bool            `json:"success"`
	Body    json.RawMessage `json:"body,omitempty"`
	ErrCode int             `json:"errCode,omitempty"`
	ErrMsg  string          `json:"errMsg,omitempty"`
}

func errResponse(r *Response, errCode int, errMsg string) {
	r.Success = false
	r.ErrCode = errCode
	r.ErrMsg = errMsg
}
