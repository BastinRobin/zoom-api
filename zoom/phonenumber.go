package zoom

import (
	"encoding/json"
	"fmt"
	"log"
)

//Structure for ListPhoneNumber

type PhoneNumbers1 struct {
	NextPageToken string `json:"next_page_token"`
	PageSize      int    `json:"page_size"`
	Source        string `json:"source"`
	Status        string `json:"status"`
	TotalRecords  int    `json:"total_records"`
}
type Numbers1 []Number1
type Number1 struct {
	Capability  []string `json:"capability,omitempty"`
	DisplayName string   `json:"display_name,omitempty"`
	ID          string   `json:"id"`
	Location    string   `json:"location"`
	Number      string   `json:"number"`
	NumberType  string   `json:"number_type"`
	Assignees   Assignees
}
type Assignees []Assignee
type Assignee struct {
	ExtensionNumber int    `json:"extension_number"`
	ID              string `json:"id"`
	Name            string `json:"name"`
	Type            string `json:"type"`
}
type SitesIDs []SiteID
type SiteID struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

//Structure for get phone number details

type PhoneNumbersDetails struct {
	Capability  []string `json:"capability"`
	DisplayName string   `json:"display_name"`
	ID          string   `json:"id"`
	Location    string   `json:"location"`
	Number      string   `json:"number"`
	NumberType  string   `json:"number_type"`
	Source      string   `json:"source"`
	Status      string   `json:"status"`
}
type Assignees1 []Assignee1
type Assignee1 struct {
	ExtensionNumber int    `json:"extension_number"`
	Name            string `json:"name"`
	Type            string `json:"type"`
}
type Carriers []Carrier
type Carrier struct {
	Code int    `json:"code"`
	Name string `json:"name"`
}

type SipGroups []SipGroup
type SipGroup struct {
	DisplayName string `json:"display_name"`
	ID          string `json:"id"`
}
type Sites2 []Site2
type Site2 struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Return the list of all phonenumbers in zoom account
func (z *Zoom) ListPhoneNumber() (PhoneNumbers1, error) {
	response, err := z.Request("/phone/numbers", "GET")
	if err != nil {
		log.Println(err)
		return PhoneNumbers1{}, err
	}

	var phone_numbers PhoneNumbers1

	// Unmarshal the response into PhoneNumber struct
	err = json.Unmarshal(response, &phone_numbers)
	if err != nil {
		log.Println(err)
		return PhoneNumbers1{}, err
	}
	//return phonenumber structure
	return PhoneNumbers1{}, nil

}

// Get Specific phone number details
func (z *Zoom) GetPhoneNumberDetails(number_id string) (PhoneNumbersDetails, error) {
	// Create the url for getting phone number details
	endpoint := fmt.Sprintf("/phone/numbers/%v", number_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return PhoneNumbersDetails{}, err
	}

	// Umarshal the response into PhoneNumberDetails
	var phonenumber_details PhoneNumbersDetails
	err = json.Unmarshal(response, &phonenumber_details)
	if err != nil {
		log.Println(err)
		return PhoneNumbersDetails{}, err
	}
	//returns phone_number structure
	return phonenumber_details, nil
}
