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
	//method for GetUser
	users, err := zoom.GetUsers()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(users)
	//method for GetmMeting
	meetings, err := zoom.GetMeeting("8722824397")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(meetings.UUID)
	//method for GetChat
	chat, err := zoom.GetChat()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(chat)
	//method for GetRoomsWithIssues
	rooms, err := zoom.GetRoomsWithIssues("8722824397")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(rooms)
	//method for GetMeetingPartcipants
	Partcipant, err := zoom.GetMeetingPartcipants("8722824397", "8722824397")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(Partcipant)
}
