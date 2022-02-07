package models

type PartyMember struct {
	First    string `yaml:"first"`
	Last     string `yaml:"last"`
	Title    string `yaml:"title"`
	Email    string `yaml:"email"`
	Info     string `yaml:"info"`
	PhotoUrl string `yaml:"photo_url"`
}
