package zoom

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
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

//Structure for MeetingPolls
type MeetingPolls struct {
	Status       string `json:"status"`
	Title        string `json:"title"`
	TotalRecords int    `json:"total_records"`
	Polls        Polls
	Questions    Questions
}
type Polls []Poll
type Poll struct {
	Anonymous bool   `json:"anonymous"`
	ID        string `json:"id"`
	PollType  int    `json:"poll_type"`
}
type Questions []Question
type Question struct {
	AnswerRequired bool     `json:"answer_required"`
	Answers        []string `json:"answers"`
	Name           string   `json:"name"`
	RightAnswers   []string `json:"right_answers"`
	Type           string   `json:"type"`
}

type MeetingRegistrants struct {
	PageCount             int    `json:"page_count"`
	PageNumber            int    `json:"page_number"`
	PageSize              int    `json:"page_size"`
	Email                 string `json:"email"`
	FirstName             string `json:"first_name"`
	ID                    string `json:"id"`
	Industry              string `json:"industry"`
	JobTitle              string `json:"job_title"`
	JoinURL               string `json:"join_url"`
	LastName              string `json:"last_name"`
	NoOfEmployees         string `json:"no_of_employees"`
	Org                   string `json:"org"`
	Phone                 string `json:"phone"`
	PurchasingTimeFrame   string `json:"purchasing_time_frame"`
	RoleInPurchaseProcess string `json:"role_in_purchase_process"`
	State                 string `json:"state"`
	Status                string `json:"status"`
	Zip                   string `json:"zip"`
	TotalRecords          int    `json:"total_records"`
	Registrants           Registrants
	CustomQuestions       CustomQuestions
}
type Registrants []Registrant
type Registrant struct {
	Address    string    `json:"address"`
	City       string    `json:"city"`
	Comments   string    `json:"comments"`
	Country    string    `json:"country"`
	CreateTime time.Time `json:"create_time"`
}
type CustomQuestions []CustomerQuestion
type CustomerQuestion struct {
	Title string `json:"title"`
	Value string `json:"value"`
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

// Get Specific meeting details
func (z *Zoom) GetMeetingPolls(meeting_id string) (MeetingPolls, error) {
	// Create the url for getting meetings
	endpoint := fmt.Sprintf("/meetings/%v/polls", meeting_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return MeetingPolls{}, err
	}

	// Umarshal the response into MeetingPoll struct
	var meeting_poll MeetingPolls
	err = json.Unmarshal(response, &meeting_poll)
	if err != nil {
		log.Println(err)
		return MeetingPolls{}, err
	}

	return meeting_poll, nil
}

func (z *Zoom) GetMeetingRegistrants(meeting_id string) (MeetingRegistrants, error) {
	// Create the url for getting meetings
	endpoint := fmt.Sprintf("/meetings/%v/registrants", meeting_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return MeetingRegistrants{}, err
	}

	// Umarshal the response into MeetingRegistrants struct
	var meeting_registrants MeetingRegistrants
	err = json.Unmarshal(response, &meeting_registrants)
	if err != nil {
		log.Println(err)
		return MeetingRegistrants{}, err
	}

	return meeting_registrants, nil
}
