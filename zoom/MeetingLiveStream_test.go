package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLiveStream(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"page_url": "https://somecompany.com/livestream/123",
			"stream_key": "Contact It@somecompany.com",
			"stream_url": "https://somecompany.com/livestream"
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	livestream, err := zoom.GetLiveStream("8722824397")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, livestream, "results are empty")
	assert.Equal(t, len(livestream.StreamKey), 26, "StreamKey is not matching")
	assert.Equal(t, livestream.PageURL, "https://somecompany.com/livestream/123", "PageURl is not matching")
	assert.Equal(t, livestream.StreamURL, "https://somecompany.com/livestream", "StreamUR is not matching")

}
