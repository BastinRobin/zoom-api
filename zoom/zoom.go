package zoom

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Zoom struct {
	BaseUrl string // https://api.zoom.com/
	Headers map[string][]string
}

// Attach functions to the given Zoom struct
func (z *Zoom) Request(url string, method string) ([]byte, error) {

	// BaseURL : https://api.zoom.com
	// url : /users or /meeting/{meeting_id}
	// https://api.zoom.com/users

	// Form the complete url for the struct
	endpoint := fmt.Sprintf("%v%v", z.BaseUrl, url)

	// Create a request object
	req, err := http.NewRequest(method, endpoint, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Attach the headers into request object
	req.Header = z.Headers

	// Execute the API Call
	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Convert the response object into byte so we can convert to struct
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return data, nil
}
