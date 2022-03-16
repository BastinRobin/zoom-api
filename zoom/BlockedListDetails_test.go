package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBlockedListDetails(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"block_type": "inbound",
			"comment": "Blocked all numbers starting with 0001",
			"id": "7bR_Ix4KSJS_heuPPKc0Zw",
			"match_type": "prefix",
			"phone_number": "0001",
			"status": "active"
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	blockedlistdetails, err := zoom.GetBlockedListDetails("7bR_Ix4KSJS_heuPPKc0Zw")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, blockedlistdetails, "results are empty")
	assert.Equal(t, len(blockedlistdetails.PhoneNumber), 4, " PhoneNumber is not matching")
	assert.Equal(t, blockedlistdetails.ID, "7bR_Ix4KSJS_heuPPKc0Zw", "ID is not matching")
	assert.Equal(t, blockedlistdetails.Comment, "Blocked all numbers starting with 0001", "Comment is not matching")

}
