package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gabe565/matrimony/internal/config"
	"github.com/gabe565/matrimony/internal/database/models"
	"github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
	"net/http"
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
	HasRSVP   bool   `json:"hasRSVP" gorm:"-"`
}

func InitRSVP(db *gorm.DB) http.HandlerFunc {
	type response struct {
		Guests          []guest `json:"guests"`
		ID              uint    `json:"id"`
		SessionPassword string  `json:"sessionPassword"`
		HeadId          uint    `json:"headId"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var queryGuest models.Guest
		queryGuest.FirstName = r.URL.Query().Get("first")
		last := r.URL.Query().Get("last")
		queryGuest.LastName = &last
		email := r.URL.Query().Get("email")

		if queryGuest.FirstName == "" || queryGuest.LastName == nil || *queryGuest.LastName == "" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		err := db.Preload("Party").
			Where(&queryGuest, "FirstName", "LastName").
			Find(&queryGuest).Error
		if err != nil {
			panic(err)
		}

		if queryGuest.ID != 0 {
			// Save email
			if email != "" {
				queryGuest.EmailAddress = &email
			}

			err = db.Select("EmailAddress", "SessionPassword").Save(&queryGuest).Error
			if err != nil {
				panic(err)
			}

			// Generate session password
			err = queryGuest.Party.GenerateSessionPassword()
			if err != nil {
				panic(err)
			}
			err = db.Select("SessionPassword").Save(&queryGuest.Party).Error

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
				ID:              queryGuest.ID,
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
				_, hasRsvp := rsvp["confirmed"]

				response.Guests = append(response.Guests, guest{
					ID:        g.ID,
					FirstName: g.FirstName,
					LastName:  last,
					HasRSVP:   hasRsvp,
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
		} else {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
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
			sqliteErr, ok := err.(sqlite3.Error)
			if ok {
				if sqliteErr.Code == sqlite3.ErrConstraint {
					http.Error(w, err.Error(), 400)
					return
				}
			}
			panic(err)
			return
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
