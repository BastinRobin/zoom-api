package zoom

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGroupSetting(t *testing.T) {

	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"audio_conferencing": {
			  "toll_free_and_fee_based_toll_call": {
				"allow_webinar_attendees_dial": true,
				"enable": true,
				"numbers": [
				  {
					"code": "64",
					"country_code": "US",
					"country_name": "United States",
					"display_number": "+1 5550100(Atlanta)",
					"number": "+1 5550100"
				  },
				  {
					"code": "64",
					"country_code": "US",
					"country_name": "United States",
					"display_number": "+1 5550101(SanJose)",
					"number": "+1 5550101"
				  }
				]
			  }
			},
			"email_notification": {
			  "alternative_host_reminder": true,
			  "cancel_meeting_reminder": true,
			  "cloud_recording_available_reminder": true,
			  "jbh_reminder": true,
			  "recording_available_reminder_alternative_hosts": true,
			  "recording_available_reminder_schedulers": true,
			  "schedule_for_reminder": true
			},
			"in_meeting": {
			  "alert_guest_join": true,
			  "allow_live_streaming": true,
			  "allow_participants_chat_with": 2,
			  "allow_show_zoom_windows": true,
			  "allow_users_save_chats": 2,
			  "annotation": true,
			  "attendee_on_hold": true,
			  "auto_answer": true,
			  "auto_saving_chat": true,
			  "breakout_room": true,
			  "breakout_room_schedule": true,
			  "chat": true,
			  "closed_caption": true,
			  "closed_captioning": {
				"auto_transcribing": true,
				"enable": true,
				"save_caption": true,
				"third_party_captioning_service": false,
				"view_full_transcript": true
			  },
			  "co_host": true,
			  "custom_live_streaming_service": true,
			  "custom_service_instructions": "specific instructions",
			  "disable_screen_sharing_for_host_meetings": true,
			  "disable_screen_sharing_for_in_meeting_guests": true,
			  "e2e_encryption": true,
			  "entry_exit_chime": "all",
			  "far_end_camera_control": true,
			  "feedback": true,
			  "group_hd": true,
			  "join_from_desktop": true,
			  "join_from_mobile": true,
			  "live_streaming_facebook": true,
			  "live_streaming_youtube": true,
			  "manual_captioning": {
				"allow_to_type": false,
				"auto_generated_captions": false,
				"full_transcript": true,
				"manual_captions": true,
				"save_captions": true,
				"third_party_captioning_service": true
			  },
			  "meeting_survey": true,
			  "non_verbal_feedback": true,
			  "only_host_view_device_list": false,
			  "original_audio": true,
			  "polling": true,
			  "post_meeting_feedback": true,
			  "private_chat": true,
			  "record_play_own_voice": false,
			  "remote_control": true,
			  "remote_support": true,
			  "screen_sharing": true,
			  "sending_default_email_invites": true,
			  "show_a_join_from_your_browser_link": true,
			  "show_browser_join_link": true,
			  "show_device_list": false,
			  "show_meeting_control_toolbar": true,
			  "slide_control": true,
			  "stereo_audio": true,
			  "use_html_format_email": true,
			  "virtual_background": true,
			  "virtual_background_settings": {
				"enable": true,
				"allow_videos": true,
				"allow_upload_custom": true,
				"files": [
				  {
					"name": "San Francisco",
					"type": "image",
					"is_default": true
				  },
				  {
					"name": "Grass",
					"type": "image",
					"is_default": true
				  },
				  {
					"name": "Earth",
					"type": "image",
					"is_default": true
				  },
				  {
					"name": "Beach",
					"type": "video",
					"is_default": true
				  },
				  {
					"name": "Northern Lights",
					"type": "video",
					"is_default": true
				  },
				  {
					"id": "j6BKD1tZQaGQtmqrPciEww",
					"name": "snow.png",
					"type": "image",
					"is_default": false,
					"size": 122055
				  }
				]
			  },
			  "waiting_room": true,
			  "webinar_chat": {
				"allow_attendees_chat_with": 1,
				"allow_auto_save_local_chat_file": true,
				"allow_panelists_chat_with": 1,
				"allow_panelists_send_direct_message": true,
				"allow_users_save_chats": 1,
				"default_attendees_chat_with": 1,
				"enable": true
			  },
			  "webinar_live_streaming": {
				"custom_service_instructions": "specific instructions",
				"enable": true,
				"live_streaming_reminder": true,
				"live_streaming_service": [
				  "facebook",
				  "workplace_by_facebook",
				  "youtube",
				  "custom_live_streaming_service"
				]
			  },
			  "webinar_polling": {
				"advanced_polls": true,
				"enable": true
			  },
			  "webinar_question_answer": true,
			  "webinar_survey": true,
			  "whiteboard": true,
			  "workplace_by_facebook": true
			},
			"other_options": {
			  "allow_users_contact_support_via_chat": true,
			  "blur_snapshot": true
			},
			"recording": {
			  "account_user_access_recording": false,
			  "auto_recording": "none",
			  "cloud_recording": true,
			  "cloud_recording_download": true,
			  "cloud_recording_download_host": true,
			  "host_delete_cloud_recording": true,
			  "local_recording": true,
			  "record_audio_file": true,
			  "record_gallery_view": false,
			  "record_speaker_view": true,
			  "recording_audio_transcript": false,
			  "save_chat_text": true,
			  "show_timestamp": false
			},
			"schedule_meeting": {
			  "audio_type": "both",
			  "force_pmi_jbh_password": true,
			  "host_video": true,
			  "join_before_host": true,
			  "mute_upon_entry": true,
			  "participant_video": true,
			  "pstn_password_protected": true,
			  "require_password_for_instant_meetings": false,
			  "require_password_for_pmi_meetings": "all",
			  "require_password_for_scheduled_meetings": true,
			  "require_password_for_scheduling_new_meetings": true,
			  "upcoming_meeting_reminder": true
			},
			"telephony": {
			  "audio_conference_info": "1234656",
			  "third_party_audio": true
			}
		  }`))
	}))

	client := &Zoom{
		BaseUrl: testServer.URL,
	}

	group_setting, err := client.GetGroupSetting("chfhfhhfh_TKikJIX0")
	if err != nil {
		t.Error(err)
	}

	assert.NotNil(t, group_setting, "Empty meetings returned")
	assert.Equal(t, group_setting.CoHost, false, "Co Host is not matching")
	assert.Equal(t, group_setting.OriginalAudio, false, "Original Audio is not matching")
	assert.Equal(t, group_setting.MeetingSurvey, false, "Meeting survey not matching")
	assert.Equal(t, group_setting.CustomLiveStreamingService, false, "Live streaming service is not matching")
	assert.Equal(t, group_setting.E2EEncryption, false, "Encryption not matching")
}
