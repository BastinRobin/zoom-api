package zoom

import (
	"encoding/json"
	"fmt"
	"log"
)

//Structure for list the Groups
type GroupList struct {
	TotalRecords int `json:"total_records"`
	Groups       Groups
}
type Groups []Group
type Group struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	TotalMembers int    `json:"total_members"`
}

//Structure for Group Locked Setting
type Group_Locked_Setting struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	TotalMembers int    `json:"total_members"`
}

// Return the list of all groups in zoom account
func (z *Zoom) GetListGroups() (GroupList, error) {
	groups, err := z.Request("/groups", "GET")
	if err != nil {
		log.Println(err)
		return GroupList{}, err
	}

	var groups_list GroupList

	// Unmarshal the groups into GroupsList struct
	err = json.Unmarshal(groups, &groups_list)
	if err != nil {
		log.Println(err)
		return GroupList{}, err
	}
	// returns group_list structure
	return groups_list, nil

}

// Get Group Locked Setting details
func (z *Zoom) GetLockedSetting(group_id string) (Group_Locked_Setting, error) {
	// Create the url for getting meetings
	endpoint := fmt.Sprintf("/groups/%v", group_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return Group_Locked_Setting{}, err
	}

	// Umarshal the response into struct
	var group_locked_setting Group_Locked_Setting
	err = json.Unmarshal(response, &group_locked_setting)
	if err != nil {
		log.Println(err)
		return Group_Locked_Setting{}, err
	}
	//returns group_locked_setting structure
	return group_locked_setting, nil
}
