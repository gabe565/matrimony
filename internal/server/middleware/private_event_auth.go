package middleware

import (
	"github.com/gabe565/matrimony/internal/config"
	"github.com/gabe565/matrimony/internal/config/models"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func PrivateEventAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if config.Config.Privacy.Mode == models.PrivacyModePrivate {
			auth := map[string]string{
				"": config.Config.Privacy.Password,
			}
			middleware.BasicAuth("Private", auth)(next).ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	})
}
