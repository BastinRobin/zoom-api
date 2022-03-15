package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAudioItem(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"audio_id": "DP-6B4CeTjyo-FYfsFXdXQ",
			"name": "Teeeee.mp3",
			"play_url": "https://file.example.com/file?jwt=eyJhbGciOiJIUzI1NiJ9.eyJleHAiOjE2MzY3MjY0MDEsImlzcyI6ImNyb3NzZmlsZSIsImF1ZCI6ImZpbGUiLCJkaWciOiIzM2MyZGFhZjQ2ZDQ2MzFiNGJkMWUzZmRmYmI5OTBjOTUzNTEwZmE5ZDYxZjYyNDAyNGY5OWZiYmY5ZmZlMWU4In0.kPwNFT1C_twZHl3CTeyaiOLhxmJBcHb__SvDBmgGpiQ&mode=download&path=zoomfs%3A%2F%2Fpbx-voice%2F%2Fprompt%2FNNNiWOl7SSmO-qXFOSXPMA%2FhcAjVmo0SVmdSvW2Sm7VrA.mp3"
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	audio_items, err := zoom.GetAudioItem("DP-6B4CeTjyo-FYfsFXdXQ")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, audio_items, "results are empty")
	assert.Equal(t, audio_items.Name, "Teeeee.mp3", "Names are not matching")
	assert.Equal(t, audio_items.PlayURL, "https://file.example.com/file?jwt=eyJhbGciOiJIUzI1NiJ9.eyJleHAiOjE2MzY3MjY0MDEsImlzcyI6ImNyb3NzZmlsZSIsImF1ZCI6ImZpbGUiLCJkaWciOiIzM2MyZGFhZjQ2ZDQ2MzFiNGJkMWUzZmRmYmI5OTBjOTUzNTEwZmE5ZDYxZjYyNDAyNGY5OWZiYmY5ZmZlMWU4In0.kPwNFT1C_twZHl3CTeyaiOLhxmJBcHb__SvDBmgGpiQ&mode=download&path=zoomfs%3A%2F%2Fpbx-voice%2F%2Fprompt%2FNNNiWOl7SSmO-qXFOSXPMA%2FhcAjVmo0SVmdSvW2Sm7VrA.mp3", "PlayURL are not matching")

}
