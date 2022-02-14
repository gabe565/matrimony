package models

type BaseSection struct {
	Title       string `yaml:"title,omitempty" json:"title,omitempty"`
	HideFromNav bool   `yaml:"hideFromNav,omitempty" json:"hideFromNav,omitempty"`
}

type HeroSection struct {
	Title string       `yaml:"title,omitempty" json:"title,omitempty"`
	Image ImageSection `yaml:"image" json:"image"`
}

type ImageSection struct {
	Src string `yaml:"src" json:"src"`
	Alt string `yaml:"alt,omitempty" json:"alt,omitempty"`
}

type TextSection struct {
	BaseSection `yaml:",inline"`
	File        string `yaml:"file" json:"-"`
	Content     string `yaml:"content,omitempty" json:"content,omitempty"`
}

type PartySection struct {
	BaseSection `yaml:",inline"`
	Members     []PartyMember `yaml:"members" json:"members"`
}

type ScheduleSection struct {
	BaseSection `yaml:",inline"`
	Events      []Event `yaml:"events" json:"events"`
}

type FaqSection struct {
	BaseSection `yaml:",inline"`
	Items       []TextSection `yaml:"items" json:"items"`
}

type MomentsSection struct {
	Title   string `yaml:"title,omitempty" json:"title,omitempty"`
	Content string `yaml:"content" json:"content"`
}

type Section struct {
	Hero     *HeroSection     `yaml:"hero,omitempty" json:"hero,omitempty"`
	Image    *ImageSection    `yaml:"image,omitempty" json:"image,omitempty"`
	Text     *TextSection     `yaml:"text,omitempty" json:"text,omitempty"`
	Party    *PartySection    `yaml:"party,omitempty" json:"party,omitempty"`
	Schedule *ScheduleSection `yaml:"schedule,omitempty" json:"schedule,omitempty"`
	Faq      *FaqSection      `yaml:"faq,omitempty" json:"faq,omitempty"`
	Moments  *MomentsSection  `yaml:"moments,omitempty" json:"moments,omitempty"`
}
