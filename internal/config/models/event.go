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
	ShowMap     bool      `yaml:"showMap" json:"showMap"`
	DressCode   string    `yaml:"dressCode" json:"dressCode"`
	Start       EventDate `yaml:"start" json:"start"`
	End         EventDate `yaml:"end" json:"end"`
}
