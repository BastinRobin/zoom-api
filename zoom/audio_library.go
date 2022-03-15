package zoom

import (
	"encoding/json"
	"fmt"
	"log"
)

type AudioItem struct {
	AudioID string `json:"audio_id"`
	Name    string `json:"name"`
	PlayURL string `json:"play_url"`
}

//Structure for list the AudioItems

type ListAudioItems struct {
	Audios Audio
}
type Audios []Audio
type Audio struct {
	AudioID string `json:"audio_id"`
	Name    string `json:"name"`
}

// Get Specific meeting details
func (z *Zoom) GetAudioItem(audio_id string) (AudioItem, error) {
	// Create the url for getting meetings
	endpoint := fmt.Sprintf("/phone/audios/%v", audio_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return AudioItem{}, err
	}

	// Umarshal the response into struct
	var audio_item AudioItem
	err = json.Unmarshal(response, &audio_item)
	if err != nil {
		log.Println(err)
		return AudioItem{}, err
	}

	return audio_item, nil
}

// Get Specific meeting details
func (z *Zoom) GetListAudioItem(user_id string) (ListAudioItems, error) {
	// Create the url for getting meetings
	endpoint := fmt.Sprintf("/phone/users/%v/audios", user_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return ListAudioItems{}, err
	}

	// Umarshal the response into struct
	var audio_item ListAudioItems
	err = json.Unmarshal(response, &audio_item)
	if err != nil {
		log.Println(err)
		return ListAudioItems{}, err
	}

	return audio_item, nil
}
