package handlers

import (
	"encoding/json"
	"github.com/gabe565/matrimony/internal/config"
	"net/http"
)

func GetConfig(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal(config.Config)
	if err != nil {
		panic(err)
	}

	_, err = w.Write(j)
	if err != nil {
		return
	}
}
