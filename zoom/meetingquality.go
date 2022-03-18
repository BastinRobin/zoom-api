package zoom

import (
	"encoding/json"
	"log"
)

type Quality struct {
	From    string `json:"from"`
	Quality []struct {
		Audio struct {
			Bad  int `json:"bad"`
			Fair int `json:"fair"`
			Good int `json:"good"`
			Poor int `json:"poor"`
		} `json:"audio"`
		ScreenShare struct {
			Bad  int `json:"bad"`
			Fair int `json:"fair"`
			Good int `json:"good"`
			Poor int `json:"poor"`
		} `json:"screen_share"`
		Video struct {
			Bad  int `json:"bad"`
			Fair int `json:"fair"`
			Good int `json:"good"`
			Poor int `json:"poor"`
		} `json:"video"`
	} `json:"quality"`
	To string `json:"to"`
}

// return meeting quality score information
func (z *Zoom) GetMeetingQualityScore() (Quality, error) {
	response, err := z.Request("/metrics/quality/", "GET")
	if err != nil {
		log.Println(err)
		return Quality{}, err
	}

	var quality Quality

	// Unmarshal the response into Quality struct
	err = json.Unmarshal(response, &quality)
	if err != nil {
		log.Println(err)
		return Quality{}, err
	}
	//return quality struct
	return quality, nil

}
