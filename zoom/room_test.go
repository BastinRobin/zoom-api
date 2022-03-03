package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRoomsWithIssues(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"from": "2019-02-28",
			"issue_details": [
			  {
				"issue": "Zoom room is offline",
				"time": "2019-03-07T11:17:00.956Z"
			  }
			],
			"page_count": 1,
			"page_size": 1,
			"to": "2019-03-28",
			"total_records": 1
		  }`))
	}))

	client := &Zoom{
		BaseUrl: testServer.URL,
	}

	rooms, err := client.GetMeeting("8722824397")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, rooms, "Empty meetings returned")
}
