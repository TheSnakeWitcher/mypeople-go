package base

import (
	"encoding/json"
	"time"
)

type Event struct {
	Name          string    `json:"name"`
	Date          time.Time `json:"date"`
    EventOpts
}

type EventOpts struct {
	Requirements  []string  `json:"requirements"`
	RelatedPeople []string  `json:"related_people"`
	Notes         []string  `json:"notes"`
}

func NewEvent(name string, date time.Time,opts EventOpts) *Event {
	return &Event{
		Name:          name,
		Date:          date,
		EventOpts: EventOpts{
		    Requirements:  opts.Requirements,
		    RelatedPeople: opts.RelatedPeople,
		    Notes:         opts.Notes,
		},
	}
}

func (self Event) Json() []byte {
	data, err := json.Marshal(self)
	if err != nil {
	   return []byte{}
	}
	return data
}

func (self Event) JsonString() string {
	data, err := json.Marshal(self)
	if err != nil {
	   return ""
	}
	return string(data)
}
