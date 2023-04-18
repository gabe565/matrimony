package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gabe565/matrimony/internal/config"
	"github.com/gabe565/matrimony/internal/config/models"
)

func ListParty(w http.ResponseWriter, r *http.Request) {
	var section *models.PartySection
	for _, s := range config.Config.Sections {
		if s.Party != nil {
			section = s.Party
		}
	}
	if section == nil {
		w.WriteHeader(404)
	}
	j, err := json.Marshal(section)
	if err != nil {
		panic(err)
	}

	_, err = w.Write(j)
	if err != nil {
		return
	}
}

func UpdateParty(w http.ResponseWriter, r *http.Request) {
	var section *models.PartySection
	for _, s := range config.Config.Sections {
		if s.Party != nil {
			section = s.Party
		}
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(section)
	if err != nil {
		panic(err)
	}

	err = config.SaveConfig()
	if err != nil {
		panic(err)
	}

	ListParty(w, r)
}
