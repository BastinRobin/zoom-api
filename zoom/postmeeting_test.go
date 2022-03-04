package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPostMeetingFeedback(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"next_page_token": "",
			"page_size": 1,
			"participants": [
			  {
				"date_time": "2021-02-21T18:48:06.423Z",
				"email": "example@example.com",
				"quality": "GOOD",
				"user_id": "aegr46312rum"
			  }
			]
		  }`))
	}))

	client := &Zoom{
		BaseUrl: testServer.URL,
	}

	postmeetings, err := client.GetPostMeetingFeedback("8722824397")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, postmeetings, "Empty meetings returned")
}
