package handlers

import (
	"encoding/json"
	"github.com/gabe565/matrimony/internal/config"
	"net/http"
)

func GetInfo(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal(config.Config.EventInfo)
	if err != nil {
		panic(err)
	}

	_, err = w.Write(j)
	if err != nil {
		return
	}
}

func UpdateInfo(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&config.Config.EventInfo)
	if err != nil {
		panic(err)
	}

	err = config.SaveConfig()
	if err != nil {
		panic(err)
	}

	GetInfo(w, r)
}
