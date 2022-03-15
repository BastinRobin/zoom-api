package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetListMeetingParticipantsQoS(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"next_page_token": "string",
			"page_count": "integer [int64]",
			"page_size": "integer",
			"participants": [
			  {
				"device": "Android",
				"domain": "user-android",
				"harddisk_id": "",
				"ip_address": "192.0.2.0",
				"join_time": "2021-06-24T20:00:00Z",
				"leave_time": "2021-06-24T20:00:00Z",
				"location": "San Jose (US)",
				"mac_addr": "",
				"pc_name": "User's Phone",
				"user_id": "1670000000",
				"user_name": "User",
				"user_qos": [
				  {
					"as_input": {
					  "avg_loss": "0.0%",
					  "bitrate": "70 kbps",
					  "frame_rate": "",
					  "jitter": "8 ms",
					  "latency": "135 ms",
					  "max_loss": "0.0%",
					  "resolution": "0*0"
					},
					"as_output": {
					  "avg_loss": "0.0%",
					  "bitrate": "70 kbps",
					  "frame_rate": "",
					  "jitter": "8 ms",
					  "latency": "135 ms",
					  "max_loss": "0.0%",
					  "resolution": "0*0"
					},
					"audio_input": {
					  "avg_loss": "0.3%",
					  "bitrate": "23 kbps",
					  "jitter": "6 ms",
					  "latency": "126 ms",
					  "max_loss": "1.9%"
					},
					"audio_output": {
					  "avg_loss": "0.0%",
					  "bitrate": "63 kbps",
					  "jitter": "6 ms",
					  "latency": "134 ms",
					  "max_loss": "0.0%"
					},
					"cpu_usage": {
					  "system_max_cpu_usage": "40%",
					  "zoom_avg_cpu_usage": "12%",
					  "zoom_max_cpu_usage": "18%",
					  "zoom_min_cpu_usage": "8%"
					},
					"date_time": "2021-06-24T20:00:00Z",
					"video_input": {
					  "avg_loss": "0.0%",
					  "bitrate": "1055 kbps",
					  "frame_rate": "12 fps",
					  "jitter": "11 ms",
					  "latency": "129 ms",
					  "max_loss": "4.9%",
					  "resolution": "1280*720"
					},
					"video_output": {
					  "avg_loss": "0.0%",
					  "bitrate": "673 kbps",
					  "frame_rate": "22 fps",
					  "jitter": "11 ms",
					  "latency": "135 ms",
					  "max_loss": "0.0%",
					  "resolution": "640*360"
					}
				  }
				],
				"version": "4.4.55383.0726"
			  }
			],
			"total_records": "integer [int64]"
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	response, err := zoom.GetListMeetingParticipantsQoS("8722824397")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, response, "results are empty")
	assert.NotNil(t, response.TotalRecords, "No records presents")
	assert.NotNil(t, response.PageCount, 1.4557887, "Page count are not matching")
}
