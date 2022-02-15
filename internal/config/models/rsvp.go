package models

type RSVP struct {
	Enabled             bool        `yaml:"enabled" json:"enabled"`
	Protected           bool        `yaml:"protected" json:"protected"`
	RestrictToGuestList bool        `yaml:"restrictToGuestList" json:"restrictToGuestList"`
	AllowPlusOnes       bool        `yaml:"allowPlusOnes" json:"allowPlusOnes"`
	Questions           []Questions `yaml:"questions" json:"questions"`
}

type Questions struct {
	Attendance *AttendanceQuestion `yaml:"attendance,omitempty" json:"attendance,omitempty"`
	Text       *TextQuestion       `yaml:"text,omitempty" json:"text,omitempty"`
	Number     *NumberQuestion     `yaml:"number,omitempty" json:"number,omitempty"`
	Choice     *ChoiceQuestion     `yaml:"choice,omitempty" json:"choice,omitempty"`
}

type Question struct {
	Field  string `yaml:"field" json:"field"`
	Prompt string `yaml:"prompt" json:"prompt"`
}

type ConditionalQuestion struct {
	Per string `yaml:"per" json:"per"`
}

type AttendanceQuestion struct {
	Question `yaml:",inline"`
	Yes      Answer `yaml:"yes" json:"yes"`
	No       Answer `yaml:"no" json:"no"`
}

type TextQuestion struct {
	Question            `yaml:",inline"`
	ConditionalQuestion `yaml:",inline"`
}

type NumberQuestion struct {
	Question            `yaml:",inline"`
	ConditionalQuestion `yaml:",inline"`
}

type ChoiceQuestion struct {
	Question            `yaml:",inline"`
	ConditionalQuestion `yaml:",inline"`
	Choices             []Answer `yaml:"choices" json:"choices"`
}

type Answer struct {
	Answer   string      `yaml:"answer" json:"answer"`
	Followup []Questions `yaml:"followup" json:"followup,omitempty"`
}
