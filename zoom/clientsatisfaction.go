package zoom

import (
	"encoding/json"
	"log"
)

type Client struct {
	From                string `json:"from"`
	To                  string `json:"to"`
	TotalRecords        int    `json:"total_records"`
	ClientSatisfactions ClientSatisfactions
}
type ClientSatisfactions []ClientSatisfaction
type ClientSatisfaction struct {
	Date      string `json:"date"`
	GoodCount int    `json:"good_count"`
	NoneCount int    `json:"none_count"`
}

// return the list of attendees meeting satisfaction
func (z *Zoom) GetClientMeetingSatisfaction() (Client, error) {
	response, err := z.Request("/metrics/client/satisfaction", "GET")
	if err != nil {
		log.Println(err)
		return Client{}, err
	}

	var client Client

	// Unmarshal the response into  struct
	err = json.Unmarshal(response, &client)
	if err != nil {
		log.Println(err)
		return Client{}, err
	}
	// return client struct
	return client, nil

}
