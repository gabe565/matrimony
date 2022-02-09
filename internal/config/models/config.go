package models

type Config struct {
	EventInfo EventInfo     `yaml:"info" json:"info"`
	Privacy   Privacy       `yaml:"privacy" json:"-"`
	Partners  []PartyMember `yaml:"partners" json:"partners"`
	Admin     Admin         `yaml:"admin" json:"-"`
	Sections  []Section     `yaml:"sections" json:"sections"`
}
