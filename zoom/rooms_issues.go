package zoom

import (
	"encoding/json"
	"log"
)

type Top25_zoom_roomsissues struct {
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

// Return the list of 25 issues in zoom Room
func (z *Zoom) Get25IssuesOfZoomRooms() (Top25_zoom_roomsissues, error) {
	roomissue, err := z.Request("/metrics/zoomrooms/issues", "GET")
	if err != nil {
		log.Println(err)
		return Top25_zoom_roomsissues{}, err
	}

	var issues Top25_zoom_roomsissues

	// Unmarshal the response into Top25_zoom_roomsissues struct
	err = json.Unmarshal(roomissue, &issues)
	if err != nil {
		log.Println(err)
		return Top25_zoom_roomsissues{}, err
	}
	//return Top25_zoom_roomsissue structure
	return issues, nil

}
