package zoom

import (
	"encoding/json"
	"fmt"
	"log"
)

type IMDirectoryGroup struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	SearchByAccount   string `json:"search_by_account"`
	SearchByDomain    string `json:"search_by_domain"`
	SearchByMaAccount string `json:"search_by_ma_account"`
	TotalMembers      string `json:"total_members"`
	Type              string `json:"type"`
}

// Get a specific IM Directory Groups
func (z *Zoom) GetIMDirectoryGroup(group_id string) (IMDirectoryGroup, error) {
	// Create the url for getting groups
	endpoint := fmt.Sprintf("/im/groups/%v", group_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return IMDirectoryGroup{}, err
	}

	// Umarshal the response into IMDirectoryGroup struct
	var im_directory_group IMDirectoryGroup
	err = json.Unmarshal(response, &im_directory_group)
	if err != nil {
		log.Println(err)
		return IMDirectoryGroup{}, err
	}
	//returns im_directory_group structure
	return im_directory_group, nil
}
