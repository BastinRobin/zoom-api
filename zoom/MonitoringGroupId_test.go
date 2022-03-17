package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMonitoringGroupId(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
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
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	monitoringgroupid, err := zoom.GetMonitoringGroupId("8f71O6rWT8KFUGQmJIFAdQ")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, monitoringgroupid, "results are empty")
	assert.Equal(t, len(monitoringgroupid.ID), 22, "ID is not matching")
	assert.Equal(t, monitoringgroupid.Name, "test", "Name is not matching")

}
