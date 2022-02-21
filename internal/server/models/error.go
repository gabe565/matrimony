package models

import (
	"github.com/go-chi/render"
	"net/http"
)

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
