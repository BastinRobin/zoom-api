package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListBlockedList(t *testing.T) {

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

	blocked_list, err := zoom.ListBlockedList()
	if err != nil {
		t.Error(err)
	}
	assert.NotNil(t, blocked_list, "results are empty")
	assert.Equal(t, blocked_list.TotalRecords, 2, "Total records are not matching")
	assert.Equal(t, blocked_list.PageSize, 30, "pagesize not matched")
}
