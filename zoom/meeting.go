package zoom

import (
	"encoding/json"
	"fmt"
	"log"
)

type Responses struct {
	ClientFeedbackDetails []struct {
		Email           string `json:"email"`
		MeetingID       string `json:"meeting_id"`
		ParticipantName string `json:"participant_name"`
		//Time            time.Time `json:"time"`
	} `json:"client_feedback_details"`
	From          string `json:"from"`
	NextPageToken string `json:"next_page_token"`
	PageSize      int    `json:"page_size"`
	To            string `json:"to"`
}

// Get Specific meeting details
func (z *Zoom) GetClientFeedback(feedbackId string) (Responses, error) {
	// Create the url for getting meetings
	endpoint := fmt.Sprintf("/metrics/client/feedback/%v/", feedbackId)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return Responses{}, err
	}

	// Umarshal the response into struct
	var meeting Responses
	err = json.Unmarshal(response, &meeting)
	if err != nil {
		log.Println(err)
		return Responses{}, err
	}

	return meeting, nil
}
