package config

import (
	"github.com/gabe565/matrimony/internal/config/models"
	"log"
	"time"
)

func ExampleConfig() *models.Config {
	admin := models.User{
		Username: "admin",
	}
	err := admin.HashPassword("password")
	if err != nil {
		log.Fatalln(err)
	}

	return &models.Config{
		EventInfo: models.EventInfo{
			DisplayName: "Eesha & Will",
			Date:        time.Now(),
			Location:    "Oklahoma City, OK, USA",
			Timezone:    "America/Chicago",
			Greeting:    "We can't wait to share our special day with you!",
		},
		Privacy: models.Privacy{
			Password: "example1234",
		},
		Partners: []models.PartyMember{
			{
				First: "Eesha",
				Last:  "Keeling",
				Title: "Bride",
			},
			{
				First: "Will",
				Last:  "Barrow",
				Title: "Groom",
			},
		},
		Admin: models.Admin{
			AuthMethod: models.AuthMethodBasic,
			Users: []models.User{
				admin,
			},
		},
		Party: models.Party{
			Enabled: true,
			Members: []models.PartyMember{
				{
					First: "Daniela",
					Last:  "Dalton",
					Title: "Maid of Honor",
				},
				{
					First: "Davey",
					Last:  "Gibson",
					Title: "Man of Honor",
				},
				{
					First: "Lily-Rose",
					Last:  "Cook",
					Title: "Bridesmaid",
				},
				{
					First: "Alan",
					Last:  "Fresco",
					Title: "Groomsman",
				},
			},
		},
	}
}
