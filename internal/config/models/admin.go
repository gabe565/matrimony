package models

type Admin struct {
	AuthMethod AuthMethod `yaml:"auth_method"`
	Users      []User     `yaml:"users"`
	Header     string     `yaml:"header"`
}
