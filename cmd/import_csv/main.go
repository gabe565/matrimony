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
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	filename := flag.StringP("filename", "f", "./data/guest-list.csv", "Guest list csv")
	flag.Parse()

	f, err := os.Open(*filename)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	db, err := database.Setup()
	if err != nil {
		log.Fatalln(err)
	}

	db.Exec("DELETE FROM guests")
	db.Exec("DELETE FROM guest_tags")
	db.Exec("DELETE FROM tags")
	db.Exec("DELETE FROM parties")

	r := csv.NewReader(f)
	var count uint
	var fieldNames []string
	for {
		row, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalln(err)
		}

		// Set field names
		if count == 0 {
			fieldNames = row
			count += 1
			continue
		}

		guest, err := createGuest(db, fieldNames, row)
		if err != nil {
			log.Fatalln(err)
		}

		err = db.Create(&guest).Error
		if err != nil {
			log.Fatalln(err)
		}

		count += 1
	}
	log.Printf("Imported %d rows\n", count)
}

var ErrUnknownColumn = errors.New("unknown column")

func createGuest(db *gorm.DB, fieldNames, row []string) (models.Guest, error) {
	var err error
	var guest models.Guest
	rsvpResponse := make(map[string]any)
	for key, value := range row {
		if value == "" {
			continue
		}

		switch fieldNames[key] {
		case "first name":
			guest.FirstName = value
		case "last name":
			tmp := value
			guest.LastName = &tmp
		case "email":
			tmp := value
			guest.EmailAddress = &tmp
		case "tags":
			tagReader := csv.NewReader(strings.NewReader(value))
			tags, err := tagReader.Read()
			if err != nil {
				log.Fatalln(err)
			}
			for _, tagName := range tags {
				tag, err := findOrCreateTag(db, tagName)
				if err != nil {
					return guest, err
				}
				guest.Tags = append(guest.Tags, tag)
			}
		case "party":
			party, err := findOrCreateParty(db, value)
			if err != nil {
				return guest, err
			}
			guest.Party = party
		case "rsvp":
			switch value {
			case "I'll Be There!":
				rsvpResponse["RSVP"] = "yes"
			case "Can't Make It :(":
				rsvpResponse["RSVP"] = "no"
			}
		case "leave a note!":
			rsvpResponse["Note To The Couple"] = value
		case "how do you know...":
			rsvpResponse["How do you know the couple"] = value
		case "just for fun, where should we go on our honeymoon?":
			rsvpResponse["Honeymoon"] = value
		case "number attending:":
			rsvpResponse["Number Attending"] = value
		case "mail invite":
			rsvpResponse["Mail Invitation"] = value
		case "mailing address":
			rsvpResponse["Mailing Address"] = value
		default:
			return guest, fmt.Errorf("%w: %s", ErrUnknownColumn, value)
		}
	}
	if len(rsvpResponse) > 0 {
		guest.RSVP, err = json.Marshal(rsvpResponse)
		if err != nil {
			return guest, err
		}
	}
	return guest, nil
}

func findOrCreateTag(db *gorm.DB, tagName string) (models.Tag, error) {
	var err error
	tag := models.Tag{
		Name: tagName,
	}

	err = db.Where(tag).Find(&tag).Error
	if err != nil {
		return tag, err
	}

	if tag.ID == 0 {
		err = db.Create(&tag).Error
		if err != nil {
			return tag, err
		}
	}

	return tag, nil
}

func findOrCreateParty(db *gorm.DB, partyName string) (models.Party, error) {
	var err error
	party := models.Party{
		Name: partyName,
	}

	err = db.Where(party).Find(&party).Error
	if err != nil {
		return party, err
	}

	if party.ID == 0 {
		err = db.Create(&party).Error
		if err != nil {
			return party, err
		}
	}

	return party, nil
}
