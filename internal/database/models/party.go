package models

type Party struct {
	Model
	Name string `gorm:"uniqueIndex"`
}
