package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPostWebinarFeedback(t *testing.T) {

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

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	response, err := zoom.GetPostWebinarFeedback("8722824397")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, response, "results are empty")
	assert.NotNil(t, response.PageSize, 1, "No pagesize presents")
}
