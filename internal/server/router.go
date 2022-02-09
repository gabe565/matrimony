package server

import (
	"github.com/gabe565/matrimony/internal/server/handlers"
	middleware2 "github.com/gabe565/matrimony/internal/server/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/go-chi/render"
	"io/fs"
	"net/http"
	"os"
	"strings"
	"time"
)

func Router(rootFs fs.FS, dataFs fs.FS) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Heartbeat("/api/health"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)

	fileserver := http.FileServer(http.FS(rootFs))

	// Serve index as 404
	r.NotFound(func(res http.ResponseWriter, req *http.Request) {
		req.URL.Path = "/"
		fileserver.ServeHTTP(res, req)
	})

	r.Get("/*", func(res http.ResponseWriter, req *http.Request) {
		requestPath := strings.TrimLeft(req.URL.Path, "/")
		if _, err := fs.Stat(rootFs, requestPath); !os.IsNotExist(err) {
			fileserver.ServeHTTP(res, req)
		} else {
			r.NotFoundHandler().ServeHTTP(res, req)
		}
	})

	r.Get("/public/*", func(res http.ResponseWriter, req *http.Request) {
		rctx := chi.RouteContext(req.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(http.FS(dataFs)))
		fs.ServeHTTP(res, req)
	})

	r.Route("/api", func(r chi.Router) {
		r.Use(render.SetContentType(render.ContentTypeJSON))

		r.Group(func(r chi.Router) {
			r.Use(httprate.LimitByIP(60, time.Minute))
			r.Use(middleware2.PrivateEventAuth)

			r.Get("/config", handlers.GetConfig)

			r.Get("/info", handlers.GetInfo)
			r.Get("/sections", handlers.ListSections)
			r.Get("/privacy", handlers.GetPrivacy)
			r.Get("/partners", handlers.ListPartners)
			r.Get("/party", handlers.ListParty)
		})

		r.Group(func(r chi.Router) {
			r.Use(httprate.LimitByIP(10, time.Minute))
			r.Use(middleware2.AdminAuth)

			r.Put("/info", handlers.PutInfo)
			r.Put("/sections", handlers.PutSections)
			r.Put("/privacy", handlers.PutPrivacy)
			r.Put("/partners", handlers.PutPartners)
			r.Put("/party", handlers.PutParty)
		})
	})

	return r
}
