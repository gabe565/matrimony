package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gabe565/matrimony/internal/config"
	"github.com/gabe565/matrimony/internal/database/models"
	httpModels "github.com/gabe565/matrimony/internal/server/models"
	"github.com/go-chi/render"
	"gorm.io/gorm"
	"net/http"
	"strings"
)

func ListRSVPQuestions(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal(config.Config.RSVP.Questions)
	if err != nil {
		panic(err)
	}

	_, err = w.Write(j)
	if err != nil {
		return
	}
}

type guest struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	PartyID   uint   `json:"partyId"`
}

func InitRSVP(db *gorm.DB) http.HandlerFunc {
	type response struct {
		Guests          []guest `json:"guests"`
		ID              uint    `json:"id"`
		SessionPassword string  `json:"sessionPassword"`
		HeadId          uint    `json:"headId"`
	}

	ErrUserNotFound := httpModels.Error{
		Error:      `No match was found. Check your spelling and try again. "David" vs "Dave"`,
		StatusCode: 404,
	}

	ErrInvalidEmail := httpModels.Error{
		Error:      `We found you, but the email address provided does not match the one we have on file. Please try again with a different email address.`,
		StatusCode: 400,
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
		if queryGuest.EmailAddress != nil && *queryGuest.EmailAddress != email {
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
		response := response{
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

		j, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}

		_, err = w.Write(j)
		if err != nil {
			panic(err)
		}
		return
	}
}

var ErrInvalidSessionPassword = errors.New("invalid session password")

func RSVPResponse(db *gorm.DB) http.HandlerFunc {
	type Body struct {
		ID              uint                   `json:"id"`
		SessionPassword string                 `json:"sessionPassword"`
		Values          map[string]interface{} `json:"values"`
	}

	checkAuth := func(r *http.Request) (*models.Guest, Body, error) {
		var body Body
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			panic(err)
		}

		guest := models.Guest{
			Model: models.Model{
				ID: body.ID,
			},
		}
		err = db.Preload("Party").
			Where(&guest, "ID", "SessionPassword").
			Find(&guest).Error
		if err != nil {
			return nil, body, err
		}

		if guest.Party.SessionPassword != body.SessionPassword {
			panic(ErrInvalidSessionPassword)
		}

		return &guest, body, nil
	}

	return func(w http.ResponseWriter, r *http.Request) {
		guest, body, err := checkAuth(r)
		if err != nil {
			panic(err)
		}

		if guest.ID == 0 {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		guestRSVPStr, err := guest.RSVP.Value()
		if err != nil {
			panic(err)
		}

		var guestJson map[string]interface{}
		if guestRSVPStr != nil {
			err = json.Unmarshal([]byte(guestRSVPStr.(string)), &guestJson)
		} else {
			guestJson = make(map[string]interface{})
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

		w.WriteHeader(http.StatusOK)
	}
}

func RSVPCreateGuest(db *gorm.DB) http.HandlerFunc {
	type Body struct {
		PartyID         uint   `json:"partyId"`
		SessionPassword string `json:"sessionPassword"`
		Guest           guest  `json:"guest"`
	}

	checkAuth := func(r *http.Request) (*models.Party, Body, error) {
		var body Body
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			panic(err)
		}

		party := models.Party{
			Model: models.Model{
				ID: body.PartyID,
			},
			SessionPassword: body.SessionPassword,
		}
		err = db.Where(&party, "ID", "SessionPassword").Find(&party).Error
		if err != nil {
			return nil, body, err
		}

		return &party, body, nil
	}

	return func(w http.ResponseWriter, r *http.Request) {
		party, body, err := checkAuth(r)
		if err != nil {
			panic(err)
		}

		body.Guest.PartyID = party.ID
		err = db.Table("guests").Create(&body.Guest).Error
		if err != nil {
			panic(err)
		}

		j, err := json.Marshal(body.Guest)
		if err != nil {
			panic(err)
		}

		_, err = w.Write(j)
		if err != nil {
			return
		}
	}
}
