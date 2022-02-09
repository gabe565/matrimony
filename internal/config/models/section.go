package models

type ImageSection struct {
	Src string `yaml:"src" json:"src"`
	Alt string `yaml:"alt,omitempty" json:"alt,omitempty"`
}

type TextSection struct {
	Title   string `yaml:"title,omitempty" json:"title,omitempty"`
	File    string `yaml:"file" json:"-"`
	Content string `yaml:"content,omitempty" json:"content,omitempty"`
}

type PartySection struct {
	Title   string        `yaml:"title,omitempty" json:"title,omitempty"`
	Members []PartyMember `yaml:"members" json:"members"`
}

type Section struct {
	Image *ImageSection `yaml:"image,omitempty" json:"image,omitempty"`
	Text  *TextSection  `yaml:"text,omitempty" json:"text,omitempty"`
	Party *PartySection `yaml:"party,omitempty" json:"party,omitempty"`
}
