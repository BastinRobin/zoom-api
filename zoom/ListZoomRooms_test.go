package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetListZoomRooms(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"page_count": 1,
			"page_number": 1,
			"page_size": 30,
			"total_records": 1,
			"zoom_rooms": [
			  {
				"account_type": "Work Email",
				"calendar_name": "example.cal",
				"camera": "Integrated Webcam",
				"device_ip": "Computer : 192.0.2.0",
				"email": "example@example.com",
				"health": "critical",
				"id": "EbgjgjhghZY9wh0A",
				"issues": [
				  "Zoom room is offline"
				],
				"last_start_time": "2019-08-29T16:37:07Z",
				"microphone": "Microphone Array (Realtek Audio)",
				"room_name": "testZoomRoom",
				"speaker": "Speakers / Headphones (Realtek Audio)",
				"status": "Offline"
			  }
			]
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	zoomrooms, err := zoom.GetListZoomRooms()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, zoomrooms, "results are empty")
	assert.NotNil(t, zoomrooms.TotalRecords, "No records presents")
	assert.Equal(t, len(zoomrooms.ZoomRooms), 0, "ZoomRooms are not matching")

}
