package zoom

import (
	"encoding/json"
	"fmt"
	"log"
)

type List_IMDirectory struct {
	PageCount    string `json:"page_count"`
	PageNumber   string `json:"page_number"`
	PageSize     string `json:"page_size"`
	TotalRecords string `json:"total_records"`
	Groups       Groups
}
type Groups []group
type group struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	SearchByAccount   string `json:"search_by_account"`
	SearchByDomain    string `json:"search_by_domain"`
	SearchByMaAccount string `json:"search_by_ma_account"`
	TotalMembers      string `json:"total_members"`
	Type              string `json:"type"`
}

type List_IMDirectoryMembers struct {
	PageCount    string `json:"page_count"`
	PageNumber   string `json:"page_number"`
	PageSize     string `json:"page_size"`
	TotalRecords string `json:"total_records"`
}
type Members []Member
type Member struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	ID        string `json:"id"`
	LastName  string `json:"last_name"`
	Type      string `json:"type"`
}

//return List of IM directory groups.
func (z *Zoom) GetListIMDirectory() (List_IMDirectory, error) {
	response, err := z.Request("/im/groups", "GET")
	if err != nil {
		log.Println(err)
		return List_IMDirectory{}, err
	}

	var list_imdirectory List_IMDirectory

	// Unmarshal the response into ListIMDirectory struct
	err = json.Unmarshal(response, &list_imdirectory)
	if err != nil {
		log.Println(err)
		return List_IMDirectory{}, err
	}
	//return listimdirectory
	return list_imdirectory, nil
}

// Return  an IM directory group members
func (z *Zoom) GetGroupMembers(group_id string) (List_IMDirectoryMembers, error) {
	endpoint := fmt.Sprintf("/im/groups/%v/members", group_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return List_IMDirectoryMembers{}, err
	}

	// Umarshal the response into List_IMDirectoryMembers struct
	var directorymembers List_IMDirectoryMembers
	err = json.Unmarshal(response, &directorymembers)
	if err != nil {
		log.Println(err)
		return List_IMDirectoryMembers{}, err
	}
	// return  directorymembers   struct
	return directorymembers, nil
}
