package models

import "time"

type Button struct {
	Type string `yaml:"type" json:"type"`
}

type Event struct {
	Title       string     `yaml:"title" json:"title"`
	Description string     `yaml:"description" json:"description"`
	Location    string     `yaml:"location" json:"location"`
	Buttons     []Button   `yaml:"buttons,omitempty" json:"buttons,omitempty"`
	DressCode   string     `yaml:"dressCode" json:"dressCode"`
	Start       *time.Time `yaml:"start" json:"start"`
}
