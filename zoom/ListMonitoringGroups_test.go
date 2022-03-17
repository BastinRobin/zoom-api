package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetListMonitoringGroups(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"monitoring_groups": {
			  "id": "Ct6mcZYMSuWYug46puWOQg",
			  "monitor_members_count": 2,
			  "monitored_members_count": 2,
			  "monitoring_privileges": [
				"listen",
				"whisper",
				"barge",
				"take_over"
			  ],
			  "name": "test",
			  "prompt": true,
			  "site": {
				"id": "8f71O6rWT8KFUGQmJIFAdQ",
				"name": "Main Site"
			  },
			  "type": 1
			},
			"next_page_token": "abcD3944YsoYPB12",
			"page_size": 1,
			"total_records": 6
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	listmonitoringroups, err := zoom.GetListMonitoringGroups()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, listmonitoringroups, "results are empty")
	assert.Equal(t, len(listmonitoringroups.NextPageToken), 16, "NextPageToken is not matching")
	assert.Equal(t, listmonitoringroups.PageSize, 1, "PageSize is not matching")

}
