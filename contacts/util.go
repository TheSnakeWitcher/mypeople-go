package contacts

import (
	"encoding/json"
	"fmt"
	"github.com/TheSnakeWitcher/mypeople/contacts/base"
)

func (self *People) ToPeopleData() *base.PeopleData {
	//var arg = base.PeopleData{ID: self.ID, Name: self.Name}
	var arg = base.PeopleData{}
    data,err := json.Marshal(self)
    if err != nil {
        fmt.Print("error in ToPeopleData: ",err,"\n\n")
    }
    fmt.Print("marshal data: ",string(data),"\n\n")

	err = json.Unmarshal(data,&arg)
    if err != nil {
        fmt.Print("error in ToPeopleData: ",err,"\n\n")
    }
    fmt.Print("unmarshal data: ",string(data),"\n\n")


	//var arg2 = base.PeopleData{}
    //marshalGroups,errGroups := json.Marshal(self.Groups)
    //if errGroups != nil {
    //    fmt.Print("error in ToPeopleData: ",errGroups,"\n\n")
    //}
    //fmt.Print("marshal groups: ",string(marshalGroups),"\n\n")

	//errGroups = json.Unmarshal(marshalGroups,&arg2.Groups)
    //if errGroups != nil {
    //    fmt.Print("error in ToPeopleData: ",errGroups,"\n\n")
    //}
    //fmt.Print("unmarshal groups: ",arg2.Groups,"\n\n")

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
