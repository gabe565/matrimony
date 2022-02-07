package models

import (
	"errors"
	"fmt"
)

//go:generate stringer -type PrivacyMode -linecomment

type PrivacyMode uint8

const (
	PrivacyModePrivate PrivacyMode = iota // private
	PrivacyModePublic                     // public
)

var ErrInvalidPrivacyMode = errors.New("invalid privacy mode")

func (i PrivacyMode) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

func (i *PrivacyMode) UnmarshalText(b []byte) error {
	s := string(b)
	for j := 0; j < len(_PrivacyMode_index)-1; j++ {
		if s == _PrivacyMode_name[_PrivacyMode_index[j]:_PrivacyMode_index[j+1]] {
			*i = PrivacyMode(j)
			return nil
		}
	}
	return fmt.Errorf("%v: %s", ErrInvalidPrivacyMode, s)
}
