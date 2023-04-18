package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gabe565/matrimony/internal/config"
)

func ListSections(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal(config.Config.Sections)
	if err != nil {
		panic(err)
	}

	_, err = w.Write(j)
	if err != nil {
		return
	}
}

func UpdateSections(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&config.Config.Sections)
	if err != nil {
		panic(err)
	}

	err = config.SaveConfig()
	if err != nil {
		panic(err)
	}

	ListSections(w, r)
}
