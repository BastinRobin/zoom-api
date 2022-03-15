package zoom

import (
	"encoding/json"
	"log"
)

type Issues struct {
	From         string `json:"from"`
	To           string `json:"to"`
	TotalRecords int    `json:"total_records"`
	ZoomRooms    []struct {
		ID          string `json:"id"`
		IssuesCount int    `json:"issues_count"`
		RoomName    string `json:"room_name"`
	} `json:"zoom_rooms"`
}

// Return the list of all Zoom Rooms in an account with issues
func (z *Zoom) GetTop20Issues() (Issues, error) {
	response, err := z.Request("/metrics/issues/zoomrooms", "GET")
	if err != nil {
		log.Println(err)
		return Issues{}, err
	}

	var issues Issues

	// Unmarshal the response into struct
	err = json.Unmarshal(response, &issues)
	if err != nil {
		log.Println(err)
		return Issues{}, err
	}

	return issues, nil

}
