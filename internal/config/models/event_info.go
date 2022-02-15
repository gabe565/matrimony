package models

import (
	"strings"
	"time"
)

type EventInfo struct {
	Date     time.Time `yaml:"date" json:"date"`
	Location string    `yaml:"location" json:"location"`
	Timezone string    `yaml:"timezone" json:"timezone"`
	Greeting string    `yaml:"greeting" json:"greeting"`
	Keywords Keywords  `yaml:"keywords" json:"keywords"`
}

type Keywords []string

func (keywords Keywords) Join() string {
	return strings.Join(keywords, ", ")
}
