package zoom

import (
	"encoding/json"
	"fmt"
	"log"
)

//Structure for list the Groups
type GroupList struct {
	TotalRecords int `json:"total_records"`
	Groups       Groups
}
type Groups []Group
type Group struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	TotalMembers int    `json:"total_members"`
}

//Structure for Group Locked Setting
type Group_Locked_Setting struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	TotalMembers int    `json:"total_members"`
}

//Structure for Group Members
type Group_Members struct {
	PageCount    int `json:"page_count"`
	PageNumber   int `json:"page_number"`
	PageSize     int `json:"page_size"`
	TotalRecords int `json:"total_records"`
	Members      Members
}
type Members []Member
type Member struct {
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	ID        string `json:"id"`
	LastName  string `json:"last_name"`
	Type      int    `json:"type"`
}

//Structure for Group_Settings
type Group_Settings struct {
	CoHost                                 bool   `json:"co_host"`
	CustomLiveStreamingService             bool   `json:"custom_live_streaming_service"`
	CustomServiceInstructions              string `json:"custom_service_instructions"`
	DisableScreenSharingForHostMeetings    bool   `json:"disable_screen_sharing_for_host_meetings"`
	DisableScreenSharingForInMeetingGuests bool   `json:"disable_screen_sharing_for_in_meeting_guests"`
	E2EEncryption                          bool   `json:"e2e_encryption"`
	EntryExitChime                         string `json:"entry_exit_chime"`
	FarEndCameraControl                    bool   `json:"far_end_camera_control"`
	Feedback                               bool   `json:"feedback"`
	GroupHd                                bool   `json:"group_hd"`
	JoinFromDesktop                        bool   `json:"join_from_desktop"`
	JoinFromMobile                         bool   `json:"join_from_mobile"`
	LiveStreamingFacebook                  bool   `json:"live_streaming_facebook"`
	LiveStreamingYoutube                   bool   `json:"live_streaming_youtube"`
	MeetingSurvey                          bool   `json:"meeting_survey"`
	NonVerbalFeedback                      bool   `json:"non_verbal_feedback"`
	OnlyHostViewDeviceList                 bool   `json:"only_host_view_device_list"`
	OriginalAudio                          bool   `json:"original_audio"`
	Polling                                bool   `json:"polling"`
	PostMeetingFeedback                    bool   `json:"post_meeting_feedback"`
	PrivateChat                            bool   `json:"private_chat"`
	RecordPlayOwnVoice                     bool   `json:"record_play_own_voice"`
	RemoteControl                          bool   `json:"remote_control"`
	RemoteSupport                          bool   `json:"remote_support"`
	ScreenSharing                          bool   `json:"screen_sharing"`
	SendingDefaultEmailInvites             bool   `json:"sending_default_email_invites"`
	ShowAJoinFromYourBrowserLink           bool   `json:"show_a_join_from_your_browser_link"`
	ShowBrowserJoinLink                    bool   `json:"show_browser_join_link"`
	ShowDeviceList                         bool   `json:"show_device_list"`
	ShowMeetingControlToolbar              bool   `json:"show_meeting_control_toolbar"`
	SlideControl                           bool   `json:"slide_control"`
	StereoAudio                            bool   `json:"stereo_audio"`
	UseHTMLFormatEmail                     bool   `json:"use_html_format_email"`
	VirtualBackground                      bool   `json:"virtual_background"`
	WaitingRoom                            bool   `json:"waiting_room"`
	WebinarQuestionAnswer                  bool   `json:"webinar_question_answer"`
	WebinarSurvey                          bool   `json:"webinar_survey"`
	Whiteboard                             bool   `json:"whiteboard"`
	WorkplaceByFacebook                    bool   `json:"workplace_by_facebook"`
	AudioConferencings                     AudioConferencings
}
type AudioConferencings []AudioConferencing
type AudioConferencing struct {
	TollFreeAndFeeBasedTollCalls TollFreeAndFeeBasedTollCalls
	Numbers                      Numbers
	EmailNotifications           EmailNotifications
	InMeetings                   InMeetings
	ClosedCaptionings            ClosedCaptionings
	ManualCaptionings            ManualCaptionings
	VirtualBackgroundSettings    VirtualBackgroundSettings
	Files                        Files
	WebinarChats                 WebinarChats
	WebinarLiveStreamings        WebinarLiveStreamings
	OtherOptions                 OtherOptions
	Recordings                   Recordings
	ScheduleMeetings             ScheduleMeetings
	Telephonys                   Telephonys
}

type TollFreeAndFeeBasedTollCalls []TollFreeAndFeeBasedTollCall
type TollFreeAndFeeBasedTollCall struct {
	AllowWebinarAttendeesDial bool `json:"allow_webinar_attendees_dial"`
	Enable                    bool `json:"enable"`
}
type Numbers []Number
type Number struct {
	Code          string `json:"code"`
	CountryCode   string `json:"country_code"`
	CountryName   string `json:"country_name"`
	DisplayNumber string `json:"display_number"`
	Number        string `json:"number"`
}
type EmailNotifications []EmailNotification
type EmailNotification struct {
	AlternativeHostReminder                    bool `json:"alternative_host_reminder"`
	CancelMeetingReminder                      bool `json:"cancel_meeting_reminder"`
	CloudRecordingAvailableReminder            bool `json:"cloud_recording_available_reminder"`
	JbhReminder                                bool `json:"jbh_reminder"`
	RecordingAvailableReminderAlternativeHosts bool `json:"recording_available_reminder_alternative_hosts"`
	RecordingAvailableReminderSchedulers       bool `json:"recording_available_reminder_schedulers"`
	ScheduleForReminder                        bool `json:"schedule_for_reminder"`
}
type InMeetings []InMeeting
type InMeeting struct {
	AlertGuestJoin            bool `json:"alert_guest_join"`
	AllowLiveStreaming        bool `json:"allow_live_streaming"`
	AllowParticipantsChatWith int  `json:"allow_participants_chat_with"`
	AllowShowZoomWindows      bool `json:"allow_show_zoom_windows"`
	AllowUsersSaveChats       int  `json:"allow_users_save_chats"`
	Annotation                bool `json:"annotation"`
	AttendeeOnHold            bool `json:"attendee_on_hold"`
	AutoAnswer                bool `json:"auto_answer"`
	AutoSavingChat            bool `json:"auto_saving_chat"`
	BreakoutRoom              bool `json:"breakout_room"`
	BreakoutRoomSchedule      bool `json:"breakout_room_schedule"`
	Chat                      bool `json:"chat"`
	ClosedCaption             bool `json:"closed_caption"`
}
type ClosedCaptionings []ClosedCaptioning
type ClosedCaptioning struct {
	AutoTranscribing            bool `json:"auto_transcribing"`
	Enable                      bool `json:"enable"`
	SaveCaption                 bool `json:"save_caption"`
	ThirdPartyCaptioningService bool `json:"third_party_captioning_service"`
	ViewFullTranscript          bool `json:"view_full_transcript"`
}

type ManualCaptionings []ManualCaptioning
type ManualCaptioning struct {
	AllowToType                 bool `json:"allow_to_type"`
	AutoGeneratedCaptions       bool `json:"auto_generated_captions"`
	FullTranscript              bool `json:"full_transcript"`
	ManualCaptions              bool `json:"manual_captions"`
	SaveCaptions                bool `json:"save_captions"`
	ThirdPartyCaptioningService bool `json:"third_party_captioning_service"`
}

type VirtualBackgroundSettings []VirtualBackgroundSetting
type VirtualBackgroundSetting struct {
	Enable            bool `json:"enable"`
	AllowVideos       bool `json:"allow_videos"`
	AllowUploadCustom bool `json:"allow_upload_custom"`
}
type Files []File
type File struct {
	Name      string `json:"name"`
	Type      string `json:"type"`
	IsDefault bool   `json:"is_default"`
	ID        string `json:"id,omitempty"`
	Size      int    `json:"size,omitempty"`
}
type WebinarChats []WebinarChat
type WebinarChat struct {
	AllowAttendeesChatWith          int  `json:"allow_attendees_chat_with"`
	AllowAutoSaveLocalChatFile      bool `json:"allow_auto_save_local_chat_file"`
	AllowPanelistsChatWith          int  `json:"allow_panelists_chat_with"`
	AllowPanelistsSendDirectMessage bool `json:"allow_panelists_send_direct_message"`
	AllowUsersSaveChats             int  `json:"allow_users_save_chats"`
	DefaultAttendeesChatWith        int  `json:"default_attendees_chat_with"`
	Enable                          bool `json:"enable"`
}
type WebinarLiveStreamings []WebinarLiveStreaming
type WebinarLiveStreaming struct {
	CustomServiceInstructions string   `json:"custom_service_instructions"`
	Enable                    bool     `json:"enable"`
	LiveStreamingReminder     bool     `json:"live_streaming_reminder"`
	LiveStreamingService      []string `json:"live_streaming_service"`
}
type WebinarPollings []WebinarPolling
type WebinarPolling struct {
	AdvancedPolls bool `json:"advanced_polls"`
	Enable        bool `json:"enable"`
}
type OtherOptions []OtherOption
type OtherOption struct {
	AllowUsersContactSupportViaChat bool `json:"allow_users_contact_support_via_chat"`
	BlurSnapshot                    bool `json:"blur_snapshot"`
}
type Recordings []Recording
type Recording struct {
	AccountUserAccessRecording bool   `json:"account_user_access_recording"`
	AutoRecording              string `json:"auto_recording"`
	CloudRecording             bool   `json:"cloud_recording"`
	CloudRecordingDownload     bool   `json:"cloud_recording_download"`
	CloudRecordingDownloadHost bool   `json:"cloud_recording_download_host"`
	HostDeleteCloudRecording   bool   `json:"host_delete_cloud_recording"`
	LocalRecording             bool   `json:"local_recording"`
	RecordAudioFile            bool   `json:"record_audio_file"`
	RecordGalleryView          bool   `json:"record_gallery_view"`
	RecordSpeakerView          bool   `json:"record_speaker_view"`
	RecordingAudioTranscript   bool   `json:"recording_audio_transcript"`
	SaveChatText               bool   `json:"save_chat_text"`
	ShowTimestamp              bool   `json:"show_timestamp"`
}
type ScheduleMeetings []ScheduleMeeting
type ScheduleMeeting struct {
	AudioType                               string `json:"audio_type"`
	ForcePmiJbhPassword                     bool   `json:"force_pmi_jbh_password"`
	HostVideo                               bool   `json:"host_video"`
	JoinBeforeHost                          bool   `json:"join_before_host"`
	MuteUponEntry                           bool   `json:"mute_upon_entry"`
	ParticipantVideo                        bool   `json:"participant_video"`
	PstnPasswordProtected                   bool   `json:"pstn_password_protected"`
	RequirePasswordForInstantMeetings       bool   `json:"require_password_for_instant_meetings"`
	RequirePasswordForPmiMeetings           string `json:"require_password_for_pmi_meetings"`
	RequirePasswordForScheduledMeetings     bool   `json:"require_password_for_scheduled_meetings"`
	RequirePasswordForSchedulingNewMeetings bool   `json:"require_password_for_scheduling_new_meetings"`
	UpcomingMeetingReminder                 bool   `json:"upcoming_meeting_reminder"`
}
type Telephonys []Telephony
type Telephony struct {
	AudioConferenceInfo string `json:"audio_conference_info"`
	ThirdPartyAudio     bool   `json:"third_party_audio"`
}

// Return the list of all groups in zoom account
func (z *Zoom) GetListGroups() (GroupList, error) {
	groups, err := z.Request("/groups", "GET")
	if err != nil {
		log.Println(err)
		return GroupList{}, err
	}

	var groups_list GroupList

	// Unmarshal the groups into GroupsList struct
	err = json.Unmarshal(groups, &groups_list)
	if err != nil {
		log.Println(err)
		return GroupList{}, err
	}
	// returns group_list structure
	return groups_list, nil

}

// Get Group Locked Setting details
func (z *Zoom) GetLockedSetting(group_id string) (Group_Locked_Setting, error) {
	// Create the url for getting meetings
	endpoint := fmt.Sprintf("/groups/%v", group_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return Group_Locked_Setting{}, err
	}

	// Umarshal the response into struct
	var group_locked_setting Group_Locked_Setting
	err = json.Unmarshal(response, &group_locked_setting)
	if err != nil {
		log.Println(err)
		return Group_Locked_Setting{}, err
	}
	//returns group_locked_setting structure
	return group_locked_setting, nil
}

// Get Group Members details
func (z *Zoom) GetListGroupMembers(group_id string) (Group_Members, error) {
	// Create the url for getting groups
	endpoint := fmt.Sprintf("/groups/%v/members", group_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return Group_Members{}, err
	}

	// Umarshal the response into Group_Members struct
	var group_members Group_Members
	err = json.Unmarshal(response, &group_members)
	if err != nil {
		log.Println(err)
		return Group_Members{}, err
	}
	//returns group_members structure
	return group_members, nil
}

// Get Group Setting details
func (z *Zoom) GetGroupSetting(group_id string) (Group_Settings, error) {
	// Create the url for getting group
	endpoint := fmt.Sprintf("/groups/%v/settings", group_id)
	response, err := z.Request(endpoint, "GET")
	if err != nil {
		log.Println(err)
		return Group_Settings{}, err
	}

	// Umarshal the response into Group_Settings struct
	var group_setting Group_Settings
	err = json.Unmarshal(response, &group_setting)
	if err != nil {
		log.Println(err)
		return Group_Settings{}, err
	}
	//returns group_settings structure
	return group_setting, nil
}
