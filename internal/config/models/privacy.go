package models

type Privacy struct {
	Mode     PrivacyMode `yaml:"mode" json:"mode"`
	Password string      `yaml:"password" json:"password"`
}
