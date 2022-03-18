package zoom

import (
	"encoding/json"
	"log"
)

type PeeringPhoneNumber struct {
	NextToken    string `json:"next_token"`
	TotalRecords int    `json:"total_records"`
}
type Numbers []Number
type Number struct {
	Assigned           int    `json:"assigned"`
	BillingReferenceID string `json:"billing_reference_id"`
	PhoneNumber        string `json:"phone_number"`
	ServiceInfo        string `json:"service_info"`
	SipTrunkName       string `json:"sip_trunk_name"`
	Status             int    `json:"status"`
}

// Return the list of peering phone numbers
func (z *Zoom) GetListPeeringPhoneNumber() (PeeringPhoneNumber, error) {
	response, err := z.Request("/phone/peering/numbers", "GET")
	if err != nil {
		log.Println(err)
		return PeeringPhoneNumber{}, err
	}

	var peeringphonenumber PeeringPhoneNumber

	// Unmarshal the response into peeringphonenumber struct
	err = json.Unmarshal(response, &peeringphonenumber)
	if err != nil {
		log.Println(err)
		return PeeringPhoneNumber{}, err
	}
	//return peeringphonenumber struct
	return peeringphonenumber, nil

}
