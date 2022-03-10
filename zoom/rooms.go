package zoom

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Rooms_With_Issues struct {
	From         string `json:"from"`
	PageCount    int    `json:"page_count"`
	PageSize     int    `json:"page_size"`
	To           string `json:"to"`
	TotalRecords int    `json:"total_records"`
	IssueDetails IssueDetails
}
type IssueDetails []IssueDetail
type IssueDetail []struct {
	Issue string    `json:"issue"`
	Time  time.Time `json:"time"`
}

// Get Specific Rooms with issue
func (z *Zoom) GetRoomsWithIssues(zoomroomId string) (Rooms_With_Issues, error) {
	// Create the url for getting meetings
	endpoint := fmt.Sprintf("/metrics/issues/zoomrooms/%v", zoomroomId)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return Rooms_With_Issues{}, err
	}

	// Umarshal the response into Rooms_with_Issues struct
	var room Rooms_With_Issues
	err = json.Unmarshal(response, &room)
	if err != nil {
		log.Println(err)
		return Rooms_With_Issues{}, err
	}
	//return Rooms_with_issues structure
	return room, nil
}
