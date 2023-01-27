package base

import "encoding/json"

type Product struct {
	Name    string   `json:"name"`
	//Pic     string   `json:"pic"`
	Price   string   `json:"price"`
	Qty     int64    `json:"qty"`
	ProductOpts
	//Details string   `json:"details"`
	//Notes   []string `json:"notes"`
}

type ProductOpts struct {
	Pic     string   `json:"pic"`
	Details string   `json:"details"`
	Notes   []string `json:"notes"`
}

func (self Product) Json() ([]byte, error) {
	data, err := json.Marshal(self)
	return data, err
}

func (self Product) JsonString() (string, error) {
	data, err := json.Marshal(self)
	return string(data), err
}
