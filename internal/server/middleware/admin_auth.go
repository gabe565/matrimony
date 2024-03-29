package middleware

import (
	"net/http"

	"github.com/gabe565/matrimony/internal/config"
	"github.com/gabe565/matrimony/internal/config/models"
)

func AdminAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var success bool
		switch config.Config.Admin.Auth.Method {
		case models.AuthMethodBasic:
			success = adminBasicAuth(r)
		case models.AuthMethodHeader:
			success = adminHeaderAuth(r)
		}

		if success {
			next.ServeHTTP(w, r)
			return
		}

		w.Header().Add("WWW-Authenticate", `Basic realm="Admin"`)
		w.WriteHeader(http.StatusUnauthorized)
	})
}

func adminBasicAuth(r *http.Request) bool {
	user, pass, ok := r.BasicAuth()
	if !ok {
		return false
	}

	for _, admin := range config.Config.Admin.Users {
		if user == admin.Username && admin.CheckPasswordHash(pass) {
			return true
		}
	}

	return false
}

func adminHeaderAuth(r *http.Request) bool {
	header := config.Config.Admin.Auth.Header
	user := r.Header.Get(header)
	if user == "" {
		return false
	}

	for _, admin := range config.Config.Admin.Users {
		if user == admin.Username {
			return true
		}
	}

	return false
}
