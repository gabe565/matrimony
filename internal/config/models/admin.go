package models

type Admin struct {
	AuthMethod AuthMethod `yaml:"auth_method" json:"auth_method"`
	Users      []User     `yaml:"users" json:"users"`
	Header     string     `yaml:"header" json:"header"`
}
