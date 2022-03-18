package zoom

import (
	"encoding/json"
	"log"
)

type ListAutoReceptionists struct {
	NextPageToken     string `json:"next_page_token"`
	PageSize          int    `json:"page_size"`
	TotalRecords      int    `json:"total_records"`
	AutoReceptionists AutoReceptionists
	PhoneNumbers      PhoneNumbers
	Sites             Sites
}
type AutoReceptionists []AutoReceptionist
type AutoReceptionist struct {
	CostCenter      string `json:"cost_center"`
	Department      string `json:"department"`
	ExtensionID     string `json:"extension_id"`
	ExtensionNumber int    `json:"extension_number"`
	ID              string `json:"id"`
	Name            string `json:"name"`
}

type PhoneNumbers []PhoneNumber
type PhoneNumber struct {
	ID     string `json:"id"`
	Number string `json:"number"`
}
type Sites []Site
type Site struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//return the list auto receptionists.
func (z *Zoom) GetListAutoReceptionists() (ListAutoReceptionists, error) {
	response, err := z.Request("/phone/auto_receptionists", "GET")
	if err != nil {
		log.Println(err)
		return ListAutoReceptionists{}, err
	}

	var listautoreceptionists ListAutoReceptionists

	// Unmarshal the response into ListAutoReceptionists struct
	err = json.Unmarshal(response, &listautoreceptionists)
	if err != nil {
		log.Println(err)
		return ListAutoReceptionists{}, err
	}
	//return listautoreceptionists struct
	return listautoreceptionists, nil

}
