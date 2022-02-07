package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (u *User) HashPassword(password string) (err error) {
	u.Password, err = HashPassword(password)
	return err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (u User) CheckPasswordHash(password string) bool {
	return CheckPasswordHash(password, u.Password)
}
