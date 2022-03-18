package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGroupAdmins(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"admins": [
			  {
				"email": "",
				"id": "3542342",
				"name": "Ghale"
			  }
			],
			"next_page_token": "svjsecsjaiowef",
			"page_size": 1,
			"total_records": 1
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	groupadmins, err := zoom.GetGroup("3542342")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, groupadmins, "results are empty")
	assert.Equal(t, len(groupadmins.ID), 0, "Group Id is not matching")
	assert.Equal(t, groupadmins.TotalMembers, 0, "TotalMembers is not matching")

}
