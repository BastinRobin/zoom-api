package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListCallingPlans(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"calling_plans": [
			  {
				"assigned": 1,
				"available": 9,
				"name": "US/Canada metered calling plan",
				"subscribed": 10,
				"type": 100
			  },
			  {
				"assigned": 2,
				"available": 2,
				"name": "Australia/New Zealand metered calling plan",
				"subscribed": 4,
				"type": 101
			  }
			]
		  }`))
	}))

	client := &Zoom{
		BaseUrl: testServer.URL,
	}

	settings, err := client.ListCallingPlans()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, settings, "Empty meetings returned")
	// assert.Equal(t, settings.CallingPlans[0].Name, "US/Canada metered calling plan", "names not matching")
	//assert.Equal(t, settings.CallingPlans[0].Subscribed, 10, "Subscriptions are not matching")
	// assert.Equal(t, settings.CallingPlans[0].Type, 100, "Types are not matching")
}
