package zoom

import (
	"encoding/json"
	"log"
)

type ListPhoneNumbers struct {
	NextPageToken string `json:"next_page_token"`
	PageSize      int    `json:"page_size"`
	TotalRecords  int    `json:"total_records"`
}

type CarrierResellerNumbers []CarrierResellerNumber
type CarrierResellerNumber struct {
	AssignedStatus string `json:"assigned_status"`
	CarrierCode    int    `json:"carrier_code"`
	PhoneNumber    string `json:"phone_number"`
	Status         string `json:"status"`
	SubAccountID   string `json:"sub_account_id"`
	SubAccountName string `json:"sub_account_name"`
}

func (z *Zoom) GetListPhoneNumbers() (ListPhoneNumbers, error) {
	response, err := z.Request("/phone/carrier_reseller/numbers", "GET")
	if err != nil {
		log.Println(err)
		return ListPhoneNumbers{}, err
	}

	var listPhonenumbers ListPhoneNumbers

	// Unmarshal the response into ListPhoneNumbers struct
	err = json.Unmarshal(response, &listPhonenumbers)
	if err != nil {
		log.Println(err)
		return ListPhoneNumbers{}, err
	}
	//return listPhonenumbers struct
	return listPhonenumbers, nil

}
