package handlers

import (
	"encoding/json"
	"io/fs"
	"net/http"
	"strings"
)

const MomentsDir = "moments"

func ListMoments(fsys fs.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		files, err := fs.ReadDir(fsys, MomentsDir)
		if err != nil {
			panic(err)
		}

		var response []string
		for _, file := range files {
			lowerFilename := strings.ToLower(file.Name())
			if file.IsDir() || strings.HasPrefix(lowerFilename, ".") {
				continue
			}
			response = append(response, file.Name())
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
