package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"page_count": 1,
			"page_number": 1,
			"page_size": 30,
			"total_records": 1,
			"users": [
			  {
				"created_at": "2018-11-15T01:10:08Z",
				"dept": "",
				"email": "example@example.com",
				"employee_unique_id": "dddaaaa",
				"first_name": "Taylor",
				"id": "z8yAAAAA8bbbQ",
				"im_group_ids": [
				  "GroupName"
				],
				"last_client_version": "4.4.55383.0716(android)",
				"last_login_time": "2019-09-13T21:08:52Z",
				"last_name": "Kim",
				"pic_url": "https://example.com/photo.jpg",
				"pmi": 111111111,
				"status": "active",
				"timezone": "America/Los_Angeles",
				"type": 2,
				"verified": 1
			  }
			]
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	users, err := zoom.GetUsers()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, users, "results are empty")
	assert.NotNil(t, users.TotalRecords, "No records presents")
	assert.Equal(t, len(users.Users), 1, "Total no of users not matching")
	assert.Equal(t, users.Users[0].FirstName, "Taylor", "Names are not matching")
	assert.Equal(t, users.Users[0].PMI, int64(111111111), "PMI are not matching")
}
