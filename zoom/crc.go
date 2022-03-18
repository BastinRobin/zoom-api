package zoom

import (
	"encoding/json"
	"log"
	"time"
)

type Response struct {
	CrcPortsUsage []struct {
		CrcPortsHourUsage []struct {
			Hour       string `json:"hour"`
			MaxUsage   int    `json:"max_usage"`
			TotalUsage int    `json:"total_usage"`
		} `json:"crc_ports_hour_usage"`
		DateTime time.Time `json:"date_time"`
	} `json:"crc_ports_usage"`
	From string `json:"from"`
	To   string `json:"to"`
}

func (z *Zoom) GetCRCUsage() (Response, error) {
	response, err := z.Request("/metrics/crc", "GET")
	if err != nil {
		log.Println(err)
		return Response{}, err
	}

	var crc Response

	// Unmarshal the response into Response struct
	err = json.Unmarshal(response, &crc)
	if err != nil {
		log.Println(err)
		return Response{}, err
	}

	return crc, nil

}
