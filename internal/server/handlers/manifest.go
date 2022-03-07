package handlers

import (
	"encoding/json"
	"github.com/gabe565/matrimony/internal/config"
	"github.com/go-chi/render"
	"io/fs"
	"net/http"
	"os"
)

type ManifestResponse struct {
	Name            string        `json:"name"`
	ShortName       string        `json:"short_name"`
	StartUrl        string        `json:"start_url"`
	ThemeColor      string        `json:"theme_color"`
	BackgroundColor string        `json:"background_color"`
	Display         string        `json:"display"`
	Icons           []interface{} `json:"icons"`
}

func (ManifestResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func GetManifest(rootFs fs.FS) http.HandlerFunc {
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

		var response ManifestResponse
		err = json.NewDecoder(f).Decode(&response)
		if err != nil {
			panic(err)
		}

		if response.Name == "" {
			response.Name = config.Config.Title()
		}

		if response.ShortName == "" {
			response.ShortName = config.Config.TitleShort()
		}

		if response.ThemeColor == "" {
			response.ThemeColor = config.Config.Theme.Primary
		}

		if response.BackgroundColor == "" {
			response.BackgroundColor = config.Config.Theme.Primary
		}

		err = render.Render(w, r, response)
		if err != nil {
			panic(err)
		}
	}
}
