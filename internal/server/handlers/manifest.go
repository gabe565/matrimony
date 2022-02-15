package handlers

import (
	"encoding/json"
	"github.com/gabe565/matrimony/internal/config"
	"io/fs"
	"net/http"
	"os"
)

func GetManifest(rootFs fs.FS) http.HandlerFunc {
	type manifest struct {
		Name            string        `json:"name"`
		ShortName       string        `json:"short_name"`
		StartUrl        string        `json:"start_url"`
		ThemeColor      string        `json:"theme_color"`
		BackgroundColor string        `json:"background_color"`
		Display         string        `json:"display"`
		Icons           []interface{} `json:"icons"`
	}
	manifestPath := "manifest.json"
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := fs.Stat(rootFs, manifestPath); os.IsNotExist(err) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return

		}
		f, err := rootFs.Open(manifestPath)
		if err != nil {
			panic(err)
		}

		var m manifest
		err = json.NewDecoder(f).Decode(&m)
		if err != nil {
			panic(err)
		}

		if m.Name == "" {
			m.Name = config.Config.Title()
		}

		if m.ShortName == "" {
			m.ShortName = config.Config.TitleShort()
		}

		if m.ThemeColor == "" {
			m.ThemeColor = config.Config.Theme.Primary
		}

		if m.BackgroundColor == "" {
			m.BackgroundColor = config.Config.Theme.Primary
		}

		err = json.NewEncoder(w).Encode(m)
		if err != nil {
			panic(err)
		}
	}
}
