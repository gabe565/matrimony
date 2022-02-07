package models

type Party struct {
	Enabled bool          `yaml:"enabled"`
	Members []PartyMember `yaml:"members"`
}
