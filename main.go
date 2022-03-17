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

			"Content-Type": []string{"application/json"},
		},
	}

	// chat, err := zoom.GetChat()
	// if err != nil {
	// 	log.Println(err)
	// }

	//fmt.Println(chat)

	// webinar, err := zoom.GetAllWebinar("webinar_id")
	// if err != nil {
	// 	log.Println(err)
	// }

	// fmt.Println(webinar)

	// quality, err := zoom.GetMeetingQualityScore()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(quality)

	// recording, err := zoom.GetRecordingDetails("8722824397")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(recording)

	// client, err := zoom.GetClientMeetingSatisfaction()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(client)

	// meeting_participant, err := zoom.GetListMeetingParticipantsQoS("8722824397")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(meeting_participant)
	// post_webinarfeedback, err := zoom.GetPostWebinarFeedback("8722824397")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(post_webinarfeedback)

	// list_webinarparticipants, err := zoom.GetListWebinarParticipantsQoS("8722824397")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(list_webinarparticipants)

	// list_zoomrooms, err := zoom.GetListZoomRooms()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(list_zoomrooms)

	// group, err := zoom.GetGroup("chfhfhhfh_TKikJIX0")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(group)

	// group_admins, err := zoom.GetGroupAdmins("3542342")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(group_admins)

	// list_imdirectory, err := zoom.GetListIMDirectory()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(list_imdirectory)

	// group_members, err := zoom.GetGroupMembers("dlfjdhq3430394")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(group_members)

	// meeting_invitation, err := zoom.GetMeetingInvitation("8722824397")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(meeting_invitation)

	// meeting_livestream, err := zoom.GetLiveStream("8722824397")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(meeting_livestream)

	// listautoreceptionists, err := zoom.GetListAutoReceptionists()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(listautoreceptionists)

	// blockedlist, err := zoom.GetListBlockedLists()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(blockedlist)

	// blockedlistdetails, err := zoom.GetBlockedListDetails("7bR_Ix4KSJS_heuPPKc0Zw")
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(blockedlistdetails)

	// listPhonenumbers, err := zoom.GetListPhoneNumbers()
	// if err != nil {
	// 	log.Println(err)
	// }
	// fmt.Println(listPhonenumbers)

	listmonitoringroups, err := zoom.GetListMonitoringGroups()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(listmonitoringroups)

	monitoringroupsid, err := zoom.GetMonitoringGroupId("8f71O6rWT8KFUGQmJIFAdQ")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(monitoringroupsid)
}
