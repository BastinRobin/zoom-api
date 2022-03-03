package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetClientFeedback(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"client_feedbacks": [
			  {
				"feedback_id": "53532100000",
				"feedback_name": "Poor audio quality.",
				"participants_count": 1
			  },
			  {
				"feedback_id": "53532100000",
				"feedback_name": "They could not hear us.",
				"participants_count": 2
			  }
			],
			"from": "2013-03-16",
			"to": "2013-04-16",
			"total_records": 2
		  }`))
	}))
	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	response, err := zoom.GetChat()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, response, "results are empty")
	//assert.NotNil(t, response.TotalRecords, "No records presents")
	//assert.Equal(t, len(chats.Users), 1, "Total no of users not matching")
	//assert.Equal(t, chats.Users[0].FirstName, "Taylor", "Names are not matching")
	//assert.Equal(t, chats.Users[0].PMI, int64(111111111), "PMI are not matching")
}
