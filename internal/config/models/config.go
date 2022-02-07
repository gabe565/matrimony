package models

type Config struct {
	EventInfo EventInfo     `yaml:"info"`
	Privacy   Privacy       `yaml:"privacy"`
	Partners  []PartyMember `yaml:"partners"`
	Admin     Admin         `yaml:"admin"`
	Party     Party         `yaml:"party"`
}
