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
			"Authorization": []string{"Bearer "},
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

	meetings, err := zoom.GetMeeting("8722824397")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(meetings.UUID)

}
