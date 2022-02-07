package handlers

import (
	"encoding/json"
	"github.com/gabe565/matrimony/internal/config"
	"net/http"
)

func ListParty(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal(config.Config.Party.Members)
	if err != nil {
		panic(err)
	}

	_, err = w.Write(j)
	if err != nil {
		return
	}
}

func PutParty(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&config.Config.Party.Members)
	if err != nil {
		panic(err)
	}

	err = config.SaveConfig()
	if err != nil {
		panic(err)
	}

	ListParty(w, r)
}
