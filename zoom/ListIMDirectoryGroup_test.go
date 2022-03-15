package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetListIMDirectory(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"groups": [
			  {
				"id": "sjvfkjfew34535",
				"name": "MyAdminDirectoryGroup",
				"search_by_account": "false",
				"search_by_domain": "true",
				"search_by_ma_account": "false",
				"total_members": "10",
				"type": "shared"
			  }
			],
			"page_count": "1",
			"page_number": "1",
			"page_size": "1",
			"total_records": "1"
		  }`))
	}))

	zoom := &Zoom{
		BaseUrl: testServer.URL,
	}

	list_imdirectory, err := zoom.GetListIMDirectory()
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, list_imdirectory, "results are empty")
	assert.Equal(t, len(list_imdirectory.Groups), 1, "Groups is not matching")
	assert.Equal(t, list_imdirectory.PageNumber, "1", "PageNumber is not matching")

}
