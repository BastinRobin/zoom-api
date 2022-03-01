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
			"Authorization": []string{"Bearer eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOm51bGwsImlzcyI6IjgtWGZkQW5RVDZlR1FzN2JyMHFoanciLCJleHAiOjE2NDYwNjEwMDEsImlhdCI6MTY0NjA1NTYwMX0.IXV9Ub2jeU0v4IkAE1kXMDyHSAf69_2rkhzw3ckFSVk"},
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

	//fmt.Println(meetings)
	chats, err := zoom.GetChat("8722824397")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(chats)
	//for _, chats := range chats.Chats {
	//fmt.Println(chats.UserID, chats.UserName, chats.Email)
	//}
	meeting, err := zoom.GetClientFeedback("8722824397")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(meeting)
}