package zoom

import (
	"encoding/json"
	"log"
)

type AutoReceptionists struct {
	NextPageToken string `json:"next_page_token"`
	PageSize      int    `json:"page_size"`
	TotalRecords  int    `json:"total_records"`
	Sites         Sites
	Receptionists Receptionists
}
type Receptionists []Receptionist
type Receptionist struct {
	CostCenter      string `json:"cost_center"`
	Department      string `json:"department"`
	ExtensionID     string `json:"extension_id"`
	ExtensionNumber int    `json:"extension_number"`
	ID              string `json:"id"`
	Name            string `json:"name"`
	PhoneNumbers    PhoneNumbers
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

// Return the list AutoRecptionists in zoom account
func (z *Zoom) ListAutoReceptionists() (AutoReceptionists, error) {
	receptionist, err := z.Request("/phone/auto_receptionists", "GET")
	if err != nil {
		log.Println(err)
		return AutoReceptionists{}, err
	}

	var auto_receptionists AutoReceptionists

	// Unmarshal the response into AutoRecopnists struct
	err = json.Unmarshal(receptionist, &auto_receptionists)
	if err != nil {
		log.Println(err)
		return AutoReceptionists{}, err
	}
	//return auto_receptionists Structure
	return auto_receptionists, nil

}
