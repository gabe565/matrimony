package models

import "time"

type Event struct {
	Title       string     `yaml:"title" json:"title"`
	Description string     `yaml:"description" json:"description"`
	Location    string     `yaml:"location" json:"location"`
	DressCode   string     `yaml:"dressCode" json:"dressCode"`
	Start       *time.Time `yaml:"start" json:"start"`
}
