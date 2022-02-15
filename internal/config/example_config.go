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
			Date:     time.Now(),
			Location: "Oklahoma City, OK, USA",
			Timezone: "America/Chicago",
			Greeting: "We can't wait to share our special day with you!",
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
			Auth: models.AdminAuth{
				Method: models.AuthMethodBasic,
			},
			Users: []models.User{
				admin,
			},
		},
		Sections: []models.Sections{
			{
				Hero: &models.HeroSection{
					Image: models.ImageSection{
						Src: "",
					},
				},
			},
			{
				Text: &models.TextSection{
					Section: models.Section{
						Title: "About Us",
					},
					Content: "",
				},
			},
			{
				Party: &models.PartySection{
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
			},
		},
	}
}
