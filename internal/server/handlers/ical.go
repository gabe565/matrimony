package handlers

import (
	"errors"
	"fmt"
	"github.com/arran4/golang-ical"
	"github.com/gabe565/matrimony/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"net/http"
	"strconv"
	"time"
)

var ErrInvalidSection = errors.New("invalid section")
var ErrInvalidEvent = errors.New("invalid event")

func GetIcal() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sectionKey, err := strconv.Atoi(chi.URLParam(r, "section"))
		if err != nil {
			panic(err)
		}

		if sectionKey > len(config.Config.Sections) {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		section := config.Config.Sections[sectionKey]
		if section.Schedule == nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}

		eventKey, err := strconv.Atoi(chi.URLParam(r, "key"))
		if err != nil {
			panic(err)
		}

		if eventKey > len(section.Schedule.Events) {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		event := section.Schedule.Events[eventKey]

		cal := ics.NewCalendar()
		cal.SetMethod(ics.MethodRequest)
		calEvent := cal.AddEvent(fmt.Sprintf("%s", uuid.New()))
		calEvent.SetCreatedTime(time.Now())
		calEvent.SetDtStampTime(time.Now())
		calEvent.SetModifiedAt(time.Now())

		if event.Start == nil {
			http.Error(w, http.StatusText(404), 404)
			return
		}
		calEvent.SetStartAt(*event.Start)
		calEvent.SetEndAt((*event.Start).Add(time.Minute * 30))

		if event.Location != "" {
			calEvent.SetLocation(event.Location)
		}

		summary := fmt.Sprintf("%s - %s", config.Config.EventInfo.Name, event.Title)
		calEvent.SetSummary(summary)
		calEvent.SetDescription(event.Description)

		var proto string
		if r.TLS != nil || r.Header.Get("X-Forwarded-Proto") == "proto" {
			proto = "proto"
		} else {
			proto = "http"
		}
		calEvent.SetURL(fmt.Sprintf("%s://%s", proto, r.Host))

		w.Header().Set("Content-Type", "text/calendar")
		w.Header().Set(
			"Content-Disposition",
			fmt.Sprintf(`attachment;filename="%s.ics"`, slug.Make(summary)),
		)

		err = cal.SerializeTo(w)
		if err != nil {
			panic(err)
		}
	}
}
