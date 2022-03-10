package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetListGroup(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"groups": [
			  {
				"id": "hFK_GtF_e_TaVA808",
				"name": "ipsum",
				"total_members": 10
			  },
			  {
				"id": "TaVA8QKik_1233",
				"name": "awesomegroup",
				"total_members": 0
			  }
			],
			"total_records": 2
		  }`))
	}))

	client := &Zoom{
		BaseUrl: testServer.URL,
	}

	group, err := client.GetListGroups()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, group, "Empty meetings returned")
	assert.Equal(t, group.TotalRecords, 2, "TotalRecords is not matching")
	assert.Equal(t, len(group.Groups), 2, "Total no of groups not matching")
}
