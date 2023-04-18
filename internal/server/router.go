package server

import (
	"io/fs"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gabe565/matrimony/internal/server/handlers"
	"github.com/gabe565/matrimony/internal/server/handlers/rsvp"
	middleware2 "github.com/gabe565/matrimony/internal/server/middleware"
	"github.com/gabe565/matrimony/internal/server/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"github.com/go-chi/render"
	"gorm.io/gorm"
)

func Router(db *gorm.DB, rootFs fs.FS, dataFs fs.FS) *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Heartbeat("/api/health"))
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.CleanPath)

	fileserver := http.FileServer(http.FS(rootFs))

	r.Get("/", handlers.GetIndex(rootFs))
	r.With(render.SetContentType(render.ContentTypeJSON)).
		Get("/manifest.json", handlers.GetManifest(rootFs))

	// Serve index as 404
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/"
		handlers.GetIndex(rootFs).ServeHTTP(w, r)
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
			r.Use(middleware2.PrivateEventAuth)

			r.Group(func(r chi.Router) {
				r.Use(httprate.LimitByIP(60, time.Minute))
				r.Get("/config", handlers.GetConfig)
				r.Get("/info", handlers.GetInfo)
				r.Get("/sections", handlers.ListSections)
				r.Get("/privacy", handlers.GetPrivacy)
				r.Get("/partners", handlers.ListPartners)
				r.Get("/party", handlers.ListParty)
				r.Get("/moments", handlers.ListMoments(dataFs))
				r.Get("/ical/{section}/{key}", handlers.GetIcal())
				r.Get("/rsvp/questions", rsvp.ListQuestions)
			})

			r.Group(func(r chi.Router) {
				r.Use(httprate.LimitByIP(15, time.Minute))
				r.Get("/rsvp/init", rsvp.Init(db))
				r.Group(func(r chi.Router) {
					r.Put("/rsvp/response", rsvp.Respond(db))
					r.Get("/rsvp/responses", rsvp.GetResponses(db))
					r.Post("/rsvp/guest/add", rsvp.CreateGuest(db))
				})
			})
		})

		r.Group(func(r chi.Router) {
			r.Use(httprate.LimitByIP(10, time.Minute))
			r.Use(middleware2.AdminAuth)

			r.Put("/info", handlers.UpdateInfo)
			r.Put("/sections", handlers.UpdateSections)
			r.Put("/privacy", handlers.UpdatePrivacy)
			r.Put("/partners", handlers.UpdatePartners)
			r.Put("/party", handlers.UpdateParty)

			r.Get("/guest", handlers.ListGuests(db))
			r.Get("/guest/{id}", handlers.GetGuest(db))
			r.Post("/guest", handlers.CreateGuest(db))
			r.Put("/guest/{id}", handlers.UpdateGuest(db))
			r.Delete("/guest/{id}", handlers.DeleteGuest(db))
		})

		r.HandleFunc("/*", func(w http.ResponseWriter, r *http.Request) {
			if err := render.Render(w, r, models.ErrNotFound); err != nil {
				panic(err)
			}
		})
	})

	return r
}
