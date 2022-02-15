package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gabe565/matrimony/internal/database/models"
	"github.com/go-chi/chi/v5"
	"github.com/mattn/go-sqlite3"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
)

func ListGuests(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var guests []models.Guest
		err := db.Preload(clause.Associations).Find(&guests).Error
		if err != nil {
			panic(err)
		}

		j, err := json.Marshal(guests)
		if err != nil {
			panic(err)
		}

		_, err = w.Write(j)
		if err != nil {
			return
		}
	}
}

func GetGuest(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		var guest models.Guest
		err := db.Preload(clause.Associations).First(&guest, id).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
			panic(err)
		}

		j, err := json.Marshal(guest)
		if err != nil {
			panic(err)
		}

		_, err = w.Write(j)
		if err != nil {
			return
		}
	}
}

func CreateGuest(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var guest models.Guest
		err := json.NewDecoder(r.Body).Decode(&guest)
		if err != nil {
			panic(err)
		}

		err = db.Create(&guest).Error
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

		j, err := json.Marshal(guest)
		if err != nil {
			panic(err)
		}

		_, err = w.Write(j)
		if err != nil {
			return
		}
	}
}

func UpdateGuest(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		var guest models.Guest
		err := db.First(&guest, id).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
			panic(err)
		}

		err = json.NewDecoder(r.Body).Decode(&guest)
		if err != nil {
			panic(err)
		}

		err = db.Save(&guest).Error
		if err != nil {
			panic(err)
		}

		j, err := json.Marshal(guest)
		if err != nil {
			panic(err)
		}

		_, err = w.Write(j)
		if err != nil {
			return
		}
	}
}

func DeleteGuest(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		var guest models.Guest
		err := db.First(&guest, id).Error
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}
			panic(err)
		}

		j, err := json.Marshal(guest)
		if err != nil {
			panic(err)
		}

		err = db.Delete(&guest).Error
		if err != nil {
			panic(err)
		}

		_, err = w.Write(j)
		if err != nil {
			return
		}
	}
}
