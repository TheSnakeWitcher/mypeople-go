package base

import (
	"testing"
	"time"

	//"github.com/stretchr/testify/suite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

//type TestEventSuite struct {
//    suite.Suite
//}

type TestEventMock struct { 
    mock.Mock
    Event
}


func TestEvent(t *testing.T)  {
    var (
        testDate = time.Now()
        testName = "my event"
        testRequirements = []string{"req"}
        testRelatedPeople = []string{"people"}
        testNotes = []string{"note"}

        testEventOpts = EventOpts{ 
            Requirements: testRequirements,
            RelatedPeople: testRelatedPeople,
            Notes: testNotes,
        }
    )

    event := NewEvent( testName, testDate, testEventOpts)

    mockData := new(TestEventMock)
    mockData.On("NewEvent",testName,testDate,testEventOpts).Return(event)
    mockData.On("Json").Return(event.Json())


    require.Equal(t,event.Name,testName)
    require.Equal(t,event.Date,testDate)
    assert.Equal(t,event.Requirements,testRequirements)
}
