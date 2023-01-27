package base

import (
// "encoding/json"
// "encoding/json"
// "time"
)

type PeopleData struct {
	ID     int64    `db:"id" json:"id"`
	Name   string   `db:"name" json:"name"`
	Pic    string   `db:"pic" json:"pic"`
	Groups []string `db:"groups" json:"groups"`
	//Emails     sql.NullString `db:"emails" json:"emails"`
	//Phones     sql.NullString `db:"phones" json:"phones"`
	//SocialNets sql.NullString `db:"social_nets" json:"social_nets"`
	//Wallets    sql.NullString `db:"wallets" json:"wallets"`
	Locations Location `db:"locations" json:"locations"`
	Events    Event    `db:"events" json:"events"`
	Notes     []string `db:"notes" json:"notes"`
}

func NewPeopleData(id int64, name string) *PeopleData {

	out := &PeopleData{
		ID:   id,
		Name: name,
	}

	return out
}
