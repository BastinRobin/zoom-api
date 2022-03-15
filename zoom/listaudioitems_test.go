package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListGetAudioItem(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"audios": [
			  {
				"audio_id": "DP-6B4CeTjyo-FYfsFXdXQ",
				"name": "Teeeee.mp3"
			  },
			  {
				"audio_id": "K3xnsN2kSpSODzIRmDj0Rg",
				"name": "TestAdd01.mp3"
			  }
			]
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	audios, err := zoom.GetListAudioItem("")
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, audios, "results are empty")
	assert.Equal(t, audios.Audios.Name, "TestAdd01.mp3", "Names are not matching")
}
