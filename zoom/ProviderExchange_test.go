package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetListPeeringPhoneNumber(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"next_token": "6U86TxL8RZaoEWgM-xlEHw",
			"numbers": [
			  {
				"assigned": 0,
				"billing_reference_id": "1234abcd",
				"phone_number": "15556660100",
				"service_info": "Service info.",
				"sip_trunk_name": "example-carrier-trunk",
				"status": 1
			  }
			],
			"total_records": 1
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	peeringphonenumber, err := zoom.GetListPeeringPhoneNumber()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, peeringphonenumber, "results are empty")
	assert.Equal(t, len(peeringphonenumber.NextToken), 22, "NextToken is not matching")
	assert.Equal(t, peeringphonenumber.TotalRecords, 1, "TotalRecords is not matching")

}
