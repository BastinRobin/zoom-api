package zoom

import (
	"encoding/json"
	"fmt"
	"log"
)

type GroupAdmins struct {
	NextPageToken string `json:"next_page_token"`
	PageSize      int    `json:"page_size"`
	TotalRecords  int    `json:"total_records"`
}
type Admins []admin
type admin struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	Name  string `json:"name"`
}

// return a list of group administrators
func (z *Zoom) GetGroupAdmins(group_id string) (GroupAdmins, error) {
	endpoint := fmt.Sprintf("/groups/%v/admins", group_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return GroupAdmins{}, err
	}

	// Umarshal the response into GroupAdmins struct
	var groupadmins GroupAdmins
	err = json.Unmarshal(response, &groupadmins)
	if err != nil {
		log.Println(err)
		return GroupAdmins{}, err
	}
	// return groupadmins struct
	return groupadmins, nil
}
