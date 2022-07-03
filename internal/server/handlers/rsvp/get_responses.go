package rsvp

import (
	"encoding/json"
	"github.com/gabe565/matrimony/internal/database/models"
	"github.com/gabe565/matrimony/internal/server/helpers"
	httpModels "github.com/gabe565/matrimony/internal/server/models"
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type GetResponsesResponse struct {
	ResponseSaved map[uint]bool            `json:"responseSaved"`
	Responses     map[uint]json.RawMessage `json:"responses"`
}

func (response GetResponsesResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func GetResponses(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		partyId := strings.ToLower(r.URL.Query().Get("id"))
		sessionPassword, err := helpers.GetBearerAuthToken(r)
		if err != nil || partyId == "" || sessionPassword == "" {
			err = render.Render(w, r, httpModels.ErrUnauthorized)
			if err != nil {
				panic(err)
			}
			return
		}

		var party models.Party
		err = db.
			Where("id = ? and session_password = ?", partyId, sessionPassword).
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

		response := GetResponsesResponse{
			ResponseSaved: make(map[uint]bool),
			Responses:     make(map[uint]json.RawMessage),
		}

		// Find guests in party
		var guests []models.Guest
		err = db.Select(
			"ID",
			"FirstName",
			"LastName",
			"RSVP",
		).
			Where("party_id = ?", partyId).
			Find(&guests).
			Error
		if err != nil {
			panic(err)
		}

		for _, g := range guests {
			if g.RSVP != nil {
				response.ResponseSaved[g.ID] = true
				response.Responses[g.ID] = json.RawMessage(g.RSVP)
			}
		}

		err = render.Render(w, r, response)
		if err != nil {
			panic(err)
		}
	}
}
