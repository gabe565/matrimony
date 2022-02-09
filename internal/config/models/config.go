package models

type Config struct {
	EventInfo EventInfo     `yaml:"info" json:"info"`
	Sections  []Section     `yaml:"sections" json:"sections"`
	Privacy   Privacy       `yaml:"privacy" json:"-"`
	Partners  []PartyMember `yaml:"partners" json:"partners"`
	Admin     Admin         `yaml:"admin" json:"-"`
	Party     Party         `yaml:"party" json:"party"`
}
