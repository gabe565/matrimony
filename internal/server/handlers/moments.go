package handlers

import (
	"encoding/json"
	"github.com/gabe565/matrimony/internal/datadir"
	"io/fs"
	"net/http"
	"path"
	"strings"
)

const MomentsDir = "moments"
const ThumbDir = ".thumb"

func ListMoments(fsys fs.FS) http.HandlerFunc {
	type photo struct {
		Thumb string `json:"thumb"`
		Src   string `json:"src"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		files, err := fs.ReadDir(fsys, MomentsDir)
		if err != nil {
			panic(err)
		}

		var response []photo
		for _, file := range files {
			lowerFilename := strings.ToLower(file.Name())
			if file.IsDir() || strings.HasPrefix(lowerFilename, ".") {
				continue
			}
			response = append(response, photo{
				Thumb: path.Join("/", datadir.PublicDir, MomentsDir, ThumbDir, file.Name()),
				Src:   path.Join("/", datadir.PublicDir, MomentsDir, file.Name()),
			})
		}

		j, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}

		_, err = w.Write(j)
		if err != nil {
			return
		}
	}
}
