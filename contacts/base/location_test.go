package base

import (
    "testing"
    "fmt"
)

func TestLocation(t *testing.T)  {

    loc := NewLocation("my loc",LocationOpts{Address: "my address"})

    fmt.Println("loc json: ",loc.Json())
    fmt.Println("loc json string: ",loc.JsonString())

}
