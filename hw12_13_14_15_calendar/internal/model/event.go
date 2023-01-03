package model

import "time"

type Event struct {
	ID               string
	Title            string
	DateTime         time.Time
	Duration         int
	Description      string
	OwnerID          string
	NotificationTime time.Time
}
