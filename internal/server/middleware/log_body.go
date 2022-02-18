package middleware

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func LogBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		log.Println("Request body: " + string(b))
		r.Body = ioutil.NopCloser(bytes.NewReader(b))
		next.ServeHTTP(w, r)
	})
}
