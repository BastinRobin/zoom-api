package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMeetingInvitation(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"invitation": "Shrijana G is inviting you to a scheduled Zoom meeting.\r\n\r\nTopic: MyTestMeeting\r\nTime: Jul 31, 2019 04:00 PM Pacific Time (US and Canada)\r\n\r\nJoin Zoom Meeting\r\nhttps://zoom.us/j/000000\r\n\r\nOne tap mobile\r\n+000000"
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	meeting_invitation, err := zoom.GetMeetingInvitation("8722824397")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, meeting_invitation, "results are empty")
	assert.Equal(t, len(meeting_invitation.Invitation), 211, "Invitation is not matching")

}
