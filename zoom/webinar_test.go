package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllWebinar(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"host": "Henry Chao",
			"duration": 12,
			"email": "hchao@example.com",
			"end_time": "2019-07-16T17:26:20Z",
			"has_3rd_party_audio": false,
			"has_archiving": false,
			"has_pstn": false,
			"has_recording": false,
			"has_screen_share": false,
			"has_sip": false,
			"has_video": false,
			"has_voip": false,
			"id": 33281536,
			"participants": 1,
			"start_time": "2019-07-16T17:14:39Z",
			"topic": "MyTestWebinar",
			"user_type": "Pro|Webinar1000",
			"uuid": "CJaaaaaaaEV6A=="
			}`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	response, err := zoom.GetAllWebinar("8722824397")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, response, "results are empty")

}
