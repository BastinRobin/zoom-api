package zoom

import (
	"encoding/json"
	"log"
)

type Response_roomsissue struct {
	From string `json:"from"`
	//To string `json:"to"`
	TotalRecords int `json:"total_records"`
	RoomsIssues  RoomsIssues
}
type RoomsIssues []RoomsIssue
type RoomsIssue struct {
	IssueName      string `json:"issue_name"`
	ZoomRoomsCount int    `json:"zoom_rooms_count"`
}

// Return the list of all issues in zoom Room
func (z *Zoom) Get25IssuesOfZoomRooms() (Response_roomsissue, error) {
	roomissue, err := z.Request("/metrics/zoomrooms/issues", "GET")
	if err != nil {
		log.Println(err)
		return Response_roomsissue{}, err
	}

	var response Response_roomsissue

	// Unmarshal the response into Response struct
	err = json.Unmarshal(roomissue, &response)
	if err != nil {
		log.Println(err)
		return Response_roomsissue{}, err
	}

	return response, nil

}
