package zoom

import (
	"encoding/json"
	"log"
)

type Responses struct {
	From          string `json:"from"`
	NextPageToken string `json:"next_page_token"`
	PageSize      int    `json:"page_size"`
	To            string `json:"to"`
	Chats         Chats
}
type Chats []Chat
type Chat struct {
	AudioSent int `json:"audio_sent"`
	//CodeSippetSent int    `json:"code_sippet_sent"`
	Email      string `json:"email"`
	GiphysSent int    `json:"giphys_sent"`
	GroupSent  int    `json:"group_sent"`
	ImagesSent int    `json:"images_sent"`
	P2PSent    int    `json:"p2p_sent"`
	TextSent   int    `json:"text_sent"`
	TotalSent  int    `json:"total_sent"`
	UserID     string `json:"user_id"`
	UserName   string `json:"user_name"`
	VideoSent  int    `json:"video_sent"`
}

// Return the list of all chats in zoom account
func (z *Zoom) GetChat() (Responses, error) {
	chats, err := z.Request("/metrics/chat", "GET")
	if err != nil {
		log.Println(err)
		return Responses{}, err
	}

	var response Responses

	// Unmarshal the response into Response struct
	err = json.Unmarshal(chats, &response)
	if err != nil {
		log.Println(err)
		return Responses{}, err
	}
	return response, nil
}
