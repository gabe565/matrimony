package models

type Privacy struct {
	Mode     PrivacyMode `yaml:"mode"`
	Password string      `yaml:"password"`
}
