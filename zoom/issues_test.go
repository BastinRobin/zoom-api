package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTop20Issues(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"from": "2019-08-24",
			"to": "2019-08-24",
			"total_records": 0,
			"zoom_rooms": [
			  {
				"id": "string",
				"issues_count": 0,
				"room_name": "string"
			  }
			]
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	response, err := zoom.GetTop20Issues()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, response, "results are empty")
	//assert.NotNil(t, response.TotalRecords, "No records presents")

}
