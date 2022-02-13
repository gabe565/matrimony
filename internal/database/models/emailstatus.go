package models

//go:generate stringer -type EmailStatus -linecomment

type EmailStatus uint8

const (
	EmailStatusNone          EmailStatus = iota // None
	EmailStatusSent                             // Sent
	EmailStatusEmailOpened                      // Email Opened
	EmailStatusWebsiteOpened                    // Website Opened
)
