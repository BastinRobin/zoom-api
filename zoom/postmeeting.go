package zoom

import (
	"encoding/json"
	"fmt"
	"log"
)

type PostMeetingFeedback struct {
	NextPageToken string `json:"next_page_size"`
	PageSize      int    `json:"page_size"`
	Partcipantss  Partcipantss
}
type Partcipantss []Partcipant
type Partcipant struct {
	Email   string `json:"email"`
	Quality string `json:"quality"`
	UserID  string `json:"user_id"`
}

// Get Specific postmeeting details
func (z *Zoom) GetPostMeetingFeedback(meeting_id string) (PostMeetingFeedback, error) {
	// Create the url for getting meetings
	endpoint := fmt.Sprintf("/metrics/meetings/%v/participants/satisfaction", meeting_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return PostMeetingFeedback{}, err
	}

	// Umarshal the response into PostMeetingFeedback struct
	var postmeetings PostMeetingFeedback
	err = json.Unmarshal(response, &postmeetings)
	if err != nil {
		log.Println(err)
		return PostMeetingFeedback{}, err
	}
	//return PostMeetingFeedback structure
	return postmeetings, nil
}
