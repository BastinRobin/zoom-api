package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bastinrobin/zoom-api/zoom"
)

func main() {

	// Create a new zoom client to authenticate the api
	zoom := zoom.Zoom{
		BaseUrl: "https://api.zoom.us/v2",
		Headers: http.Header{
			"Authorization": []string{"Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6IjgtWGZkQW5RVDZlR1FzN2JyMHFoanciLCJleHAiOjE2NDY3MjIwNzIsImlhdCI6MTY0NjExNzI3MH0.KlOE4NfT_U-TI64EoLA_5OJn6aM-3rSH-MS-gNHOBZE"},
			"Content-Type":  []string{"application/json"},
		},
	}

	// Start calling the methods
	// users, err := zoom.GetUsers()
	// if err != nil {
	// 	log.Println(err)
	// }

	// for _, user := range users.Users {
	// 	fmt.Println(user.FirstName, user.LastName, user.PMI)
	// }

	//meetings, err := zoom.GetMeeting("8722824397")
	//if err != nil {
	//	log.Println(err)
	//}

	//fmt.Println(meetings.UUID)
	meeting, err := zoom.GetClientFeedback("8722824397")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(meeting)
	//chat, err := zoom.GetChat()
	//if err != nil {
	//	log.Println(err)
	//}

	//fmt.Println(chat)
	postmeetings, err := zoom.GetMeeting("8722824397")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(postmeetings)
}
