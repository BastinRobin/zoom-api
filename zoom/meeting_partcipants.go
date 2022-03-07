package zoom

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Meeting_Partcipants struct {
	Device       string    `json:"device"`
	Domain       string    `json:"domain"`
	HarddiskID   string    `json:"harddisk_id"`
	IPAddress    string    `json:"ip_address"`
	JoinTime     time.Time `json:"join_time"`
	LeaveTime    time.Time `json:"leave_time"`
	Location     string    `json:"location"`
	MacAddr      string    `json:"mac_addr"`
	PcName       string    `json:"pc_name"`
	UserID       string    `json:"user_id"`
	UserName     string    `json:"user_name"`
	UserQos      UserQos
	AsOutputs    AsOutputs
	AudioInputs  AudioInputs
	AudioOutputs AudioOutputs
	CPUUsages    CPUUsages
	VideoInputs  VideoInputs
	VideoOutputs VideoOutputs
}
type UserQos []Userqo
type Userqo struct {
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

// Get Specific meeting details
func (z *Zoom) GetMeetingPartcipants(meetingId string, participantId string) (Meeting_Partcipants, error) {
	// Create the url for getting meetings
	endpoint := fmt.Sprintf("/metrics/meetings/%v/participants/%v/qos", meetingId, participantId)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return Meeting_Partcipants{}, err
	}

	// Umarshal the response into Meeting_Participant struct
	var Partcipant Meeting_Partcipants
	err = json.Unmarshal(response, &Partcipant)
	if err != nil {
		log.Println(err)
		return Meeting_Partcipants{}, err
	}
	//return Meeting_Partcipant structure
	return Partcipant, nil
}
