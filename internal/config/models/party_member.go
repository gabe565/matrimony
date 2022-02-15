package models

type PartyMember struct {
	First string `yaml:"first" json:"first"`
	Last  string `yaml:"last" json:"last"`
	Title string `yaml:"title" json:"title"`
	Email string `yaml:"email,omitempty" json:"email,omitempty"`
	Info  string `yaml:"info,omitempty" json:"info,omitempty"`
	Image string `yaml:"image,omitempty" json:"image,omitempty"`
}

func (member PartyMember) Fullname() string {
	return member.First + " " + member.Last
}
