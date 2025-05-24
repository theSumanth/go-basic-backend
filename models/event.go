package models

import "time"

type Event struct {
	ID          int
	Title       string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	Date        time.Time
	UserID      int
}

var events = []Event{}

func (e Event) Save() {
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
