package zoom

import (
	"encoding/json"
	"fmt"
	"log"
)

type Meeting struct {
	UUID  string `json:"uuid"`
	Id    int    `json:"id"`
	Topic string `json:"topic"`
	Type  int    `json:"type"`
}

// Get Specific meeting details
func (z *Zoom) GetMeeting(meeting_id string) (Meeting, error) {
	// Create the url for getting meetings
	endpoint := fmt.Sprintf("%v/meetings/%v", z.BaseUrl, meeting_id)
	fmt.Println(endpoint)

	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return Meeting{}, err
	}

	// Umarshal the response into struct
	var meeting Meeting
	err = json.Unmarshal(response, &meeting)
	if err != nil {
		log.Println(err)
		return Meeting{}, err
	}

	return meeting, nil
}
