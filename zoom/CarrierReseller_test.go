package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetListPhoneNumbers(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"carrier_reseller_numbers": [
			  {
				"assigned_status": "assigned",
				"carrier_code": 232,
				"phone_number": "12052334444",
				"status": "active",
				"sub_account_id": "343432",
				"sub_account_name": "xxxxx"
			  }
			],
			"next_page_token": "",
			"page_size": 30,
			"total_records": 1
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	listPhonenumbers, err := zoom.GetListPhoneNumbers()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, listPhonenumbers, "results are empty")
	assert.Equal(t, len(listPhonenumbers.NextPageToken), 0, "NextPageToken is not matching")
	assert.Equal(t, listPhonenumbers.PageSize, 30, "PageSize is not matching")
	assert.Equal(t, listPhonenumbers.TotalRecords, 1, "TotalRecords is not matching")

}
