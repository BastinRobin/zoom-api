package zoom

import (
	"encoding/json"
	"fmt"
	"log"
)

type Response_recording struct {
	NextPageToken string `json:"next_page_token"`
	PageCount     string `json:"page_count"`
	PageSize      string `json:"page_size"`
	Participants  Participants
	Details       Details
}

type Participants []Participant
type Participant struct {
	ID       string `json:"id"`
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	Details  Details
}
type Details []Detail
type Detail struct {
	Content   string `json:"content"`
	EndTime   string `json:"end_time"`
	StartTime string `json:"start_time"`
}

// return recording details of the participant
func (z *Zoom) GetRecordingDetails(meeting_id string) (Response_recording, error) {
	// Create the url for getting meetings
	endpoint := fmt.Sprintf("/metrics/meetings/%v/participants/sharing", meeting_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return Response_recording{}, err
	}

	// Umarshal the response into struct
	var recording Response_recording
	err = json.Unmarshal(response, &recording)
	if err != nil {
		log.Println(err)
		return Response_recording{}, err
	}
	// return recording struct
	return recording, nil
}
