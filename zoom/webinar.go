package zoom

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Webinar struct {
	Host             string    `json:"host"`
	Duration         int       `json:"duration"`
	Email            string    `json:"email"`
	EndTime          time.Time `json:"end_time"`
	Has3RdPartyAudio bool      `json:"has_3rd_party_audio"`
	HasArchiving     bool      `json:"has_archiving"`
	HasPstn          bool      `json:"has_pstn"`
	HasRecording     bool      `json:"has_recording"`
	HasScreenShare   bool      `json:"has_screen_share"`
	HasSip           bool      `json:"has_sip"`
	HasVideo         bool      `json:"has_video"`
	HasVoip          bool      `json:"has_voip"`
	ID               int       `json:"id"`
	Participants     int       `json:"participants"`
	StartTime        time.Time `json:"start_time"`
	Topic            string    `json:"topic"`
	UserType         string    `json:"user_type"`
	UUID             string    `json:"uuid"`
}

type Post_WebinarFeedback struct {
	NextPageToken string `json:"next_page_token"`
	PageSize      int    `json:"page_size"`
	Feedbacks     Feedbacks
}
type Feedbacks []Feedback
type Feedback struct {
	DateTime time.Time `json:"date_time"`
	Email    string    `json:"email"`
	Quality  string    `json:"quality"`
	UserID   string    `json:"user_id"`
}

type List_WebinarParticipants struct {
	NextPageToken       string    `json:"next_page_token"`
	PageCount           string    `json:"page_count"`
	PageSize            string    `json:"page_size"`
	DateTime            time.Time `json:"date_time"`
	WebinarParticipants WebinarParticipants
	Web_UserQoss        Web_UserQoss
}
type WebinarParticipants []WebinarParticipant
type WebinarParticipant struct {
	Device     string    `json:"device"`
	Domain     string    `json:"domain"`
	HarddiskID string    `json:"harddisk_id"`
	IPAddress  string    `json:"ip_address"`
	JoinTime   time.Time `json:"join_time"`
	LeaveTime  time.Time `json:"leave_time"`
	Location   string    `json:"location"`
	MacAddr    string    `json:"mac_addr"`
	PcName     string    `json:"pc_name"`
	UserID     string    `json:"user_id"`
	UserName   string    `json:"user_name"`
}
type Web_UserQoss []Web_UserQos
type Web_UserQos struct {
	Web_AsInputs     Web_AsInputs
	Web_AsOutputs    Web_AsOutputs
	Web_AudioInputs  Web_AudioInputs
	Web_AudioOutputs Web_AudioOutputs
	Web_CPUUsages    Web_CPUUsages
	Web_VideoInputs  Web_VideoInputs
	Web_VideoOutputs Web_VideoOutputs
}
type Web_AsInputs []Web_AsInput
type Web_AsInput struct {
	AvgLoss    string `json:"avg_loss"`
	Bitrate    string `json:"bitrate"`
	FrameRate  string `json:"frame_rate"`
	Jitter     string `json:"jitter"`
	Latency    string `json:"latency"`
	MaxLoss    string `json:"max_loss"`
	Resolution string `json:"resolution"`
}
type Web_AsOutputs []Web_AsOutput
type Web_AsOutput struct {
	AvgLoss    string `json:"avg_loss"`
	Bitrate    string `json:"bitrate"`
	FrameRate  string `json:"frame_rate"`
	Jitter     string `json:"jitter"`
	Latency    string `json:"latency"`
	MaxLoss    string `json:"max_loss"`
	Resolution string `json:"resolution"`
}
type Web_AudioInputs []Web_AudioInput
type Web_AudioInput struct {
	AvgLoss string `json:"avg_loss"`
	Bitrate string `json:"bitrate"`
	Jitter  string `json:"jitter"`
	Latency string `json:"latency"`
	MaxLoss string `json:"max_loss"`
}
type Web_AudioOutputs []Web_AudioOutput
type Web_AudioOutput struct {
	AvgLoss string `json:"avg_loss"`
	Bitrate string `json:"bitrate"`
	Jitter  string `json:"jitter"`
	Latency string `json:"latency"`
	MaxLoss string `json:"max_loss"`
}
type Web_CPUUsages []Web_CPUUsage
type Web_CPUUsage struct {
	SystemMaxCPUUsage string `json:"system_max_cpu_usage"`
	ZoomAvgCPUUsage   string `json:"zoom_avg_cpu_usage"`
	ZoomMaxCPUUsage   string `json:"zoom_max_cpu_usage"`
	ZoomMinCPUUsage   string `json:"zoom_min_cpu_usage"`
}

type Web_VideoInputs []Web_VideoInput
type Web_VideoInput struct {
	AvgLoss    string `json:"avg_loss"`
	Bitrate    string `json:"bitrate"`
	FrameRate  string `json:"frame_rate"`
	Jitter     string `json:"jitter"`
	Latency    string `json:"latency"`
	MaxLoss    string `json:"max_loss"`
	Resolution string `json:"resolution"`
}
type Web_VideoOutputs []Web_VideoOutput
type Web_VideoOutput struct {
	AvgLoss    string `json:"avg_loss"`
	Bitrate    string `json:"bitrate"`
	FrameRate  string `json:"frame_rate"`
	Jitter     string `json:"jitter"`
	Latency    string `json:"latency"`
	MaxLoss    string `json:"max_loss"`
	Resolution string `json:"resolution"`
}

// return the list of all  webinar details
func (z *Zoom) GetAllWebinar(webinar_id string) (Webinar, error) {
	//fmt.Print("Enter Webinar ID")
	// Create the url for getting webinar
	endpoint := fmt.Sprintf("/metrics/webinars/%v", webinar_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return Webinar{}, err
	}

	// Umarshal the response into Webinar struct
	var webinar Webinar
	err = json.Unmarshal(response, &webinar)
	if err != nil {
		log.Println(err)
		return Webinar{}, err
	}
	//return webinar struct
	return webinar, nil
}

// return all post webinar feedback
func (z *Zoom) GetPostWebinarFeedback(webinar_id string) (Post_WebinarFeedback, error) {
	// Create the url for getting webinar
	endpoint := fmt.Sprintf("/metrics/webinars/%v/participants/satisfaction", webinar_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return Post_WebinarFeedback{}, err
	}

	// Umarshal the response into PostWebinarFeedback struct
	var post_webinarfeedback Post_WebinarFeedback
	err = json.Unmarshal(response, &post_webinarfeedback)
	if err != nil {
		log.Println(err)
		return Post_WebinarFeedback{}, err
	}
	//return postwebinarfeedback struct
	return post_webinarfeedback, nil
}

// return a list of webinar participants and their quality of service
func (z *Zoom) GetListWebinarParticipantsQoS(webinar_id string) (List_WebinarParticipants, error) {
	// Create the url for getting webinar
	endpoint := fmt.Sprintf("/metrics/webinars/%v/participants/qos", webinar_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return List_WebinarParticipants{}, err
	}

	// Umarshal the response into listWebinarParticipants struct
	var list_webinarparticipants List_WebinarParticipants
	err = json.Unmarshal(response, &list_webinarparticipants)
	if err != nil {
		log.Println(err)
		return List_WebinarParticipants{}, err
	}
	//return listwebinarparticipants struct
	return list_webinarparticipants, nil
}
