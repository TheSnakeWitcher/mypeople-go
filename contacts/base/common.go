package base

type BookType interface {
	string | []string |
	Event | []Event |
	Location | []Location
}

type Book[T BookType] map[string]T
