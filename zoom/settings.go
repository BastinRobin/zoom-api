package zoom

import (
	"encoding/json"
	"log"
	"time"
)

type Settings struct {
	CallingPlans CallingPlans
}
type CallingPlans []CallingPlan
type CallingPlan struct {
	Assigned   int    `json:"assigned"`
	Available  int    `json:"available"`
	Name       string `json:"name"`
	Subscribed int    `json:"subscribed"`
	Type       int    `json:"type"`
}

type PortedNumberSetting struct {
	NextPageToken     string    `json:"next_page_token"`
	PageSize          int       `json:"page_size"`
	Status            string    `json:"status"`
	SubmittedDateTime time.Time `json:"submitted_date_time"`
	TotalRecords      int       `json:"total_records"`
	PortedNumbers     PortedNumbers
}
type PortedNumbers []PortedNumber
type PortedNumber struct {
	Numbers          []string `json:"numbers"`
	OrderID          string   `json:"order_id"`
	ReplacingNumbers ReplacingNumbers
}
type ReplacingNumbers []ReplacingNumber
type ReplacingNumber struct {
	SourceNumber string `json:"source_number"`
	TargetNumber string `json:"target_number"`
}

// Return the list of all users in zoom account
func (z *Zoom) ListCallingPlans() (Settings, error) {
	calling_plans, err := z.Request("/phone/calling_plans", "GET")
	if err != nil {
		log.Println(err)
		return Settings{}, err
	}

	var settings Settings

	// Unmarshal the response into Response struct
	err = json.Unmarshal(calling_plans, &settings)
	if err != nil {
		log.Println(err)
		return Settings{}, err
	}

	return settings, nil
}

// Return the list of all users in zoom account
func (z *Zoom) ListPortedNumbers() (PortedNumberSetting, error) {
	settings, err := z.Request("/phone/ported_numbers/orders", "GET")
	if err != nil {
		log.Println(err)
		return PortedNumberSetting{}, err
	}

	var ported_numbers PortedNumberSetting

	// Unmarshal the response into Response struct
	err = json.Unmarshal(settings, &ported_numbers)
	if err != nil {
		log.Println(err)
		return PortedNumberSetting{}, err
	}

	return ported_numbers, nil

}
