package models

import "strings"

type Config struct {
	EventInfo EventInfo     `yaml:"info" json:"info"`
	Theme     Theme         `yaml:"theme" json:"theme"`
	Privacy   Privacy       `yaml:"privacy" json:"-"`
	Partners  []PartyMember `yaml:"partners" json:"partners"`
	Admin     Admin         `yaml:"admin" json:"-"`
	Sections  []Sections    `yaml:"sections" json:"sections"`
	RSVP      RSVP          `yaml:"rsvp" json:"rsvp"`
}

func (config Config) Title() string {
	names := make([]string, 0, len(config.Partners))
	for _, p := range config.Partners {
		names = append(names, p.Fullname())
	}
	return strings.Join(names, " & ")
}

func (config Config) TitleShort() string {
	names := make([]string, 0, len(config.Partners))
	for _, p := range config.Partners {
		names = append(names, p.First)
	}
	return strings.Join(names, " & ")
}
