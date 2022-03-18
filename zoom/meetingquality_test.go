package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMeetingQualityScore(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"from": "2019-02-28",
			"quality": [
			{
			"audio": {
			"bad": 0,
			"fair": 2,
			"good": 13,
			"poor": 1
			},
			"screen_share": {
			"bad": 3,
			"fair": 4,
			"good": 6,
			"poor": 3
			},
			"video": {
			"bad": 0,
			"fair": 0,
			"good": 16,
			"poor": 0
			}
			}
			],
			"to": "2019-03-28"
			}`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	response, err := zoom.GetMeetingQualityScore()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, response, "results are empty")

}
