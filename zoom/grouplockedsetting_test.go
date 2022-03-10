package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLockedSetting(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"id": "chfhfhhfh_TKikJIX0",
			"name": "My test group",
			"total_members": 0
		  }`))
	}))

	client := &Zoom{
		BaseUrl: testServer.URL,
	}

	group_locked_setting, err := client.GetLockedSetting("chfhfhhfh_TKikJIX0")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, group_locked_setting, "Empty meetings returned")
	assert.Equal(t, group_locked_setting.TotalMembers, 0, "TotalMembers is not matching")
	assert.Equal(t, group_locked_setting.Name, "My test group", "Names are not matching")
}
