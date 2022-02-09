package models

type PartyMember struct {
	First    string `yaml:"first" json:"first"`
	Last     string `yaml:"last" json:"last"`
	Title    string `yaml:"title" json:"title"`
	Email    string `yaml:"email" json:"email"`
	Info     string `yaml:"info" json:"info"`
	PhotoUrl string `yaml:"photo_url" json:"photo_url"`
}
