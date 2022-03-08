package zoom

import (
	"encoding/json"
	"log"
)

type Zoom_Device struct {
	PageCount    int `json:"page_count"`
	PageNumber   int `json:"page_number"`
	PageSize     int `json:"page_size"`
	TotalRecords int `json:"total_records"`
	Devices      Devices
}
type Devices []Device
type Device struct {
	Encryption string `json:"encryption"`
	ID         string `json:"id"`
	IP         string `json:"ip"`
	Name       string `json:"name"`
	Protocol   string `json:"protocol"`
}

//Strucrure for Create SIP Devices
type CreateSIPDevice struct {
	Encryption string `json:"encryption"`
	IP         string `json:"ip"`
	Name       string `json:"name"`
	Protocol   string `json:"protocol"`
}

// Returns all SIPDevices in zoom account
func (z *Zoom) GetSIPDevice() (Zoom_Device, error) {
	sipdevices, err := z.Request("/h323/devices", "GET")
	if err != nil {
		log.Println(err)
		return Zoom_Device{}, err
	}

	var devices Zoom_Device

	// Unmarshal the response into Zoom_Device struct
	err = json.Unmarshal(sipdevices, &devices)
	if err != nil {
		log.Println(err)
		return Zoom_Device{}, err
	}
	//return device structure
	return devices, nil

}

func (z *Zoom) CreateSIPDevice() (CreateSIPDevice, error) {
	sipdevices, err := z.Request("/h323/devices", "POST")
	if err != nil {
		log.Println(err)
		return CreateSIPDevice{}, err
	}

	var devices CreateSIPDevice

	// Unmarshal the response into Zoom_Device struct
	err = json.Unmarshal(sipdevices, &devices)
	if err != nil {
		log.Println(err)
		return CreateSIPDevice{}, err
	}
	//return device structure
	return devices, nil

}
