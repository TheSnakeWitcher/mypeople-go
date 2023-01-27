package base

import (
    "fmt"
    "time"
    "testing"
)

func TestEvent(t *testing.T)  {

    event := NewEvent(
        "my event",
        time.Now(),
        EventOpts{ Requirements: []string{"my req"} },
    )

    fmt.Println("event json: ",event.Json())
    fmt.Println("event json: ",event.JsonString())
}
