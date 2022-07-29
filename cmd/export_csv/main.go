package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gabe565/matrimony/internal/database"
	"github.com/gabe565/matrimony/internal/database/models"
	flag "github.com/spf13/pflag"
	"gorm.io/gorm"
	"log"
	"os"
	"strings"
	"time"
)

var ErrInvalidFieldNum = errors.New("invalid number of fields")

func main() {
	start := time.Now()

	filename := flag.StringP("filename", "f", "./data/export.csv", "Guest list csv")
	flag.Parse()

	f, err := os.OpenFile(*filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	db, err := database.Setup()
	if err != nil {
		log.Fatalln(err)
	}

	w := csv.NewWriter(f)

	header := []string{
		"First Name",
		"Last Name",
		"Email",
		"Tags",
		"Party ID",
		"RSVP",
		"Number attending",
		"Leave a note!",
		"How do you know...",
		"Honeymoon",
		"Mail invite",
		"Mailing address",
		"Create date",
		"Update date",
	}
	err = w.Write(header)
	if err != nil {
		log.Fatalln(err)
	}

	var count uint
	var guests []models.Guest
	err = db.Model(models.Guest{}).
		Preload("Tags").
		Preload("Party", "ID").
		Order("party_id").
		Find(&guests).Error
	if err != nil {
		log.Fatalln(err)
	}

	for _, guest := range guests {
		fields, err := guestFields(db, guest)
		if err != nil {
			log.Fatalln(err)
		}

		if len(fields) != len(header) {
			log.Fatalln(fmt.Errorf(
				"%w. expected %d, got %d. %v",
				ErrInvalidFieldNum,
				len(header),
				len(fields), fields),
			)
		}

		err = w.Write(fields)
		if err != nil {
			log.Fatalln(err)
		}

		count += 1
	}

	w.Flush()
	err = w.Error()
	if err != nil {
		log.Fatalln(err)
	}

	err = f.Close()
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Exported %d rows in %v\n", count, time.Since(start).Round(time.Millisecond))
}

func guestFields(db *gorm.DB, guest models.Guest) ([]string, error) {
	var err error
	r := []string{
		guest.FirstName,
	}

	r = append(r, dereferenceOrEmpty[string](guest.LastName))
	r = append(r, dereferenceOrEmpty[string](guest.EmailAddress))

	var tagNames strings.Builder
	for i, tag := range guest.Tags {
		if i > 0 {
			tagNames.WriteString(", ")
		}
		tagNames.WriteString(tag.Name)
	}
	r = append(r, tagNames.String())

	r = append(r, fmt.Sprintf("%d", guest.Party.ID))

	var rsvp map[string]any
	if guest.RSVP != nil {
		err = json.Unmarshal(guest.RSVP, &rsvp)
		if err != nil {
			return r, err
		}
	}

	r = append(r, castOrEmpty[string](rsvp["RSVP"]))
	switch rsvp["Number Attending"].(type) {
	case float64:
		s := fmt.Sprintf("%.0f", castOrEmpty[float64](rsvp["Number Attending"]))
		r = append(r, s)
	case string:
		r = append(r, castOrEmpty[string](rsvp["Number Attending"]))
	default:
		r = append(r, "")
	}
	r = append(r, castOrEmpty[string](rsvp["Note To The Couple"]))
	r = append(r, castOrEmpty[string](rsvp["How do you know the couple"]))
	r = append(r, castOrEmpty[string](rsvp["Honeymoon"]))
	r = append(r, castOrEmpty[string](rsvp["Mail Invitation"]))
	r = append(r, castOrEmpty[string](rsvp["Mailing Address"]))
	r = append(r, guest.CreatedAt.Format("2006-01-02 15:04:05"))
	r = append(r, guest.UpdatedAt.Format("2006-01-02 15:04:05"))

	return r, nil
}

func castOrEmpty[T any](val any) T {
	if val != nil {
		return val.(T)
	}
	var empty T
	return empty
}

func dereferenceOrEmpty[T any](val *T) T {
	if val != nil {
		return *val
	}
	var empty T
	return empty
}
