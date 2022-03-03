package zoom

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Response_rooms struct {
	From         string `json:"from"`
	IssueDetails []struct {
		Issue string    `json:"issue"`
		Time  time.Time `json:"time"`
	} `json:"issue_details"`
	PageCount    int    `json:"page_count"`
	PageSize     int    `json:"page_size"`
	To           string `json:"to"`
	TotalRecords int    `json:"total_records"`
}

// Get Specific Rooms with issues details
func (z *Zoom) GetRoomsWithIssues(zoomroomId string) (Response_rooms, error) {
	// Create the url for getting meetings
	endpoint := fmt.Sprintf("/metrics/issues/zoomrooms/%v", zoomroomId)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return Response_rooms{}, err
	}

	// Umarshal the response into struct
	var room Response_rooms
	err = json.Unmarshal(response, &room)
	if err != nil {
		log.Println(err)
		return Response_rooms{}, err
	}

	return room, nil
}
