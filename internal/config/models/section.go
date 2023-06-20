package models

type Sections struct {
	Hero     *HeroSection     `yaml:"hero,omitempty" json:"hero,omitempty"`
	Image    *ImageSection    `yaml:"image,omitempty" json:"image,omitempty"`
	Text     *TextSection     `yaml:"text,omitempty" json:"text,omitempty"`
	Party    *PartySection    `yaml:"party,omitempty" json:"party,omitempty"`
	Schedule *ScheduleSection `yaml:"schedule,omitempty" json:"schedule,omitempty"`
	Faq      *FaqSection      `yaml:"faq,omitempty" json:"faq,omitempty"`
	Links    *LinksSection    `yaml:"links,omitempty" json:"links,omitempty"`
}

type Section struct {
	Title       string `yaml:"title" json:"title,omitempty"`
	HideFromNav bool   `yaml:"hideFromNav" json:"hideFromNav,omitempty"`
}

type HeroSection struct {
	Title string       `yaml:"title" json:"title,omitempty"`
	Image ImageSection `yaml:"image" json:"image"`
}

type ImageSection struct {
	Src string `yaml:"src" json:"src"`
	Alt string `yaml:"alt" json:"alt,omitempty"`
}

type TextSection struct {
	Section `yaml:",inline"`
	File    string `yaml:"file" json:"-"`
	Content string `yaml:"content" json:"content,omitempty"`
}

type PartySection struct {
	Section `yaml:",inline"`
	Members []PartyMember `yaml:"members" json:"members"`
}

type ScheduleSection struct {
	Section `yaml:",inline"`
	Events  []Event `yaml:"events" json:"events"`
}

type FaqSection struct {
	Section `yaml:",inline"`
	Items   []TextSection `yaml:"items" json:"items"`
}

type LinksSection struct {
	Section `yaml:",inline"`
	Moments Link `yaml:"moments" json:"moments"`
	RSVP    Link `yaml:"rsvp" json:"rsvp"`
}

type Link struct {
	Hidden  bool   `yaml:"hidden" json:"hidden,omitempty"`
	Title   string `yaml:"title" json:"title,omitempty"`
	Content string `yaml:"content" json:"content"`
}
