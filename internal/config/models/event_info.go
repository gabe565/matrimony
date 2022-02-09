package models

import "time"

type EventInfo struct {
	DisplayName string    `yaml:"display_name" json:"display_name"`
	Date        time.Time `yaml:"date" json:"date"`
	Location    string    `yaml:"location" json:"location"`
	Timezone    string    `yaml:"timezone" json:"timezone"`
	Greeting    string    `yaml:"greeting" json:"greeting"`
	HeroImage   string    `yaml:"hero_image" json:"hero_image"`
}
