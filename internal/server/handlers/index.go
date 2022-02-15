package handlers

import (
	"github.com/gabe565/matrimony/internal/config"
	"github.com/gabe565/matrimony/internal/server/helpers"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"path"
)

func GetIndex(rootFs fs.FS) http.HandlerFunc {
	indexPath := "index.html"
	return func(w http.ResponseWriter, r *http.Request) {
		if _, err := fs.Stat(rootFs, indexPath); os.IsNotExist(err) {
			http.Error(w, http.StatusText(404), 404)
			return

		}
		tpl, err := template.New(indexPath).ParseFS(rootFs, indexPath)
		if err != nil {
			panic(err)
		}

		keywords := config.Config.EventInfo.Keywords
		for _, p := range config.Config.Partners {
			keywords = append(keywords, p.Fullname())
		}

		var heroImage string
		for _, sec := range config.Config.Sections {
			if sec.Hero != nil {
				heroImage = path.Join("/", sec.Hero.Image.Src)
			}
		}

		err = tpl.Execute(w, map[string]interface{}{
			"title":       config.Config.Title(),
			"title_short": config.Config.TitleShort(),
			"description": config.Config.EventInfo.Greeting,
			"keywords":    keywords,
			"author":      config.Config.Title(),
			"heroImage":   heroImage,
			"url":         helpers.PublicUrl(r),
			"theme":       config.Config.Theme,
		})
		if err != nil {
			panic(err)
		}
	}
}
