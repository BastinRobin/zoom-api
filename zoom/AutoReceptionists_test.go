package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetListAutoReceptionists(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"auto_receptionists": [
			  {
				"cost_center": "test",
				"department": "test",
				"extension_id": "wh1KurOzThOXrtqTe0cXvw",
				"extension_number": 18300,
				"id": "3VnxCE-iQ1SrX5hKYTtSfg",
				"name": "test",
				"phone_numbers": [
				  {
					"id": "wLXTuEbgQ9i3maoawZ4Orw",
					"number": "+12058945692"
				  }
				],
				"site": {
				  "id": "4ziOo5WxRQeXdJSpYlDv0Q",
				  "name": "TestRange"
				}
			  }
			],
			"next_page_token": "eKTSRXixDFXKyiRa7gscqzgxVyv2E4daqU2",
			"page_size": 1,
			"total_records": 2
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	autoreceptionists, err := zoom.GetListAutoReceptionists()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, autoreceptionists, "results are empty")
	assert.Equal(t, len(autoreceptionists.NextPageToken), 35, "NextPageToken is not matching")
	assert.Equal(t, autoreceptionists.PageSize, 1, "PageSize is not matching")
	assert.Equal(t, autoreceptionists.TotalRecords, 2, "TotalRecords is not matching")

}
