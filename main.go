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

	// // Start calling the methods
	// //Method for GetUser Details
	// users, err := zoom.GetUsers()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(users)

	// //Method for GetMeeting
	// meetings, err := zoom.GetMeeting("8722824397")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(meetings.UUID)

	// //Method for GetPostMeetingFeedback
	// postmeeting, err := zoom.GetPostMeetingFeedback("8722824397")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(postmeeting)

	// //Method for Get25IssuesOfZoomRooms
	// roomissue, err := zoom.Get25IssuesOfZoomRooms()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(roomissue)

	// //Method for GetRoomsWithIssues
	// room, err := zoom.GetRoomsWithIssues("8722824397")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(room)

	// //Method for GetMeetingPartcipants
	// Partcipant, err := zoom.GetMeetingPartcipants("8722824397", "8722824397")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(Partcipant)

	// //Method for GetWebinarDetails
	// webinar, err := zoom.GetWebinarDetails()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(webinar)

	// //Method for GetWebinarPartcipants
	// webinar_partcipant, err := zoom.GetWebinarPartcipants("8722824397")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(webinar_partcipant)

	//Method for GetSIPDevices
	// SIP_devices, err := zoom.GetSIPDevice()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(SIP_devices)

	//Method for GetListGroups
	// list_groups, err := zoom.GetListGroups()
	// if err != nil {
	// 	log.Println(err)
	// }

	// fmt.Println(list_groups)

	// //Method for GetLockedSetting
	// Group_Locked_Setting, err := zoom.GetLockedSetting("chfhfhhfh_TKikJIX0")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(Group_Locked_Setting)

	//Method for GetGroupMembers
	// group_members, err := zoom.GetListGroupMembers("3542342")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(group_members)

	// //Method for GetGroupSettings
	// group_setting, err := zoom.GetGroupSetting("chfhfhhfh_TKikJIX0")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(group_setting)
	// im_groups, err := zoom.GetIMDirectoryGroup("3542342")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(im_groups)

	// meeting_polls, err := zoom.GetMeetingPolls("8722824397")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(meeting_polls)

	// meeting_registrants, err := zoom.GetMeetingRegistrants("8722824397")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(meeting_registrants)
	//Method for GetAudioItem
	// AudioItem, err := zoom.GetAudioItem("DP-6B4CeTjyo-FYfsFXdXQ")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(AudioItem)
	// //Method for ListAudioItem
	// list_audio_item, err := zoom.GetListAudioItem("K3xnsN2kSpSODzIRmDj0Rg")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(list_audio_item)

	//Method for ListBlockedList
	// blocked_list, err := zoom.ListBlockedList()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(blocked_list)

	// //Method for List AutoReceptionists
	// auto_receptionists, err := zoom.ListAutoReceptionists()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(auto_receptionists)

	//Method for ListCallingPlans
	// calling_plans, err := zoom.ListCallingPlans()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(calling_plans)

	// //Method for ListPortedNumbers
	// ported_numbers, err := zoom.ListPortedNumbers()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(ported_numbers)

	//Methof for ListPhoneNumber
	phone_numbers, err := zoom.ListPhoneNumber()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(phone_numbers)

	//Methof for GetPhoneNumberDetails
	phonenumbers_details, err := zoom.GetPhoneNumberDetails("Hfdgdfgdfg1ew")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(phonenumbers_details)
}
