package contacts

import (
	"encoding/json"
	"fmt"
	"github.com/TheSnakeWitcher/mypeople/contacts/base"
)

func (self *People) ToPeopleData() *base.PeopleData {
	var arg = base.PeopleData{ID: self.ID, Name: self.Name}

    data,err := json.Marshal(self)
    if err != nil {
        fmt.Println("error marshaling in ToPeopleData: ",err)
    }

	err = json.Unmarshal(data,&arg)
    if err != nil {
        fmt.Println("error unmarshaling in ToPeopleData: ",err)
    }

	return &arg
}

func (self People) String() string {
	return fmt.Sprintf("\nid: %d  name: %s pic: %t groups: %s\n" ,
		self.ID,
		self.Name,
		len(self.Pic) > 1,
		self.Groups,
	)   
}
