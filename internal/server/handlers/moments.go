package handlers

import (
	"encoding/json"
	"io/fs"
	"net/http"
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
			if !file.IsDir() {
				response = append(response, "/public/moments/"+file.Name())
			}
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
