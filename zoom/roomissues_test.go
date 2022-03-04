package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet25IssuesOfZoomRooms(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"from": "2019-08-15",
			"room_issues": [
			  {
				"issue_name": "Controller disconnected",
				"zoom_rooms_count": 1
			  },
			  {
				"issue_name": "Controller is not charging",
				"zoom_rooms_count": 1
			  }
			],
			"to": "2019-09-15",
			"total_records": 2
		  }`))
	}))
	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	response, err := zoom.Get25IssuesOfZoomRooms()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, response, "results are empty")
	assert.NotNil(t, response.TotalRecords, "No records presents")
}
