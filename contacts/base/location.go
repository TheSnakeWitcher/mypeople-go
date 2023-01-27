package base

import "encoding/json"

type Location struct {
	Name        string   `json:"name"`
    LocationOpts	
}

type LocationOpts struct {
	Pic         string   `json:"pic"`
	Address     string   `json:"address"`
	Coordinates string   `json:"coordinates"`
	Notes       []string `json:"notes"`
}


func NewLocation(name string, opts LocationOpts) *Location {
	return &Location{
		Name: name,
		LocationOpts : LocationOpts{
		    Pic: opts.Pic,
		    Address:     opts.Address,
		    Coordinates: opts.Coordinates,
		    Notes:       opts.Notes,
		},
	}
}

func (self Location) Json() []byte {
	data, err := json.Marshal(self)
	if err != nil {
	   return []byte{}
	}
	return data
}

func (self Location) JsonString() string {
	data, err := json.Marshal(self)
	if err != nil {
	   return ""
	}
	return string(data)
}
