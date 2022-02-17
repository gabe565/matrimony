package models

import "github.com/gabe565/matrimony/internal/util"

type Party struct {
	Model
	Name            string `json:"name"`
	SessionPassword string `json:"sessionPassword"`
}

func (p *Party) GenerateSessionPassword() (err error) {
	p.SessionPassword, err = util.RandomString(64)
	return err
}
