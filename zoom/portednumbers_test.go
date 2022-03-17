package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListPortedNumber(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"next_page_token": "w8jQ7vUYK2VIKMeuurezwkFDhlVLPvcw8L2",
			"page_size": 1,
			"ported_numbers": [
			  {
				"numbers": [
				  "+12058945751",
				  "+12058945752"
				],
				"order_id": "2021080201485902918",
				"replacing_numbers": [
				  {
					"source_number": "+12058945752",
					"target_number": "+12058945755"
				  },
				  {
					"source_number": "+12058945751",
					"target_number": "+12058945753"
				  }
				],
				"status": "Processing",
				"submitted_date_time": "2021-08-01T17:48:59Z"
			  }
			],
			"total_records": 2
		  }`))
	}))

	client := &Zoom{
		BaseUrl: testServer.URL,
	}

	ported_numbers, err := client.ListPortedNumbers()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, ported_numbers, "Empty meetings returned")
	assert.Equal(t, ported_numbers.TotalRecords, 2, "Total Records Not Matching")
	assert.Contains(t, ported_numbers.NextPageToken, "w8jQ7vUYK2VIKMeuurezwkFDhlVLPvcw8L2")
	//assert.Equal(t, settings.CallingPlans[0].Subscribed, 10, "Subscriptions are not matching")
	// assert.Equal(t, settings.CallingPlans[0].Type, 100, "Types are not matching")
}
