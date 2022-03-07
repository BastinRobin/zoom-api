package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWebinarDetails(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"from": "2019-08-24",
			"to": "2019-08-24",
			"next_page_token": "string",
			"page_count": 0,
			"page_size": 30,
			"total_records": 0,
			"webinars": [
			  {
				"host": "string",
				"custom_keys": [
				  {
					"key": "string",
					"value": "string"
				  }
				],
				"dept": "string",
				"duration": "string",
				"email": "string",
				"end_time": "2019-08-24T14:15:22Z",
				"has_3rd_party_audio": true,
				"has_archiving": true,
				"has_pstn": true,
				"has_recording": true,
				"has_screen_share": true,
				"has_sip": true,
				"has_video": true,
				"has_voip": true,
				"id": 0,
				"participants": 0,
				"start_time": "2019-08-24T14:15:22Z",
				"topic": "string",
				"user_type": "string",
				"uuid": "string",
				"audio_quality": "good",
				"video_quality": "good",
				"screen_share_quality": "good"
			  }
			]
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	webinar, err := zoom.GetWebinarDetails()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, webinar, "results are empty")
	assert.NotNil(t, webinar.TotalRecords, "No records presents")
	assert.Equal(t, len(webinar.Webinars), 1, "Total no of users not matching")
}
