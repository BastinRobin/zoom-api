package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListPhoneNumber(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"next_page_token": "",
			"page_size": 30,
			"phone_numbers": [
			  {
				"assignee": {
				  "extension_number": 12,
				  "id": "cgfdgfghghim",
				  "name": "Peter Jenner",
				  "type": "user"
				},
				"capability": [
				  "incoming",
				  "outgoing"
				],
				"display_name": "abc",
				"id": "execvbfgbgr",
				"location": "Milpitas,California,United States",
				"number": "0000111100",
				"number_type": "tollfree",
				"site": {
				  "id": "sdfsdfgrg",
				  "name": "SF office"
				},
				"source": "internal",
				"status": "pending"
			  },
			  {
				"assignee": {
				  "extension_number": 1,
				  "id": "dfgdfghdfhgh",
				  "name": "Receptionist",
				  "type": "autoReceptionist"
				},
				"id": "fdgfdgfdh",
				"location": "San Jose,California,United States",
				"number": "111111111",
				"number_type": "toll",
				"site": {
				  "id": "jhdfsdghfdg",
				  "name": "San Jose office"
				},
				"source": "external",
				"status": "available"
			  }
			],
			"total_records": 2
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	phone_number, err := zoom.GetUsers()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, phone_number, "results are empty")
	assert.NotNil(t, phone_number.TotalRecords, 2, "No records presents")
}
