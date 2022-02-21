package models

import (
	"github.com/go-chi/render"
	"net/http"
)

var ErrUnauthorized = CreateErrorFromStatusCode(http.StatusUnauthorized)

type Error struct {
	Error      string `json:"error"`
	StatusCode int    `json:"-"`
}

func (e Error) Render(w http.ResponseWriter, r *http.Request) error {
	if e.StatusCode == 0 {
		e.StatusCode = 500
	}
	render.Status(r, e.StatusCode)
	return nil
}

func CreateErrorFromStatusCode(statusCode int) Error {
	return Error{
		Error:      http.StatusText(statusCode),
		StatusCode: statusCode,
	}
}
