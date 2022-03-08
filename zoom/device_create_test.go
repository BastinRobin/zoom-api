package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateSIPDevice(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"encryption": "auto",
			"ip": "string",
			"name": "string",
			"protocol": "H.323"
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	device, err := zoom.CreateSIPDevice()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, device, "results are empty")
	assert.NotNil(t, device.IP, 987979786876, "IP incorrect")
}
