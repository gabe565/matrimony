package handlers

import (
	"github.com/gabe565/matrimony/internal/config"
	"github.com/gabe565/matrimony/internal/config/models"
	"github.com/go-chi/render"
	"net/http"
)

type GetConfigResponse models.Config

func (GetConfigResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func GetConfig(w http.ResponseWriter, r *http.Request) {
	response := (*GetConfigResponse)(config.Config)
	if err := render.Render(w, r, response); err != nil {
		return
	}
}
