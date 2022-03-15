package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRecordingDetails(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"next_page_token": "string",
			"page_count": "integer",
			"page_size": "integer",
			"participants": [
			  {
				"details": [
				  {
					"content": "string",
					"end_time": "string",
					"start_time": "string"
				  }
				],
				"id": "string",
				"user_id": "string",
				"user_name": "string"
			  }
			],
			"total_records": "integer"
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	response, err := zoom.GetRecordingDetails("8722824397")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, response, "results are empty")
	//assert.NotNil(t, response.TotalRecords, "No records presents")

}
