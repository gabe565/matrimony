package models

import (
	"errors"
	"fmt"
)

//go:generate stringer -type AuthMethod -linecomment

type AuthMethod uint8

const (
	AuthMethodBasic  AuthMethod = iota // basic
	AuthMethodHeader                   // header
)

var ErrInvalidAuthMethod = errors.New("invalid auth method")

func (i AuthMethod) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

func (i *AuthMethod) UnmarshalText(b []byte) error {
	s := string(b)
	for j := 0; j < len(_AuthMethod_index)-1; j++ {
		if s == _AuthMethod_name[_AuthMethod_index[j]:_AuthMethod_index[j+1]] {
			*i = AuthMethod(j)
			return nil
		}
	}
	return fmt.Errorf("%v: %s", ErrInvalidAuthMethod, s)
}
