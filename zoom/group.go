package zoom

import (
	"encoding/json"
	"fmt"
	"log"
)

type Group struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	TotalMembers int    `json:"total_members"`
}

// return a group
func (z *Zoom) GetGroup(group_id string) (Group, error) {
	// Create the url for getting group
	endpoint := fmt.Sprintf("/groups/%v", group_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return Group{}, err
	}

	// Umarshal the response into Group struct
	var group Group
	err = json.Unmarshal(response, &group)
	if err != nil {
		log.Println(err)
		return Group{}, err
	}
	// return group struct
	return group, nil
}
