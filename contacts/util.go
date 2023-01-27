package contacts

import (
	"encoding/json"
	"fmt"
	"github.com/TheSnakeWitcher/mypeople/contacts/base"
)

func (self *People) ToPeopleData() *base.PeopleData {
	var arg = base.PeopleData{ID: self.ID, Name: self.Name}

	if self.Pic.Valid {
		arg.Pic = self.Pic.String
	}

	if self.Groups.Valid {
		json.Unmarshal([]byte(self.Groups.String), &arg.Groups)
	}

	return &arg
}

func (self People) String() string {
	return fmt.Sprintf("\nid: %d  name: %s pic: %t groups: %s\n" ,
		self.ID,
		self.Name,
		self.Pic.Valid,
		self.Groups.String,
	)   
}
