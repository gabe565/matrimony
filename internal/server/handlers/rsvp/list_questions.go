package rsvp

import (
	"github.com/gabe565/matrimony/internal/config"
	configModels "github.com/gabe565/matrimony/internal/config/models"
	"github.com/go-chi/render"
	"net/http"
)

type ListQuestionsResponse []configModels.Questions

func (ListQuestionsResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func ListQuestions(w http.ResponseWriter, r *http.Request) {
	response := ListQuestionsResponse(config.Config.RSVP.Questions)
	if err := render.Render(w, r, response); err != nil {
		return
	}
}
