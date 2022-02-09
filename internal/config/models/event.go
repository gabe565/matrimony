package models

import "time"

type EventDate struct {
	Show  bool      `yaml:"show" json:"show"`
	Value time.Time `yaml:"value" json:"value"`
}

type Event struct {
	Title       string    `yaml:"title" json:"title"`
	Description string    `yaml:"description" json:"description"`
	Location    string    `yaml:"location" json:"location"`
	DisplayMap  bool      `yaml:"displayMap" json:"displayMap"`
	DressCode   string    `yaml:"dress_code" json:"dressCode"`
	Start       EventDate `yaml:"start" json:"start"`
	End         EventDate `yaml:"end" json:"end"`
}
