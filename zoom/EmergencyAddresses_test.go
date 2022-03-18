package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetListEmergencyAddresses(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"emergency_addresses": [
			  {
				"address_line1": "55 Almaden Boulevard",
				"city": "SAN JOSE",
				"country": "US",
				"id": "CCc8zYT1SN60i7uDMzDbXA",
				"is_default": false,
				"level": 1,
				"owner": {
				  "extension_number": 1000000000,
				  "id": "q5C69v95SPKsZ5uUi-Xbcw",
				  "name": "someownername"
				},
				"site": {
				  "id": "8f71O6rWT8KFUGQmJIFAdQ",
				  "name": "somesitename"
				},
				"state_code": "CA",
				"status": 1,
				"zip": "95113"
			  }
			],
			"next_page_token": "",
			"page_size": 1,
			"total_records": 1
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	emergencyaddresses, err := zoom.GetListEmergencyAddresses()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, emergencyaddresses, "results are empty")
	assert.Equal(t, len(emergencyaddresses.NextPageToken), 0, "NextPageToken is not matching")
	assert.Equal(t, emergencyaddresses.PageSize, 1, "PageSize is not matching")
	assert.Equal(t, emergencyaddresses.TotalRecords, 1, "TotalRecords is not matching")

}
