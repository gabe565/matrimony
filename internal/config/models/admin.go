package models

type Admin struct {
	Auth  AdminAuth `yaml:"auth" json:"auth"`
	Users []User    `yaml:"users" json:"users"`
}

type AdminAuth struct {
	Method AuthMethod `yaml:"method" json:"method"`
	Header string     `yaml:"header" json:"header"`
}
