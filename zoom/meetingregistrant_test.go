package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMeetingRegistrant(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"page_count": 1,
			"page_number": 1,
			"page_size": 1,
			"registrants": [
			  {
				"address": "8 Rue du Nom Fictif",
				"city": "Paris",
				"comments": "Love using Zoom APIs",
				"country": "France",
				"create_time": "2012-04-14T16:55:30.231Z",
				"custom_questions": [
				  {
					"title": "Did you enjoy the registration process?",
					"value": "Yes, alot."
				  },
				  {
					"title": "Would you like to register for our next meeting?",
					"value": "Absolutely."
				  }
				],
				"email": "example@example.com",
				"first_name": "Tim",
				"id": "zjkfsdfjdfhg",
				"industry": "Tech",
				"job_title": "Developer",
				"join_url": "https://success.zoom.us/j/0000000000000",
				"last_name": "S.",
				"no_of_employees": "1-20",
				"org": "Growth",
				"phone": "5550100",
				"purchasing_time_frame": "Within a month",
				"role_in_purchase_process": "Not involved",
				"state": "Ile-de-France",
				"status": "approved",
				"zip": "314"
			  }
			],
			"total_records": 1
		  }`))
	}))

	client := &Zoom{
		BaseUrl: testServer.URL,
	}

	meetings_registrant, err := client.GetMeetingRegistrants("8722824397")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, meetings_registrant, "Empty meetings returned")
	assert.Equal(t, meetings_registrant.TotalRecords, 1, "No records presents")
	assert.Equal(t, meetings_registrant.PageCount, 1, "Total count are not matching")

}
