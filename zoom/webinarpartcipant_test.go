package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetWebinarPartcipants(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"next_page_token": "",
			"page_count": 1,
			"page_size": 30,
			"participants": [
			  {
				"camera": "FaceTime HD Camera",
				"connection_type": "P2P",
				"data_center": "SC",
				"device": "WIN",
				"domain": "Dojo-workspace",
				"from_sip_uri": "sip:sipp@10.100.112.140:11029",
				"full_data_center": "United States;United States (US03_SC CRC)",
				"harddisk_id": "sed proident in",
				"id": "d52f19c548b88490b5d16fcbd38",
				"ip_address": "127.0.0.1",
				"join_time": "2019-09-07T13:15:02.837Z",
				"leave_reason": "Dojo left the meeting.<br>Reason: Host ended the meeting.",
				"leave_time": "2019-09-07T13:15:09.837Z",
				"location": "New York",
				"mac_addr": " 00:0a:95:9d:68:16",
				"microphone": "Plantronics BT600",
				"network_type": "Wired",
				"pc_name": "dojo's pc",
				"recording": false,
				"role": "panelist",
				"share_application": false,
				"share_desktop": true,
				"share_whiteboard": true,
				"sip_uri": "sip:sipp@10.100.112.140:11029",
				"speaker": "Plantronics BT600",
				"status": "in_waiting_room",
				"user_id": "32dsfsd4g5gd",
				"user_name": "dojo",
				"version": "4.4.55383.0716"
			  }
			],
			"total_records": 1
		  }`))
	}))
	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	WebinarParcipants, err := zoom.GetWebinarPartcipants("8722824397")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, WebinarParcipants, "results are empty")
	assert.NotNil(t, WebinarParcipants.TotalRecords, "No records presents")
	assert.Equal(t, len(WebinarParcipants.Participants), 1, "Total no of partcipants not matching")
}
