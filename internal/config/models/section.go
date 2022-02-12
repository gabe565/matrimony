package models

type HeroSection struct {
	Title string       `yaml:"title,omitempty" json:"title,omitempty"`
	Image ImageSection `yaml:"image" json:"image"`
}

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

type ScheduleSection struct {
	Title  string  `yaml:"title,omitempty" json:"title,omitempty"`
	Events []Event `yaml:"events" json:"events"`
}

type FaqSection struct {
	Title string        `yaml:"title" json:"title"`
	Items []TextSection `yaml:"items" json:"items"`
}

type Section struct {
	Hero     *HeroSection     `yaml:"hero,omitempty" json:"hero,omitempty"`
	Image    *ImageSection    `yaml:"image,omitempty" json:"image,omitempty"`
	Text     *TextSection     `yaml:"text,omitempty" json:"text,omitempty"`
	Party    *PartySection    `yaml:"party,omitempty" json:"party,omitempty"`
	Schedule *ScheduleSection `yaml:"schedule,omitempty" json:"schedule,omitempty"`
	Faq      *FaqSection      `yaml:"faq,omitempty" json:"faq,omitempty"`
}
