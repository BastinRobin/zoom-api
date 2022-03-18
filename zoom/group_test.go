package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGroup(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"id": "chfhfhhfh_TKikJIX0",
			"name": "My test group",
			"total_members": 0
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	group, err := zoom.GetGroup("chfhfhhfh_TKikJIX0")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, group, "results are empty")
	assert.Equal(t, group.TotalMembers, 0, "TotalMembers is not matching")
	assert.Equal(t, len(group.ID), 18, " group  id is not matching")

}
