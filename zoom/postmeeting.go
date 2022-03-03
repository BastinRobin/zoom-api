package zoom

import (
	"encoding/json"
	"fmt"
	"log"
)

type Response_meeting struct {
	NextPageToken string `json:"next_page_token"`
	PageSize      int    `json:"page_size"`
	Participants  []struct {
		//	DateTime time.Time `json:"date_time"`
		Email   string `json:"email"`
		Quality string `json:"quality"`
		UserID  string `json:"user_id"`
	} `json:"participants"`
}

// Get Specific postmeeting details
func (z *Zoom) GetPostMeeting(meeting_id string) (Response_meeting, error) {
	// Create the url for getting meetings
	endpoint := fmt.Sprintf("/metrics/meetings/%v/participants/satisfaction", meeting_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return Response_meeting{}, err
	}

	// Umarshal the response into struct
	var postmeetings Response_meeting
	err = json.Unmarshal(response, &postmeetings)
	if err != nil {
		log.Println(err)
		return Response_meeting{}, err
	}

	return postmeetings, nil
}
