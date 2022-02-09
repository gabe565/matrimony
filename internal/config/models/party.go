package models

type Party struct {
	Enabled bool          `yaml:"enabled" json:"enabled"`
	Members []PartyMember `yaml:"members" json:"members"`
}
