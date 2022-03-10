package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetListGroupMembers(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"members": [
			  {
				"email": "",
				"first_name": "Ram",
				"id": "3542342",
				"last_name": "Ghale",
				"type": 1
			  }
			],
			"page_count": 1,
			"page_number": 1,
			"page_size": 1,
			"total_records": 1
		  }`))
	}))

	client := &Zoom{
		BaseUrl: testServer.URL,
	}

	group_members, err := client.GetListGroupMembers("chfhfhhfh_TKikJIX0")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, group_members, "Empty meetings returned")
	assert.Equal(t, group_members.TotalRecords, 1, "TotalRecords is not matching")
	assert.Equal(t, group_members.PageCount, 1, "Page count is not matching")
	assert.Equal(t, group_members.Members[0].FirstName, "Ram", "Names are not matching")
}
