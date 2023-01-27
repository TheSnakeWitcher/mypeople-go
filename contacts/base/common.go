package base

import (
    "encoding/json"
)

type BookType interface {
	[]string |
	Event |
	Location
}

type Book[T BookType] map[string]T

func (self Book[BookType]) Json() (string) {
    data ,err := json.MarshalIndent(self,"","\t")
    if err != nil {
        return ""
    }

    return string(data)
}
