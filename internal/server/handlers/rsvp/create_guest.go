package rsvp

import (
	"github.com/gabe565/matrimony/internal/database/models"
	httpModels "github.com/gabe565/matrimony/internal/server/models"
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"net/http"
)

type CreateGuestRequest struct {
	PartyID         uint   `json:"partyId"`
	SessionPassword string `json:"sessionPassword"`
	Guest           guest  `json:"guest"`
}

func (request CreateGuestRequest) Bind(r *http.Request) error {
	if request.PartyID == 0 || request.SessionPassword == "" {
		return httpModels.ErrUnauthorized
	}
	return nil
}

func CreateGuest(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var body CreateGuestRequest
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

		var party models.Party
		err = db.
			Where("id = ? and session_password = ?", body.PartyID, body.SessionPassword).
			Find(&party).Error
		if err != nil {
			panic(err)
		}

		if party.ID == 0 {
			err = render.Render(w, r, httpModels.ErrUnauthorized)
			if err != nil {
				panic(err)
			}
			return
		}

		body.Guest.PartyID = party.ID
		err = db.Table("guests").Create(&body.Guest).Error
		if err != nil {
			panic(err)
		}

		err = render.Render(w, r, body.Guest)
		if err != nil {
			panic(err)
		}
	}
}
