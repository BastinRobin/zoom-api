package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetListBlockedLists(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"blocked_list": [
			  {
				"block_type": "inbound",
				"comment": "Blocked numbers starting with 777",
				"id": "7buyurtKc0Zw",
				"match_type": "prefix",
				"phone_number": "15550100",
				"status": "active"
			  },
			  {
				"block_type": "inbound",
				"comment": "Blocked an unknown caller",
				"id": "cYDtrtyrgyQw",
				"match_type": "phoneNumber",
				"phone_number": "15550101",
				"status": "active"
			  }
			],
			"next_page_token": "",
			"page_size": 30,
			"total_records": 2
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	blockedlist, err := zoom.GetListBlockedLists()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, blockedlist, "results are empty")
	assert.Equal(t, len(blockedlist.NextPageToken), 0, "NextPageToken is not matching")
	assert.Equal(t, blockedlist.PageSize, 30, "PageSize is not matching")
	assert.Equal(t, blockedlist.TotalRecords, 2, "TotalRecords is not matching")

}
