package zoom

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Webinar_Details struct {
	From          string `json:"from"`
	To            string `json:"to"`
	NextPageToken string `json:"next_page_token"`
	PageCount     int    `json:"page_count"`
	PageSize      int    `json:"page_size"`
	TotalRecords  int    `json:"total_records"`
	Webinars      Webinars
}
type Webinars []Webinar
type Webinar struct {
	Host               string    `json:"host"`
	Dept               string    `json:"dept"`
	Duration           string    `json:"duration"`
	Email              string    `json:"email"`
	EndTime            time.Time `json:"end_time"`
	Has3RdPartyAudio   bool      `json:"has_3rd_party_audio"`
	HasArchiving       bool      `json:"has_archiving"`
	HasPstn            bool      `json:"has_pstn"`
	HasRecording       bool      `json:"has_recording"`
	HasScreenShare     bool      `json:"has_screen_share"`
	HasSip             bool      `json:"has_sip"`
	HasVideo           bool      `json:"has_video"`
	HasVoip            bool      `json:"has_voip"`
	ID                 int       `json:"id"`
	Participants       int       `json:"participants"`
	StartTime          time.Time `json:"start_time"`
	Topic              string    `json:"topic"`
	UserType           string    `json:"user_type"`
	UUID               string    `json:"uuid"`
	AudioQuality       string    `json:"audio_quality"`
	VideoQuality       string    `json:"video_quality"`
	ScreenShareQuality string    `json:"screen_share_quality"`
	CustomerKeys       CustomerKeys
}
type CustomerKeys []CustomerKey
type CustomerKey struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

//Structure for WebinarParcipants

type Webinar_Partcipants struct {
	NextPageToken string `json:"next_page_token"`
	PageCount     int    `json:"page_count"`
	PageSize      int    `json:"page_size"`
	TotalRecords  int    `json:"total_records"`
	Participants  Participants
}
type Participants []Participant
type Participant struct {
	Camera           string    `json:"camera"`
	ConnectionType   string    `json:"connection_type"`
	DataCenter       string    `json:"data_center"`
	Device           string    `json:"device"`
	Domain           string    `json:"domain"`
	FromSipURI       string    `json:"from_sip_uri"`
	FullDataCenter   string    `json:"full_data_center"`
	HarddiskID       string    `json:"harddisk_id"`
	ID               string    `json:"id"`
	IPAddress        string    `json:"ip_address"`
	JoinTime         time.Time `json:"join_time"`
	LeaveReason      string    `json:"leave_reason"`
	LeaveTime        time.Time `json:"leave_time"`
	Location         string    `json:"location"`
	MacAddr          string    `json:"mac_addr"`
	Microphone       string    `json:"microphone"`
	NetworkType      string    `json:"network_type"`
	PcName           string    `json:"pc_name"`
	Recording        bool      `json:"recording"`
	Role             string    `json:"role"`
	ShareApplication bool      `json:"share_application"`
	ShareDesktop     bool      `json:"share_desktop"`
	ShareWhiteboard  bool      `json:"share_whiteboard"`
	SipURI           string    `json:"sip_uri"`
	Speaker          string    `json:"speaker"`
	Status           string    `json:"status"`
	UserID           string    `json:"user_id"`
	UserName         string    `json:"user_name"`
	Version          string    `json:"version"`
}

// Return the list of all users in zoom account
func (z *Zoom) GetWebinarDetails() (Webinar_Details, error) {
	webinar, err := z.Request("/metrics/webinars", "GET")
	if err != nil {
		log.Println(err)
		return Webinar_Details{}, err
	}

	var webinar_details Webinar_Details

	// Unmarshal the response into Response struct
	err = json.Unmarshal(webinar, &webinar_details)
	if err != nil {
		log.Println(err)
		return Webinar_Details{}, err
	}

	return webinar_details, nil

}

// Get WebinarPartcipants details
func (z *Zoom) GetWebinarPartcipants(webinarId string) (Webinar_Partcipants, error) {
	// Create the url for getting webinarid
	endpoint := fmt.Sprintf("/metrics/webinars/%v/participants", webinarId)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return Webinar_Partcipants{}, err
	}

	// Umarshal the response into Webinar_Parcipant struct
	var webinar_partcipants Webinar_Partcipants
	err = json.Unmarshal(response, &webinar_partcipants)
	if err != nil {
		log.Println(err)
		return Webinar_Partcipants{}, err
	}
	//return Webinar_Partcipant structure
	return webinar_partcipants, nil
}
