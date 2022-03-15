package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCRCUsage(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"crc_ports_usage": [
			  {
				"crc_ports_hour_usage": [
				  {
					"hour": "sed",
					"max_usage": 8,
					"total_usage": 14
				  }
				],
				"date_time": "2019-05-01T15:13:39.424Z"
			  }
			],
			"from": "2019-04-03",
			"to": "2019-04-04"
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	response, err := zoom.GetCRCUsage()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, response, "results are empty")

}
