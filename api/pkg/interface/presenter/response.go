package presenter

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	code int
	body any
}

func NewResponse(code int, body any) *Response {
	return &Response{
		code: code,
		body: body,
	}
}

func (r *Response) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.code)
	json.NewEncoder(w).Encode(r.body)
}
