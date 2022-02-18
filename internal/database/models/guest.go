package models

import (
	"gorm.io/datatypes"
)

type Guest struct {
	Model
	FirstName    string  `json:"first"`
	LastName     *string `json:"last"`
	EmailAddress *string `json:"email"`
	Tags         []Tag   `json:"tags" gorm:"many2many:guest_tags"`

	EmailStatus *EmailStatus `json:"emailStatus"`

	RSVP datatypes.JSON `json:"rsvp"`

	PartyID uint  `json:"partyId"`
	Party   Party `json:"party"`
}
