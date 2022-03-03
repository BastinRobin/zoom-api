package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMeeting(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"agenda": "API overview",
			"created_at": "2019-09-09T15:54:24Z",
			"duration": 60,
			"host_id": "ABcdofjdogh11111",
			"id": 1234555466,
			"join_url": "https://zoom.us/j/1234555466",
			"settings": {
			  "alternative_hosts": "example@example.com",
			  "approval_type": 2,
			  "audio": "both",
			  "auto_recording": "local",
			  "close_registration": false,
			  "cn_meeting": false,
			  "enforce_login": false,
			  "enforce_login_domains": "example.com",
			  "global_dial_in_countries": [
				"US"
			  ],
			  "global_dial_in_numbers": [
				{
				  "city": "New York",
				  "country": "US",
				  "country_name": "US",
				  "number": "+1 000011100",
				  "type": "toll"
				},
				{
				  "city": "San Jose",
				  "country": "US",
				  "country_name": "US",
				  "number": "+1 6699006833",
				  "type": "toll"
				},
				{
				  "city": "San Jose",
				  "country": "US",
				  "country_name": "US",
				  "number": "+1 221122112211",
				  "type": "toll"
				}
			  ],
			  "host_video": false,
			  "in_meeting": false,
			  "join_before_host": true,
			  "mute_upon_entry": false,
			  "participant_video": false,
			  "registrants_confirmation_email": true,
			  "registrants_email_notification": true,
			  "use_pmi": false,
			  "waiting_room": false,
			  "watermark": false
			},
			"start_time": "2019-08-30T22:00:00Z",
			"start_url": "https://zoom.us/1234555466/cdknfdffgggdfg4MDUxNjY0LCJpYXQiOjE1NjgwNDQ0NjQsImFpZCI6IjRBOWR2QkRqVHphd2J0amxoejNQZ1EiLCJjaWQiOiIifQ.Pz_msGuQwtylTtYQ",
			"status": "waiting",
			"timezone": "America/New_York",
			"topic": "My API Test",
			"type": 2,
			"uuid": "iAABBBcccdddd7A=="
		  }`))
	}))

	client := &Zoom{
		BaseUrl: testServer.URL,
	}

	meetings, err := client.GetMeeting("8722824397")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, meetings, "Empty meetings returned")
	assert.Equal(t, meetings.UUID, "iAABBBcccdddd7A==", "UUID is not matching")

}
