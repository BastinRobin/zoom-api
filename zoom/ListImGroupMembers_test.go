package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGroupMembers(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"members": [
			  {
				"email": "example@example.com",
				"first_name": "Rahul",
				"id": "dlfjdhq3430394",
				"last_name": "Ghimire",
				"type": "2"
			  }
			],
			"page_count": "1",
			"page_number": "1",
			"page_size": "1",
			"total_records": "1"
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	groupmembers, err := zoom.GetGroupMembers("dlfjdhq3430394")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, groupmembers, "results are empty")
	assert.Equal(t, len(groupmembers.TotalRecords), 1, "TotalRecords is not matching")
	assert.Equal(t, groupmembers.PageNumber, "1", "PageNumber is not matching")

}
