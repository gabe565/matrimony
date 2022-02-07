package models

import "time"

type EventInfo struct {
	DisplayName string    `yaml:"display_name"`
	Date        time.Time `yaml:"date"`
	Location    string    `yaml:"location"`
	Timezone    string    `yaml:"timezone"`
	Greeting    string    `yaml:"greeting"`
}
