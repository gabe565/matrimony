package middleware

import (
	"context"
	"encoding/json"
	"github.com/gabe565/matrimony/internal/database/models"
	"gorm.io/gorm"
	"net/http"
)

type ContextKey string

const GuestKey = ContextKey("guest")

func GuestRSVPAuth(db *gorm.DB) func(http.Handler) http.Handler {
	type Body struct {
		ID              uint   `json:"id"`
		SessionPassword string `json:"sessionPassword"`
	}

	checkAuth := func(r *http.Request) (*models.Guest, error) {
		var body Body
		err := json.NewDecoder(r.Body).Decode(&body)
		if err != nil {
			panic(err)
		}

		guest := models.Guest{
			Model: models.Model{
				ID: body.ID,
			},
			Party: models.Party{
				SessionPassword: body.SessionPassword,
			},
		}
		var found bool
		err = db.Model(models.Guest{}).
			Select("count(*) > 1").
			Where(&guest, "ID", "SessionPassword").
			Find(&found).Error
		if err != nil {
			return nil, err
		}

		return &guest, nil
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			guest, err := checkAuth(r)
			if err != nil {
				panic(err)
			}

			if guest != nil {
				ctx := context.WithValue(r.Context(), GuestKey, guest)
				next.ServeHTTP(w, r.WithContext(ctx))
				return
			}

			w.WriteHeader(http.StatusUnauthorized)
			return
		})
	}
}
