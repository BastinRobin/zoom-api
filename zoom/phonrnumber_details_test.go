package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListPhoneNumberDetails(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"assignee": {
			  "extension_number": 10000,
			  "name": "Main Auto Receptionist",
			  "type": "autoReceptionist"
			},
			"capability": [
			  "incoming",
			  "outgoing"
			],
			"carrier": {
			  "code": 13,
			  "name": "ContactCenter"
			},
			"display_name": "abc",
			"id": "Hfdgdfgdfg1ew",
			"location": "Milpitas,California,United States",
			"number": "+140000007",
			"number_type": "toll",
			"sip_group": {
			  "display_name": "RRRR",
			  "id": "8MhK7ea4Q4ihIQ4TD_g0kw"
			},
			"site": {
			  "id": "CEfdfdfgA",
			  "name": "Main Site"
			},
			"source": "internal",
			"status": "available"
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	phone_number_details, err := zoom.GetPhoneNumberDetails("Hfdgdfgdfg1ew")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, phone_number_details, "results are empty")
	assert.Contains(t, phone_number_details.Status, "available")
	assert.Equal(t, phone_number_details.NumberType, "toll", "number type not matching")
	assert.Contains(t, phone_number_details.Source, "internal")
}
