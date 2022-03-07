package rsvp

import (
	"encoding/json"
	"github.com/gabe565/matrimony/internal/database/models"
	httpModels "github.com/gabe565/matrimony/internal/server/models"
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type guest struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	PartyID   uint   `json:"partyId"`
}

func (guest) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type InitResponse struct {
	Guests          []guest `json:"guests"`
	ID              uint    `json:"id"`
	SessionPassword string  `json:"sessionPassword"`
	HeadId          uint    `json:"headId"`
}

func (InitResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func Init(db *gorm.DB) http.HandlerFunc {
	ErrUserNotFound := httpModels.ErrorResponse{
		Message:    `No match was found. Check your spelling and try again. "David" vs "Dave"`,
		StatusCode: http.StatusNotFound,
	}

	ErrInvalidEmail := httpModels.ErrorResponse{
		Message:    `We found you, but the email address provided does not match the one we have on file. Please try again with a different email address.`,
		StatusCode: http.StatusBadRequest,
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		first := strings.ToLower(r.URL.Query().Get("first"))
		last := strings.ToLower(r.URL.Query().Get("last"))
		email := strings.ToLower(r.URL.Query().Get("email"))

		// First name is required
		if first == "" {
			err = render.Render(w, r, ErrUserNotFound)
			if err != nil {
				panic(err)
			}
			return
		}

		var queryGuest models.Guest
		err = db.Preload("Party").
			Where("lower(first_name) = ? and lower(last_name) = ?", first, last).
			Find(&queryGuest).Error
		if err != nil {
			panic(err)
		}

		// Guest not found
		if queryGuest.ID == 0 {
			err = render.Render(w, r, ErrUserNotFound)
			if err != nil {
				panic(err)
			}
			return
		}

		// Email does not match
		if queryGuest.EmailAddress != nil && strings.ToLower(*queryGuest.EmailAddress) != email {
			err = render.Render(w, r, ErrInvalidEmail)
			if err != nil {
				panic(err)
			}
			return
		}

		// Save email
		if email != "" {
			queryGuest.EmailAddress = &email

			err = db.Select("EmailAddress").Save(&queryGuest).Error
			if err != nil {
				panic(err)
			}
		}

		// Generate session password
		if queryGuest.Party.SessionPassword == "" {
			err = queryGuest.Party.GenerateSessionPassword()
			if err != nil {
				panic(err)
			}
			err = db.Select("SessionPassword").Save(&queryGuest.Party).Error
		}

		// Find guests in party
		var guests []models.Guest
		err = db.Select(
			"ID",
			"FirstName",
			"LastName",
			"RSVP",
		).
			Where("party_id = ?", queryGuest.PartyID).
			Find(&guests).
			Error
		if err != nil {
			panic(err)
		}

		// Build response
		response := InitResponse{
			ID:              queryGuest.PartyID,
			SessionPassword: queryGuest.Party.SessionPassword,
		}

		var headId uint
		for _, g := range guests {
			var last string
			if g.LastName != nil {
				last = *g.LastName
			}

			var rsvp map[string]interface{}
			err = json.Unmarshal(g.RSVP, &rsvp)

			response.Guests = append(response.Guests, guest{
				ID:        g.ID,
				FirstName: g.FirstName,
				LastName:  last,
			})

			if g.ID < headId || headId == 0 {
				headId = g.ID
			}
		}
		response.HeadId = headId

		err = render.Render(w, r, response)
		if err != nil {
			panic(err)
		}
		return
	}
}
