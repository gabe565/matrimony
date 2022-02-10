package models

import "time"

type Event struct {
	Title       string     `yaml:"title" json:"title"`
	Description string     `yaml:"description" json:"description"`
	Location    string     `yaml:"location" json:"location"`
	ShowMap     bool       `yaml:"showMap" json:"showMap"`
	DressCode   string     `yaml:"dressCode" json:"dressCode"`
	Start       *time.Time `yaml:"start" json:"start"`
	End         *time.Time `yaml:"end" json:"end"`
}
