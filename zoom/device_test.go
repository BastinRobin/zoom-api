package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSIPDevice(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"devices": [
			  {
				"encryption": "auto",
				"id": "abceHewahkrehwiK",
				"ip": "127.0.0.1",
				"name": "api_test_20190508",
				"protocol": "H.323"
			  }
			],
			"page_count": 1,
			"page_number": 1,
			"page_size": 1,
			"total_records": 1
		  }`))
	}))

	client := &Zoom{
		BaseUrl: testServer.URL,
	}

	devices, err := client.GetSIPDevice()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, devices, "Empty meetings returned")
	assert.Equal(t, devices.TotalRecords, 1, "TotalRecords is not matching")
	assert.Equal(t, len(devices.Devices), 1, "Total no of devices not matching")
}
