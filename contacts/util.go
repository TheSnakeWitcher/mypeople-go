package contacts

import (
	"encoding/json"
	"fmt"

	"github.com/TheSnakeWitcher/mypeople/contacts/base"
)

func checkError(err error) {
    if err != nil {
        fmt.Print("error in ToPeopleData: ",err,"\n\n")
    }
}

func (self *People) ToPeopleData() *base.PeopleData {
	out := base.NewPeopleData(
        self.ID,
        self.Name,
        self.Pic,
        base.PeopleDataOpts{},
	)

	err := json.Unmarshal([]byte(self.Groups),&out.Groups)
	checkError(err)

	err = json.Unmarshal([]byte(self.Emails),&out.Emails)
	checkError(err)

	err = json.Unmarshal([]byte(self.Phones),&out.Phones)
	checkError(err)

	err = json.Unmarshal([]byte(self.SocialNets),&out.SocialNets)
	checkError(err)

	err = json.Unmarshal([]byte(self.Wallets),&out.Wallets)
	checkError(err)

	err = json.Unmarshal([]byte(self.Events),&out.Events)
	checkError(err)

	err = json.Unmarshal([]byte(self.Locations),&out.Locations)
	checkError(err)
	    
	err = json.Unmarshal([]byte(self.Notes),&out.Notes)
	checkError(err)
	      

	return out
}

func (self People) String() string {
	return fmt.Sprintf("\n" +
	    "Id: %d  Name: %s Pic: %t\n" +
	    "Groups: %s\n" +
	    "Emails: %s\n" +
	    "Phones: %s\n" +
	    "SocialNet: %s\n" +
	    "Wallets: %s\n" +
	    "Events: %s\n" +
	    "Locations: %s\n" +
	    "Notes: %s\n\n",
		self.ID,
		self.Name,
		len(self.Pic) > 1,
		self.Groups,
		self.Emails,
		self.Phones,
		self.SocialNets,
		self.Wallets,
		self.Events,
		self.Locations,
		self.Notes,
	)   
}

func (self People)MarshalJSON() ([]byte,error) {
	var out base.PeopleData

	errGroups := json.Unmarshal([]byte(self.Groups),&out.Groups)
    if errGroups != nil {
        fmt.Print("error in ToPeopleData: ",errGroups,"\n\n")
    }
    fmt.Print("unmarshal groups: ",out.Groups,"\n\n")

    return []byte{},nil
}
