package models

import (
	"github.com/go-chi/render"
	"net/http"
)

var ErrUnauthorized = CreateErrorFromStatusCode(http.StatusUnauthorized)
var ErrNotFound = CreateErrorFromStatusCode(http.StatusNotFound)

type ErrorResponse struct {
	Message    string `json:"error"`
	StatusCode int    `json:"-"`
}

func (e ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	if e.StatusCode == 0 {
		e.StatusCode = 500
	}
	render.Status(r, e.StatusCode)
	return nil
}

func CreateErrorFromStatusCode(statusCode int) ErrorResponse {
	return ErrorResponse{
		Message:    http.StatusText(statusCode),
		StatusCode: statusCode,
	}
}

func (e ErrorResponse) Error() string {
	return e.Message
}
