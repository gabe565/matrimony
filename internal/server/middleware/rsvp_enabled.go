package middleware

import (
	"net/http"

	"github.com/gabe565/matrimony/internal/config"
)

func RsvpEnabled(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !config.Config.RSVP.Enabled {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
