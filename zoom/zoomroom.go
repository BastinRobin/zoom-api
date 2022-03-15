package zoom

import (
	"encoding/json"
	"log"
	"time"
)

type List_ZoomRooms struct {
	PageCount    int `json:"page_count"`
	PageNumber   int `json:"page_number"`
	PageSize     int `json:"page_size"`
	TotalRecords int `json:"total_records"`
	ZoomRooms    ZoomRooms
}
type ZoomRooms []ZoomRoom
type ZoomRoom struct {
	AccountType   string    `json:"account_type"`
	CalendarName  string    `json:"calendar_name"`
	Camera        string    `json:"camera"`
	DeviceIP      string    `json:"device_ip"`
	Email         string    `json:"email"`
	Health        string    `json:"health"`
	ID            string    `json:"id"`
	Issues        []string  `json:"issues"`
	LastStartTime time.Time `json:"last_start_time"`
	Microphone    string    `json:"microphone"`
	RoomName      string    `json:"room_name"`
	Speaker       string    `json:"speaker"`
	Status        string    `json:"status"`
}

func (z *Zoom) GetListZoomRooms() (List_ZoomRooms, error) {
	response, err := z.Request("/metrics/zoomrooms", "GET")
	if err != nil {
		log.Println(err)
		return List_ZoomRooms{}, err
	}

	var zoomrooms List_ZoomRooms

	// Unmarshal the response into ListZoomRooms struct
	err = json.Unmarshal(response, &zoomrooms)
	if err != nil {
		log.Println(err)
		return List_ZoomRooms{}, err
	}
	//return zoomrooms struct
	return zoomrooms, nil

}
