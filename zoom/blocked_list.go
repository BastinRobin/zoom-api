package zoom

import (
	"encoding/json"
	"log"
)

type BlockedList struct {
	NextPageToken string `json:"next_page_token"`
	PageSize      int    `json:"page_size"`
	TotalRecords  int    `json:"total_records"`
	Blocks        Blocks
}
type Blocks []Block
type Block struct {
	BlockType   string `json:"block_type"`
	Comment     string `json:"comment"`
	ID          string `json:"id"`
	MatchType   string `json:"match_type"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
}

// Return the list of blocked_list in zoom account
func (z *Zoom) ListBlockedList() (BlockedList, error) {
	blocked_list, err := z.Request("/phone/blocked_list", "GET")
	if err != nil {
		log.Println(err)
		return BlockedList{}, err
	}

	var block_list BlockedList

	// Unmarshal the response into Blocked_List struct
	err = json.Unmarshal(blocked_list, &block_list)
	if err != nil {
		log.Println(err)
		return BlockedList{}, err
	}
	//return Block_list Structure
	return block_list, nil

}
