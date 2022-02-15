package handlers

import (
	"encoding/json"
	"github.com/gabe565/matrimony/internal/config"
	"github.com/gabe565/matrimony/internal/database/models"
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

func CheckUser(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var guest models.Guest
		guest.FirstName = r.URL.Query().Get("first")
		last := r.URL.Query().Get("last")
		guest.LastName = &last
		email := r.URL.Query().Get("email")

		if guest.FirstName == "" || guest.LastName == nil || *guest.LastName == "" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		err := db.Where(&guest, "FirstName", "LastName").Find(&guest).Error
		if err != nil {
			panic(err)
		}

		if guest.ID != 0 {
			if email != "" {
				guest.EmailAddress = &email
				db.Select("EmailAddress").Save(guest)
			}
			w.WriteHeader(http.StatusNoContent)
		} else {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
	}
}
