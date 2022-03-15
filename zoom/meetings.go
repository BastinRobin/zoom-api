package zoom

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Meeting struct {
	UUID  string `json:"uuid"`
	Id    int    `json:"id"`
	Topic string `json:"topic"`
	Type  int    `json:"type"`
}
type Meeting_Participants struct {
	NextPageToken     string    `json:"next_page_token"`
	PageCount         string    `json:"page_count"`
	PageSize          string    `json:"page_size"`
	DateTime          time.Time `json:"date_time"`
	TotalRecords      string    `json:"total_records"`
	List_Participants List_Participants
	UserQoss          UserQoss
}
type List_Participants []List_Participant
type List_Participant struct {
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
type UserQoss []UserQos
type UserQos struct {
	AsInputs     AsInputs
	AsOutputs    AsOutputs
	AudioInputs  AudioInputs
	AudioOutputs AudioOutputs
	CPUUsages    CPUUsages
	VideoInputs  VideoInputs
	VideoOutputs VideoOutputs
}
type AsInputs []AsInput
type AsInput struct {
	AvgLoss    string `json:"avg_loss"`
	Bitrate    string `json:"bitrate"`
	FrameRate  string `json:"frame_rate"`
	Jitter     string `json:"jitter"`
	Latency    string `json:"latency"`
	MaxLoss    string `json:"max_loss"`
	Resolution string `json:"resolution"`
}
type AsOutputs []AsOutput
type AsOutput struct {
	AvgLoss    string `json:"avg_loss"`
	Bitrate    string `json:"bitrate"`
	FrameRate  string `json:"frame_rate"`
	Jitter     string `json:"jitter"`
	Latency    string `json:"latency"`
	MaxLoss    string `json:"max_loss"`
	Resolution string `json:"resolution"`
}

type AudioInputs []AudioInput
type AudioInput struct {
	AvgLoss string `json:"avg_loss"`
	Bitrate string `json:"bitrate"`
	Jitter  string `json:"jitter"`
	Latency string `json:"latency"`
	MaxLoss string `json:"max_loss"`
}
type AudioOutputs []AudioOutput
type AudioOutput struct {
	AvgLoss string `json:"avg_loss"`
	Bitrate string `json:"bitrate"`
	Jitter  string `json:"jitter"`
	Latency string `json:"latency"`
	MaxLoss string `json:"max_loss"`
}
type CPUUsages []CPUUsage
type CPUUsage struct {
	SystemMaxCPUUsage string `json:"system_max_cpu_usage"`
	ZoomAvgCPUUsage   string `json:"zoom_avg_cpu_usage"`
	ZoomMaxCPUUsage   string `json:"zoom_max_cpu_usage"`
	ZoomMinCPUUsage   string `json:"zoom_min_cpu_usage"`
}

type VideoInputs []VideoInput
type VideoInput struct {
	AvgLoss    string `json:"avg_loss"`
	Bitrate    string `json:"bitrate"`
	FrameRate  string `json:"frame_rate"`
	Jitter     string `json:"jitter"`
	Latency    string `json:"latency"`
	MaxLoss    string `json:"max_loss"`
	Resolution string `json:"resolution"`
}
type VideoOutputs []VideoOutput
type VideoOutput struct {
	AvgLoss    string `json:"avg_loss"`
	Bitrate    string `json:"bitrate"`
	FrameRate  string `json:"frame_rate"`
	Jitter     string `json:"jitter"`
	Latency    string `json:"latency"`
	MaxLoss    string `json:"max_loss"`
	Resolution string `json:"resolution"`
}

type MeetingInvitation struct {
	Invitation string `json:"invitation"`
}

type MeetingLiveStream struct {
	PageURL   string `json:"page_url"`
	StreamKey string `json:"stream_key"`
	StreamURL string `json:"stream_url"`
}

// return Specific meeting details
func (z *Zoom) GetMeeting(meeting_id string) (Meeting, error) {
	// Create the url for getting meetings
	endpoint := fmt.Sprintf("/meetings/%v", meeting_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return Meeting{}, err
	}

	// Umarshal the response into struct
	var meeting Meeting
	err = json.Unmarshal(response, &meeting)
	if err != nil {
		log.Println(err)
		return Meeting{}, err
	}
	//return meeting struct
	return meeting, nil
}

// return a list of meeting participants and their quality of service
func (z *Zoom) GetListMeetingParticipantsQoS(meeting_id string) (Meeting_Participants, error) {
	// Create the url for getting meetings
	endpoint := fmt.Sprintf("/metrics/meetings/%v/participants/qos", meeting_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return Meeting_Participants{}, err
	}

	// Umarshal the response into Meeting_articipant struct
	var meeting_participant Meeting_Participants
	err = json.Unmarshal(response, &meeting_participant)
	if err != nil {
		log.Println(err)
		return Meeting_Participants{}, err
	}
	//return meeting_participant struct
	return meeting_participant, nil
}

// return a meeting invitation
func (z *Zoom) GetMeetingInvitation(meeting_id string) (MeetingInvitation, error) {
	// Create the url for getting meeting invitation
	endpoint := fmt.Sprintf("/meetings/%v/invitation", meeting_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return MeetingInvitation{}, err
	}

	// Umarshal the response into MeetingInvitation struct
	var meetinginvitation MeetingInvitation
	err = json.Unmarshal(response, &meetinginvitation)
	if err != nil {
		log.Println(err)
		return MeetingInvitation{}, err
	}
	//return meetinginvitation struct
	return meetinginvitation, nil
}

// return a meeting live stream details
func (z *Zoom) GetLiveStream(meeting_id string) (MeetingLiveStream, error) {
	// Create the url for getting meeting live stream details
	endpoint := fmt.Sprintf("/meetings/%v/livestream", meeting_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return MeetingLiveStream{}, err
	}

	// Umarshal the response into MeetingLiveStream struct
	var meetinglivestream MeetingLiveStream
	err = json.Unmarshal(response, &meetinglivestream)
	if err != nil {
		log.Println(err)
		return MeetingLiveStream{}, err
	}
	//return meetinglivestream struct
	return meetinglivestream, nil
}
