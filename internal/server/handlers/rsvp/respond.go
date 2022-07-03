package rsvp

import (
	"encoding/json"
	"github.com/gabe565/matrimony/internal/database/models"
	httpModels "github.com/gabe565/matrimony/internal/server/models"
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"net/http"
)

type RespondRequest struct {
	ID              uint           `json:"id"`
	SessionPassword string         `json:"sessionPassword"`
	Values          map[string]any `json:"values"`
}

func (request RespondRequest) Bind(r *http.Request) error {
	if request.ID == 0 || request.SessionPassword == "" {
		return httpModels.ErrUnauthorized
	}
	return nil
}

type RespondResponse struct{}

func (RespondResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func Respond(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body RespondRequest
		err := render.Bind(r, &body)
		if err != nil {
			switch err.(type) {
			case httpModels.ErrorResponse:
				err = render.Render(w, r, httpModels.ErrUnauthorized)
				if err != nil {
					panic(err)
				}
				return
			default:
				panic(err)
			}
		}

		// Find guest by given ID
		var guest models.Guest
		err = db.Preload("Party").
			Where("id = ?", body.ID).
			Find(&guest).Error
		if err != nil {
			panic(err)
		}

		// Trigger error if guest not found or invalid password
		if guest.ID == 0 || guest.Party.SessionPassword != body.SessionPassword {
			err = render.Render(w, r, httpModels.ErrUnauthorized)
			if err != nil {
				panic(err)
			}
			return
		}

		guestRSVPStr, err := guest.RSVP.Value()
		if err != nil {
			panic(err)
		}

		var guestJson map[string]any
		if guestRSVPStr != nil {
			err = json.Unmarshal([]byte(guestRSVPStr.(string)), &guestJson)
			if err != nil {
				panic(err)
			}
		} else {
			guestJson = make(map[string]any)
		}

		for k, v := range body.Values {
			guestJson[k] = v
		}

		guestJsonStr, err := json.Marshal(guestJson)
		if err != nil {
			panic(err)
		}

		err = guest.RSVP.Scan(guestJsonStr)
		if err != nil {
			panic(err)
		}
		err = db.Select("RSVP").Save(&guest).Error
		if err != nil {
			panic(err)
		}

		var response RespondResponse
		err = render.Render(w, r, response)
		if err != nil {
			panic(err)
		}
	}
}
