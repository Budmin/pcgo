package services

import (
	"time"

	"github.com/Budmin/pcgo"
)

type Arrangement struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		BPM                  float32       `json:"bpm"`
		CreatedAt            time.Time     `json:"created_at"`
		HasChords            bool          `json:"has_chords"`
		Length               int           `json:"length"`
		Meter                string        `json:"meter"`
		Name                 string        `json:"name"`
		Notes                string        `json:"notes"`
		PrintMargin          string        `json:"print_margin"`
		PrintOrientation     string        `json:"print_orientation"`
		PrintPageSize        string        `json:"print_page_size"`
		UpdatedAt            time.Time     `json:"updated_at"`
		ChordChart           string        `json:"chord_chart"`
		ChordChartFont       string        `json:"chord_chart_font"`
		ChordChartKey        string        `json:"chord_chart_key"`
		ChordChartColumns    int           `json:"chord_chart_columns"`
		ChordChartFontSize   int           `json:"chord_chart_font_size"`
		HasChordChart        bool          `json:"has_chord_chart"`
		LyricsEnabled        bool          `json:"lyrics_enabled"`
		NumberChartEnabled   bool          `json:"number_chart_enabled"`
		NumeralChartEnabled  bool          `json:"numeral_chart_enabled"`
		Sequence             []interface{} `json:"sequence"`
		SequenceShort        []interface{} `json:"sequence_short"`
		SequenceFull         []interface{} `json:"sequence_full"`
		ChordChartChordColor int           `json:"chord_chart_chord_color"`
		ArchivedAt           time.Time     `json:"archived_at"`
		Lyrics               string        `json:"lyrics"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type ArrangementSections struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Sections []interface{} `json:"sections"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Attachment struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt         time.Time `json:"created_at"`
		PageOrder         string    `json:"page_order"`
		UpdatedAt         time.Time `json:"updated_at"`
		Filename          string    `json:"fileman"`
		FileSize          int       `json:"file_size"`
		LicensesPurchased int       `json:"licenses_purchased"`
		LicensesRemaining int       `json:"licenses_remaining"`
		LicensesUsed      int       `json:"licenses_used"`
		ContentType       string    `json:"content_type"`
		DisplayName       string    `json:"display_name"`
		Filetype          string    `json:"filetype"`
		LinkedUrl         string    `json:"linked_url"`
		PCOType           string    `json:"pco_type"`
		RemoteLink        string    `json:"remote_link"`
		ThumbnailUrl      string    `json:"thumbnail_url"`
		Url               string    `json:"url"`
		AllowMp3Download  bool      `json:"allow_mp3_download"`
		WebStreamable     bool      `json:"web_streamable"`
		Downloadable      bool      `json:"downloadable"`
		Transposable      bool      `json:"transposable"`
		Streamable        bool      `json:"streamable"`
		HasPreview        bool      `json:"has_preview"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type AttachmentActivity struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Date          string `json:"date"`
		AttachmentUrl string `json:"attachment_url"`
		ActivityType  string `json:"activity_type"`
	} `json:"attributes"`

	Relationship map[string]pcgo.RelationshipNode `json:"relationship"`
}

type AttachmentType struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name              string `json:"name"`
		Aliases           string `json:"aliases"`
		CapoedChordCharts bool   `json:"capoed_chord_charts"`
		ChordCharts       bool   `json:"chord_charts"`
		Exclusions        string `json:"exclusions"`
		Lyrics            bool   `json:"lyrics"`
		NumberCharts      bool   `json:"number_charts"`
		NumeralCharts     bool   `json:"numeral_charts"`
		BuiltIn           bool   `json:"built_in"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type AvailableSignup struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		OrganizationName  string `json:"organization_name"`
		PlanningCenterUrl string `json:"planning_center_url"`
		ServiceTypeName   string `json:"service_type_name"`
		SignupsAvailable  bool   `json:"signups_available"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type Blockout struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Description      string    `json:"description"`
		GroupIdentifier  string    `json:"group_itentifier"`
		OrganizationName string    `json:"organization_name"`
		Reason           string    `json:"reason"`
		RepeatFrequency  string    `json:"repeat_frequency"`
		RepeatInterval   string    `json:"repeat_interval"`
		RepeatPeriod     string    `json:"repeat_period"`
		Settings         string    `json:"settings"`
		TimeZone         string    `json:"time_zone"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		RepeatUntil      string    `json:"repeat_until"`
		StartsAt         time.Time `json:"starts_at"`
		EndsAt           time.Time `json:"ends_at"`
		Share            bool      `json:"share"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type BlockoutDate struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		GroupItentifier string    `json:"group_identifier"`
		Reason          string    `json:"reason"`
		TimeZone        string    `json:"time_zone"`
		Share           bool      `json:"share"`
		StartsAt        time.Time `json:"starts_at"`
		EndsAt          time.Time `json:"ends_at"`
		StartsAtUtc     time.Time `json:"starts_at_utc"`
		EndsAtUtc       time.Time `json:"ends_at_utc"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type BlockoutException struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Date      string    `json:"date"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type BlockoutScheduleConflict struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Dates                string    `json:"date"`
		OrganizationName     string    `json:"organization_name"`
		PersonAvatar         string    `json:"person_avatar"`
		PersonName           string    `json:"person_name"`
		PositionDisplayTimes string    `json:"position_display_times"`
		ServiceTypeName      string    `json:"service_type_name"`
		ShortDates           string    `json:"short_dates"`
		Status               string    `json:"status"`
		TeamName             string    `json:"team_name"`
		TeamPositionName     string    `json:"team_position_name"`
		SortDate             time.Time `json:"sort_date"`
		CanAcceptPartial     bool      `json:"can_accept_partial"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type Contributor struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt             time.Time `json:"created_at"`
		UpdatedAt             time.Time `json:"updated_at"`
		ContributableAction   string    `json:"contributable_action"`
		ContributableCategory string    `json:"contributable_category"`
		ContributableType     string    `json:"contributable_type"`
		FullName              string    `json:"full_name"`
		PhotoThumbnailURL     string    `json:"photo_thumbnail_url"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type CustomSlide struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Body    string `json:"body"`
		Label   string `json:"label"`
		Order   int    `json:"order"`
		Enabled bool   `json:"enabled"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type EmailTemplate struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Kind      string    `json:"kind"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		HTMLBody  string    `json:"html_body"`
		Subject   string    `json:"subject"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type EmailTemplateRenderedResponse struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Body    string `json:"body"`
		Subject string `json:"subject"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type Folder struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt time.Time `json:"created_at"`
		Name      string    `json:"name"`
		UpdatedAt time.Time `json:"updated_at"`
		Container string    `json:"container"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type FolderPath struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Path []string `json:"path"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Item struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Title                          string    `json:"title"`
		Sequence                       int       `json:"sequence"`
		CreatedAt                      time.Time `json:"created_at"`
		UpdatedAt                      time.Time `json:"updated_at"`
		Length                         int       `json:"length"`
		ItemType                       string    `json:"item_type"`
		HTMLDetails                    string    `json:"html_details"`
		ServicePosition                string    `json:"service_position"`
		Description                    string    `json:"description"`
		KeyName                        string    `json:"key_name"`
		CustomArrangementSequence      []string  `json:"custom_arrangement_sequence"`
		CustomArrangementSequenceShort []string  `json:"custom_arrangement_sequence_short"`
		CustomArrangementSequenceFull  []string  `json:"custom_arrangement_sequence_full"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type ItemNote struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		Content      string    `json:"content"`
		CategoryName string    `json:"category_name"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type ItemNoteCategory struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt      time.Time `json:"created_at"`
		DeletedAt      time.Time `json:"deleted_at"`
		Name           string    `json:"name"`
		Sequence       int       `json:"sequence"`
		UpdatedAt      time.Time `json:"updated_at"`
		FrequentlyUsed bool      `json:"frequently_used"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type ItemTime struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		LiveStartAt  time.Time `json:"live_start_at"`
		LiveEndAt    time.Time `json:"live_end_at"`
		Exclude      bool      `json:"exclude"`
		Length       int       `json:"length"`
		LengthOffset int       `json:"length_offset"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Key struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		Name          string    `json:"name"`
		AlternateKeys string    `json:"alternate_keys"`
		EndingKey     string    `json:"ending_key"`
		StartingKey   string    `json:"starting_key"`
		StartingMinor bool      `json:"starting_minor"`
		EndingMinor   bool      `json:"ending_minor"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Layout struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes interface{} `json:"attributes"`

	Relationships interface{} `json:"relationship"`
}

type Live struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		SeriesTitle         string `json:"series_title"`
		Title               string `json:"title"`
		Dates               string `json:"dates"`
		LiveChannel         string `json:"live_channel"`
		ChatRoomChannel     string `json:"chat_room_channel"`
		CanControl          bool   `json:"can_control"`
		CanTakeControl      bool   `json:"can_take_control"`
		CanChat             bool   `json:"can_chat"`
		CanControlVideoFeed bool   `json:"can_control_video_feed"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type LiveController struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
		FullName          string    `json:"full_name"`
		PhotoThumbnailURL string    `json:"photo_thumbnail_url"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Media struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt            time.Time `json:"created_at"`
		UpdatedAt            time.Time `json:"updated_at"`
		Themes               string    `json:"themes"`
		Title                string    `json:"title"`
		ThumbnailFileName    string    `json:"thumbnail_file_name"`
		ThumbnailContentType string    `json:"thumbnail_content_type"`
		ThumbnailFileSize    int       `json:"thumbnail_file_size"`
		ThumbnailUpdatedAt   time.Time `json:"thumbnail_updated_at"`
		PreviewFileName      string    `json:"preview_file_name"`
		PreviewContentType   string    `json:"preview_content_type"`
		PreviewFileSize      int       `json:"preview_file_size"`
		PreviewUpdatedAt     time.Time `json:"preview_updated_at"`
		Length               int       `json:"length"`
		MediaType            string    `json:"media_type"`
		MediaTypeName        string    `json:"media_type_name"`
		ThumbnailURL         string    `json:"thumbnail_url"`
		CreatorName          string    `json:"creator_name"`
		PreviewURL           string    `json:"preview_url"`
		ImageURL             string    `json:"image_url"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type MediaSchedule struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		PlanDates       string    `json:"plan_dates"`
		PlanShortDates  string    `json:"plan_short_dates"`
		ServiceTypeName string    `json:"service_type_name"`
		PlanSortDate    time.Time `json:"plan_sort_date"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type NeededPosition struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Quantity         int    `json:"quantity"`
		TeamPositionName string `json:"team_position_name"`
		ScheduledTo      string `json:"scheduled_to"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Organization struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CCLI                            string    `json:"ccli"`
		CreatedAt                       time.Time `json:"created_at"`
		DateFormat                      int       `json:"date_format"`
		MusicStandEnabled               bool      `json:"music_stand_enabled"`
		Name                            string    `json:"name"`
		ProjectorEnabled                bool      `json:"projector_enabled"`
		TimeZone                        string    `json:"time_zone"`
		TwentyFourHourTime              bool      `json:"twenty_four_hour_time"`
		UpdatedAt                       time.Time `json:"updated_at"`
		OwnerName                       string    `json:"owner_name"`
		RequiredToSetDownloadPermission string    `json:"required_to_set_download_permission"`
		Secret                          string    `json:"secret"`
		AllowMp3Download                bool      `json:"allow_mp3_download"`
		CalendarStartsOnSunday          bool      `json:"calendar_starts_on_sunday"`
		CCLIConnected                   bool      `json:"ccli_connected"`
		CCLIReportingEnabled            bool      `json:"ccli_reporting_enabled"`
		ExtraFileStorageAllowed         bool      `json:"extra_file_storage_allowed"`
		FileStorageExceeded             bool      `json:"file_storage_exceeded"`
		FileStorageSize                 bool      `json:"file_storage_size"`
		FileStorageSizeUsed             bool      `json:"file_storage_size_used"`
		FileStorageExtraEnabled         bool      `json:"file_storage_extra_enabled"`
		RehearsalMixEnabled             bool      `json:"rehearsal_mix_enabled"`
		LegacyID                        string    `json:"legacy_id"`
		FileStorageExtraCharges         int       `json:"file_storage_extra_charges"`
		PeopleAllowed                   int       `json:"people_allowed"`
		PeopleRemaining                 int       `json:"people_remaining"`
		Beta                            bool      `json:"beta"`
		RehearsalPackConnected          bool      `json:"rehearsal_pack_connected"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Person struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		PhotoURL                  string        `json:"photo_url"`
		PhotoThumbnailURL         string        `json:"photo_thumbnail_url"`
		PreferredApp              string        `json:"preferred_app"`
		AssignedToRehearsalTeam   bool          `json:"assigned_to_rehearsal_team"`
		ArchivedAt                time.Time     `json:"archived_at"`
		CreatedAt                 time.Time     `json:"created_at"`
		FirstName                 string        `json:"first_name"`
		LastName                  string        `json:"last_name"`
		NamePrefix                string        `json:"name_prefix"`
		NameSuffix                string        `json:"name_suffix"`
		UpdatedAt                 string        `json:"updated_at"`
		FacebookID                string        `json:"facebook_id"`
		LegacyID                  string        `json:"legacy_id"`
		FullName                  string        `json:"full_name"`
		MaxPermissions            string        `json:"max_permissions"`
		Permissions               string        `json:"permissions"`
		Status                    string        `json:"status"`
		Anniversary               string        `json:"anniversary"`
		Birthdate                 string        `json:"birthdate"`
		GivenName                 string        `json:"given_name"`
		MiddleName                string        `json:"middle_name"`
		Nickname                  string        `json:"nickname"`
		AccessMediaAttachments    bool          `json:"access_media_attachments"`
		AccessPlanAttachments     bool          `json:"access_plan_attachments"`
		AccessSongAttachments     bool          `json:"access_song_attachments"`
		Archived                  bool          `json:"archived"`
		SiteAdministrator         bool          `json:"site_administrator"`
		LoggedInAt                time.Time     `json:"logged_in_at"`
		Notes                     string        `json:"notes"`
		PassedBackgroundCheck     bool          `json:"passed_background_check"`
		IcalCode                  string        `json:"ical_code"`
		PreferredMaxPlansPerDay   int           `json:"preferred_max_plans_per_day"`
		PreferredMaxPlansPerMonth int           `json:"preferred_max_plans_per_month"`
		PraiseChartsEnabled       bool          `json:"praise_charts_enabled"`
		MeTab                     string        `json:"me_tab"`
		PlansTab                  string        `json:"plans_tab"`
		SongsTab                  string        `json:"songs_tab"`
		MediaTab                  string        `json:"media_tab"`
		PeopleTab                 string        `json:"people_tab"`
		CanEditAllPeople          bool          `json:"can_edit_all_people"`
		CanViewAllPeople          bool          `json:"can_view_all_people"`
		Onboardings               []interface{} `json:"onboardings"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type PersonTeamPositionAssignment struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt          string        `json:"created_at"`
		UpdatedAt          string        `json:"updated_at"`
		SchedulePreference string        `json:"schedule_preference"`
		PreferredWeeks     []interface{} `json:"preferred_weeks"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Plan struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt            time.Time `json:"created_at"`
		Title                string    `json:"title"`
		UpdatedAt            time.Time `json:"updated_at"`
		Public               bool      `json:"public"`
		SeriesTitle          string    `json:"series_title"`
		PlanNotesCount       int       `json:"plan_notes_count"`
		OtherTimeCount       int       `json:"other_time_count"`
		RehearsalTimeCount   int       `json:"rehearsal_time_count"`
		ServiceTimeCount     int       `json:"service_time_count"`
		PlanPeopleCount      int       `json:"plan_people_count"`
		NeededPositionsCount int       `json:"needed_positions_count"`
		ItemsCount           int       `json:"items_count"`
		TotalLength          int       `json:"total_length"`
		CanViewOrder         bool      `json:"can_view_order"`
		MultiDay             bool      `json:"multi_day"`
		PrefersOrderView     bool      `json:"prefers_order_view"`
		Rehearsable          bool      `json:"rehearsable"`
		FilesExpireAt        time.Time `json:"files_expire_at"`
		SortDate             time.Time `json:"sort_date"`
		LastTimeAt           time.Time `json:"last_time_at"`
		Permissions          string    `json:"permissions"`
		Dates                string    `json:"dates"`
		ShortDates           string    `json:"short_dates"`
		PlanningCenterURL    string    `json:"planning_center_url"`
		RemindersDisabled    bool      `json:"reminders_disabled"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type PlanNote struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		CategoryName string    `json:"category_name"`
		Content      string    `json:"content"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type PlanNoteCategory struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt time.Time `json:"created_at"`
		DeletedAt time.Time `json:"deleted_at"`
		Name      string    `json:"name"`
		Sequence  int       `json:"sequence"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type PlanPerson struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Status                    string    `json:"status"`
		CreatedAt                 time.Time `json:"created_at"`
		UpdatedAt                 time.Time `json:"updated_at"`
		Notes                     string    `json:"notes"`
		DeclineReason             string    `json:"decline_reason"`
		Name                      string    `json:"name"`
		NotificationChangedByName string    `json:"notification_changed_by_name"`
		NotificationSenderName    string    `json:"notification_sender_name"`
		TeamPositionName          string    `json:"team_position_name"`
		PhotoThumbnail            string    `json:"photo_thumbnail"`
		StatusUpdatedAt           time.Time `json:"status_updated_at"`
		NotificationChangedAt     time.Time `json:"notification_changed_at"`
		NotificationPreparedAt    time.Time `json:"notification_prepared_at"`
		NotificationReadAt        time.Time `json:"notification_read_at"`
		PrepareNotification       bool      `json:"perpare_notification"`
		CanAcceptPartial          bool      `json:"can_accept_partial"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type PlanPersonTime struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Status    string    `json:"status"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type PlanTemplate struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name             string    `json:"name"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
		ItemCount        int       `json:"item_count"`
		TeamCount        int       `json:"team_count"`
		NoteCount        int       `json:"note_count"`
		CanViewOrder     bool      `json:"can_view_order"`
		MultiDay         bool      `json:"multi_day"`
		PrefersOrderView bool      `json:"prefers_order_view"`
		Rehearsable      bool      `json:"rehearsable"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type PlanTime struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt     time.Time     `json:"created_at"`
		UpdatedAt     time.Time     `json:"updated_at"`
		Name          string        `json:"name"`
		TimeType      int           `json:"time_type"`
		Recorded      bool          `json:"recorded"`
		TeamReminders []interface{} `json:"team_reminders"`
		StartsAt      time.Time     `json:"starts_at"`
		EndsAt        time.Time     `json:"ends_at"`
		LiveStartsAt  time.Time     `json:"live_starts_at"`
		LiveEndsAt    time.Time     `json:"live_ends_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type PublicView struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		SeriesAndPlanTitles bool `json:"series_and_plan_titles"`
		ItemLengths         bool `json:"item_lengths"`
		ServiceTimes        bool `json:"service_times"`
		SongItems           bool `json:"song_items"`
		MediaItems          bool `json:"media_items"`
		RegularItems        bool `json:"regular_items"`
		Headers             bool `json:"headers"`
		Itunes              bool `json:"itunes"`
		Amazon              bool `json:"amazon"`
		Spotify             bool `json:"spotify"`
		Youtube             bool `json:"youtube"`
		Vimeo               bool `json:"vimeo"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type ReportTemplate struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Body    string `json:"body"`
		Title   string `json:"title"`
		Type    string `json:"type"`
		Default bool   `json:"default"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Schedule struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		SortDate                       time.Time `json:"sort_date"`
		Dates                          string    `json:"dates"`
		DeclineReason                  string    `json:"decline_reason"`
		OrganizationName               string    `json:"organization_name"`
		OrganizationTimeZone           string    `json:"organization_time_zone"`
		OrganizationTwentyFourHourTime string    `json:"organization_twenty_four_hour_time"`
		PersonName                     string    `json:"person_name"`
		PositionDisplayTimes           string    `json:"position_display_times"`
		RespondsToName                 string    `json:"responds_to_name"`
		ServiceTypeName                string    `json:"service_type_name"`
		ShortDates                     string    `json:"short_dates"`
		Status                         string    `json:"status"`
		TeamName                       string    `json:"team_name"`
		TeamPositionName               string    `json:"team_position_name"`
		CanAcceptPartial               bool      `json:"can_accept_partial"`
		CanAcceptPartialOneTime        bool      `json:"can_accept_partial_one_time"`
		CanRehearse                    bool      `json:"can_rehearse"`
		PlanVisible                    bool      `json:"plan_visible"`
		PlanVisibleToMe                bool      `json:"plan_visible_to_me"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type ScheduledPerson struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		FullName  string `json:"full_name"`
		Status    string `json:"status"`
		Thumbnail string `json:"thumbnail"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type SchedulingPreference struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Preference string `json:"preference"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Series struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt           time.Time `json:"created_at"`
		UpdatedAt           time.Time `json:"updated_at"`
		ArtworkFileName     string    `json:"artwork_file_name"`
		ArtworkContentType  string    `json:"artwork_content_type"`
		ArtworkFileSize     int       `json:"artwork_file_size"`
		Title               string    `json:"title"`
		ArtworkForDashboard string    `json:"artwork_for_dashboard"`
		ArtworkForMobile    string    `json:"artwork_for_mobile"`
		ArtworkForPlan      string    `json:"artwork_for_plan"`
		ArtworkOriginal     string    `json:"artwork_original"`
		HasArtwork          bool      `json:"has_artwork"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type ServiceType struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		ArchivedAt                 time.Time `json:"archived_at"`
		CreatedAt                  time.Time `json:"created_at"`
		DeletedAt                  time.Time `json:"deleted_at"`
		Name                       string    `json:"name"`
		Sequence                   int       `json:"sequence"`
		UpdatedAt                  time.Time `json:"updated_at"`
		AttachmentTypesEnabled     bool      `json:"attachment_types_enabled"`
		BackgroundCheckPermissions string    `json:"background_check_permissions"`
		CommentPermissions         string    `json:"comment_permissions"`
		CustomItemTypes            string    `json:"custom_item_types"`
		Frequency                  string    `json:"frequency"`
		LastPlanFrom               string    `json:"last_plan_from"`
		Permissions                string    `json:"permissions"`
		StandardItemTypes          string    `json:"standard_item_types"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type ServiceTypePath struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Path []interface{} `json:"path"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type SignupSheet struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		SortDate     time.Time `json:"sort_date"`
		GroupKey     string    `json:"group_key"`
		TeamName     string    `json:"team_name"`
		DisplayTimes string    `json:"display_times"`
		PositionName string    `json:"position_name"`
		Title        string    `json:"title"`
		SortIndex    int       `json:"sort_index"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type SignupSheetMetadata struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Conflicts interface{} `json:"conflicts"`
		TimeType  string      `json:"time_type"`
		TimeName  string      `json:"time_name"`
		StartsAt  time.Time   `json:"starts_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type SkippedAttachment struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Skipped bool `json:"skipped"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Song struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Title                   string    `json:"title"`
		CreatedAt               time.Time `json:"created_at"`
		UpdatedAt               time.Time `json:"updated_at"`
		Admin                   string    `json:"admin"`
		Author                  string    `json:"author"`
		Copyright               string    `json:"copyright"`
		Hidden                  bool      `json:"hidden"`
		Notes                   string    `json:"notes"`
		Themes                  string    `json:"themes"`
		LastScheduledShortDates string    `json:"last_scheduled_short_dates"`
		LastScheduledAt         time.Time `json:"last_scheduled_at"`
		CCLINumber              int       `json:"ccli_number"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type SongSchedule struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		ArrangementName string `json:"arrangement_name"`
		KeyName         string `json:"key_name"`
		PlanDates       string `json:"plan_dates"`
		ServiceTypeName string `json:"service_type_name"`
		PlanSortDate    string `json:"plan_sort_date"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type SongbookStatus struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Status      string `json:"status"`
		StatusCode  string `json:"status_code"`
		StatusToken string `json:"status_token"`
		Url         string `json:"url"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type SplitTeamRehearsalAssignment struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		ScheduleSpecialServiceTimes bool `json:"schedule_special_service_times"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Tag struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name string `json:"name"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type TagGroup struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name                    string `json:"name"`
		Required                bool   `json:"required"`
		AllowMultipleSelections bool   `json:"allow_multiple_selections"`
		TagsFor                 string `json:"tags_for"`
		ServiceTypeFolderName   string `json:"service_type_folder_name"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Team struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name                        string    `json:"name"`
		RehearsalTeam               bool      `json:"rehearsal_team"`
		Sequence                    int       `json:"sequence"`
		ScheduleTo                  string    `json:"schedule_to"`
		DefaultStatus               string    `json:"default_status"`
		DefaultPrepareNotifications bool      `json:"default_prepare_notifications"`
		CreatedAt                   time.Time `json:"created_at"`
		UpdatedAt                   time.Time `json:"updated_at"`
		ArchivedAt                  time.Time `json:"archived_at"`
		AssignedDirectly            bool      `json:"assigned_directly"`
		SecureTeam                  bool      `json:"secure_team"`
		LastPlanFrom                string    `json:"last_plan_from"`
		StageColor                  string    `json:"stage_color"`
		StageVariant                string    `json:"stage_variant"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type TeamLeader struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		SendResponsesForAccepts   bool `json:"send_responses_for_accepts"`
		SendResponsesForDeclines  bool `json:"send_responses_for_declines"`
		SendResponsesForBlockouts bool `json:"send_responses_for_blockouts"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type TeamPosition struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name              string        `json:"name"`
		Tags              []interface{} `json:"tags"`
		NegativeTagGroups []interface{} `json:"negative_tag_groups"`
		TagGroups         []interface{} `json:"tag_groups"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type TextSetting struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		SchedulingRequestsEnabled bool   `json:"scheduling_requests_enabled"`
		GeneralEmailsEnabled      bool   `json:"general_emails_enabled"`
		SchedulingRepliesEnabled  bool   `json:"scheduling_replies_enabled"`
		RemindersEnabled          bool   `json:"reminders_enabled"`
		Carrier                   string `json:"carrier"`
		DisplayNumber             string `json:"display_number"`
		NormalizedNumber          string `json:"normalized_number"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type TimePreferenceOption struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		DayOfWeek   int       `json:"day_of_week"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		Description string    `json:"description"`
		SortIndex   string    `json:"sort_index"`
		TimeType    string    `json:"time_type"`
		MinuteOfDay int       `json:"minute_of_day"`
		StartsAt    time.Time `json:"starts_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Zoom struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		AspectRatio float32 `json:"aspect_ratio"`
		ZoomLevel   float32 `json:"zoom_level"`
		XOffset     float32 `json:"x_offset"`
		YOffset     float32 `json:"y_offset"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}
