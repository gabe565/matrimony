package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/arran4/golang-ical"
	"github.com/gabe565/matrimony/internal/config"
	"github.com/gabe565/matrimony/internal/server/helpers"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
)

var (
	ErrInvalidSection = errors.New("invalid section")
	ErrInvalidEvent   = errors.New("invalid event")
)

func GetIcal() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sectionKey, err := strconv.Atoi(chi.URLParam(r, "section"))
		if err != nil {
			panic(err)
		}

		if sectionKey > len(config.Config.Sections) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		section := config.Config.Sections[sectionKey]
		if section.Schedule == nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		eventKey, err := strconv.Atoi(chi.URLParam(r, "key"))
		if err != nil {
			panic(err)
		}

		if eventKey > len(section.Schedule.Events) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		event := section.Schedule.Events[eventKey]

		cal := ics.NewCalendar()
		cal.SetMethod(ics.MethodRequest)
		calEvent := cal.AddEvent(uuid.New().String())
		calEvent.SetCreatedTime(time.Now())
		calEvent.SetDtStampTime(time.Now())
		calEvent.SetModifiedAt(time.Now())

		if event.Start == nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		calEvent.SetStartAt(*event.Start)
		calEvent.SetEndAt((*event.Start).Add(time.Minute * 30))

		if event.Location != "" {
			calEvent.SetLocation(event.Location)
		}

		summary := fmt.Sprintf("%s - %s", config.Config.TitleShort(), event.Title)
		calEvent.SetSummary(summary)
		calEvent.SetDescription(event.Description)

		calEvent.SetURL(helpers.PublicUrl(r))

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
