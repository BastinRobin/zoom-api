package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetChat(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"from": "2019-04-09",
			"next_page_token": "",
			"page_size": 1,
			"to": "2019-05-09",
			"users": [
			  {
				"audio_sent": 4,
				"code_sippet_sent": 2,
				"email": "example@example.com",
				"files_sent": 3,
				"giphys_sent": 2,
				"group_sent": 5,
				"images_sent": 5,
				"p2p_sent": 30,
				"text_sent": 8,
				"total_sent": 20,
				"user_id": "sdfjk393lklrf",
				"user_name": "culpa ipsum",
				"video_sent": 5
			  }
			]
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	chat, err := zoom.GetChat()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, chat, "results are empty")
}
