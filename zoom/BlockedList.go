package zoom

import (
	"encoding/json"
	"fmt"
	"log"
)

type Blocked struct {
	NextPageToken string `json:"next_page_token"`
	PageSize      int    `json:"page_size"`
	TotalRecords  int    `json:"total_records"`
}
type BlockedLists []BlockedList
type BlockedList struct {
	BlockType   string `json:"block_type"`
	Comment     string `json:"comment"`
	ID          string `json:"id"`
	MatchType   string `json:"match_type"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
}

type BlockedListDetails struct {
	BlockType   string `json:"block_type"`
	Comment     string `json:"comment"`
	ID          string `json:"id"`
	MatchType   string `json:"match_type"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
}

//return the list of all the blocked lists
func (z *Zoom) GetListBlockedLists() (Blocked, error) {
	response, err := z.Request("/phone/blocked_list", "GET")
	if err != nil {
		log.Println(err)
		return Blocked{}, err
	}

	var blocked Blocked

	// Unmarshal the response into  Blocked struct
	err = json.Unmarshal(response, &blocked)
	if err != nil {
		log.Println(err)
		return Blocked{}, err
	}
	//return blocked struct
	return blocked, nil

}

// return information about a specific blocked list
func (z *Zoom) GetBlockedListDetails(blockedList_id string) (BlockedListDetails, error) {
	endpoint := fmt.Sprintf("/phone/blocked_list/%v", blockedList_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return BlockedListDetails{}, err
	}

	// Umarshal the response into BlockedListDetails struct
	var blockedlistdetails BlockedListDetails
	err = json.Unmarshal(response, &blockedlistdetails)
	if err != nil {
		log.Println(err)
		return BlockedListDetails{}, err
	}
	// return blockedlistdetails struct
	return blockedlistdetails, nil
}
