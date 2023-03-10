package base

import "encoding/json"

type PeopleDataOpts struct {
	Groups     []string       `db:"groups" json:"groups"`
	Emails     Book[[]string] `db:"emails" json:"emails"`
	Phones     Book[[]string] `db:"phones" json:"phones"`
	SocialNets Book[[]string] `db:"social_nets" json:"social_nets"`
	Wallets    Book[[]string] `db:"wallets" json:"wallets"`
	Locations  Book[Location] `db:"locations" json:"locations"`
	Events     Book[Event]    `db:"events" json:"events"`
	Notes      []string       `db:"notes" json:"notes"`
}

type PeopleData struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
	Pic  string `db:"pic" json:"pic"`
	PeopleDataOpts
}

func NewPeopleData(id int64, name, pic string, opts PeopleDataOpts) *PeopleData {
	return &PeopleData{
		id,
		name,
		pic,
		PeopleDataOpts{
			Groups:     opts.Groups,
			Emails:     opts.Emails,
			Phones:     opts.Phones,
			SocialNets: opts.SocialNets,
			Wallets:    opts.Wallets,
			Locations:  opts.Locations,
			Events:     opts.Events,
			Notes:      opts.Notes,
		},
	}
}

func (self PeopleData) Json() string {
	data, err := json.MarshalIndent(self, "", "\t")
	if err != nil {
		return ""
	}

	return string(data)
}
