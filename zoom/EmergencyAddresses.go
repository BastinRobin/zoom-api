package zoom

import (
	"encoding/json"
	"log"
)

type EmergencyAddresses struct {
	StateCode     string `json:"state_code"`
	Status        int    `json:"status"`
	Zip           string `json:"zip"`
	NextPageToken string `json:"next_page_token"`
	PageSize      int    `json:"page_size"`
	TotalRecords  int    `json:"total_records"`
	Addresses     Addresses
	Sites2        Sites2
	Owners        Owners
}
type Addresses []Address
type Address struct {
	AddressLine1 string `json:"address_line1"`
	City         string `json:"city"`
	Country      string `json:"country"`
	ID           string `json:"id"`
	IsDefault    bool   `json:"is_default"`
	Level        int    `json:"level"`
}
type Owners []Owners
type Owner struct {
	ExtensionNumber int    `json:"extension_number"`
	ID              string `json:"id"`
	Name            string `json:"name"`
}
type Sites2 []Site2
type Site2 struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Return the  list of  emergency addresses.
func (z *Zoom) GetListEmergencyAddresses() (EmergencyAddresses, error) {
	response, err := z.Request("/phone/emergency_addresses", "GET")
	if err != nil {
		log.Println(err)
		return EmergencyAddresses{}, err
	}

	var emergencyaddresses EmergencyAddresses

	// Unmarshal the response into EmergencyAddresses struct
	err = json.Unmarshal(response, &emergencyaddresses)
	if err != nil {
		log.Println(err)
		return EmergencyAddresses{}, err
	}
	//return emergencyaddresses struct
	return emergencyaddresses, nil

}
