package models

type Config struct {
	EventInfo EventInfo     `yaml:"info"`
	Sections  []Section     `yaml:"sections"`
	Privacy   Privacy       `yaml:"privacy" json:"-"`
	Partners  []PartyMember `yaml:"partners"`
	Admin     Admin         `yaml:"admin" json:"-"`
	Party     Party         `yaml:"party"`
}
