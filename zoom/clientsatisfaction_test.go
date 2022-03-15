package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClientMeetingSatisfaction(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"client_satisfaction": [
			  {
				"date": "2019-08-05",
				"good_count": 5,
				"none_count": 0,
				"not_good_count": 1,
				"satisfaction_percent": 100
			  },
			  {
				"date": "2019-08-06",
				"good_count": 0,
				"none_count": 0,
				"not_good_count": 0,
				"satisfaction_percent": 100
			  }
			],
			"from": "2019-08-05",
			"to": "2019-09-05",
			"total_records": 30
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	response, err := zoom.GetClientMeetingSatisfaction()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, response, "results are empty")

}
