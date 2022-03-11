package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMeetingPolls(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"polls": [
			  {
				"anonymous": false,
				"id": "YB33NcABCdg_g1AAAxTQ",
				"poll_type": 2,
				"questions": [
				  {
					"answer_required": false,
					"answers": [
					  "Extremely useful",
					  "Somewhat useful",
					  "Not useful at all"
					],
					"name": "How useful was this meeting?",
					"right_answers": [
					  "Not useful at all"
					],
					"type": "multiple"
				  }
				],
				"status": "notstart",
				"title": "Meeting Usefulness"
			  }
			],
			"total_records": 1
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	meetingpolls, err := zoom.GetMeetingPolls("8722824397")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, meetingpolls, "results are empty")
	assert.NotNil(t, meetingpolls.TotalRecords, 1, "No records presents")
}
