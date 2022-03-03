package zoom

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Partcipants struct {
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
	UserQos    []struct {
		AsInput struct {
			AvgLoss    string `json:"avg_loss"`
			Bitrate    string `json:"bitrate"`
			FrameRate  string `json:"frame_rate"`
			Jitter     string `json:"jitter"`
			Latency    string `json:"latency"`
			MaxLoss    string `json:"max_loss"`
			Resolution string `json:"resolution"`
		} `json:"as_input"`
		AsOutput struct {
			AvgLoss    string `json:"avg_loss"`
			Bitrate    string `json:"bitrate"`
			FrameRate  string `json:"frame_rate"`
			Jitter     string `json:"jitter"`
			Latency    string `json:"latency"`
			MaxLoss    string `json:"max_loss"`
			Resolution string `json:"resolution"`
		} `json:"as_output"`
		AudioInput struct {
			AvgLoss string `json:"avg_loss"`
			Bitrate string `json:"bitrate"`
			Jitter  string `json:"jitter"`
			Latency string `json:"latency"`
			MaxLoss string `json:"max_loss"`
		} `json:"audio_input"`
		AudioOutput struct {
			AvgLoss string `json:"avg_loss"`
			Bitrate string `json:"bitrate"`
			Jitter  string `json:"jitter"`
			Latency string `json:"latency"`
			MaxLoss string `json:"max_loss"`
		} `json:"audio_output"`
		CPUUsage struct {
			SystemMaxCPUUsage string `json:"system_max_cpu_usage"`
			ZoomAvgCPUUsage   string `json:"zoom_avg_cpu_usage"`
			ZoomMaxCPUUsage   string `json:"zoom_max_cpu_usage"`
			ZoomMinCPUUsage   string `json:"zoom_min_cpu_usage"`
		} `json:"cpu_usage"`
		DateTime   time.Time `json:"date_time"`
		VideoInput struct {
			AvgLoss    string `json:"avg_loss"`
			Bitrate    string `json:"bitrate"`
			FrameRate  string `json:"frame_rate"`
			Jitter     string `json:"jitter"`
			Latency    string `json:"latency"`
			MaxLoss    string `json:"max_loss"`
			Resolution string `json:"resolution"`
		} `json:"video_input"`
		VideoOutput struct {
			AvgLoss    string `json:"avg_loss"`
			Bitrate    string `json:"bitrate"`
			FrameRate  string `json:"frame_rate"`
			Jitter     string `json:"jitter"`
			Latency    string `json:"latency"`
			MaxLoss    string `json:"max_loss"`
			Resolution string `json:"resolution"`
		} `json:"video_output"`
	} `json:"user_qos"`
	Version string `json:"version"`
}

// Get Specific meeting details
func (z *Zoom) GetMeetingPartcipants(meetingId string, participantId string) (Partcipants, error) {
	// Create the url for getting meetings
	endpoint := fmt.Sprintf("/metrics/meetings/%v/participants/%v/qos", meetingId, participantId)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return Partcipants{}, err
	}

	// Umarshal the response into struct
	var Partcipant Partcipants
	err = json.Unmarshal(response, &Partcipant)
	if err != nil {
		log.Println(err)
		return Partcipants{}, err
	}

	return Partcipant, nil
}
