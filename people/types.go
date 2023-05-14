package people

import (
	"pcgo"
	"time"
)

type Address struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Links map[string]string `json:"links"`

	Attributes struct {
		City      string    `json:"city"`
		State     string    `json:"state"`
		Zip       string    `json:"zip"`
		Street    string    `json:"street"`
		Location  string    `json:"location"`
		Primary   bool      `json:"primary"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships struct {
		Person pcgo.RelationshipNode `json:"person"`
	} `json:"relationships"`
}

type AnniversaryCouples struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes    interface{} `json:"attributes"`
	Relationships interface{} `json:"relationships"`
}

type App struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type BirthdayPeople struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes    interface{} `json:"attributes"`
	Relationships interface{} `json:"relationships"`
}

type Campus struct {
	Type string `json:"campus"`
	ID   string `json:"id"`

	Attributes struct {
		Latitude               float64   `json:"latitude"`
		Longitude              float64   `json:"longitude"`
		Description            string    `json:"description"`
		Street                 string    `json:"street"`
		City                   string    `json:"city"`
		State                  string    `json:"state"`
		Zip                    string    `json:"zip"`
		Country                string    `json:"country"`
		PhoneNumber            string    `json:"phone_number"`
		Website                string    `json:"website"`
		TwentyFourHourTime     bool      `json:"twenty_four_hour_time"`
		DateFormat             int       `json:"date_format"`
		ChurchCenterEnabled    bool      `json:"church_center_enabled"`
		ContactEmailAddress    string    `json:"contact_email_address"`
		TimeZone               string    `json:"time_zone"`
		GeolocationSetManually bool      `json:"geolocation_set_manually"`
		Name                   string    `json:"name"`
		CreatedAt              time.Time `json:"created_at"`
		UpdatedAt              time.Time `json:"updated_at"`
		AvatarUrl              string    `json:"avatar_url"`
	} `json:"attributes"`

	Relationships struct {
		Organization pcgo.RelationshipNode `json:"organization"`
	} `json:"relationships"`
}

type Carrier struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Value         string `json:"value"`
		Name          string `json:"name"`
		International bool   `json:"international"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Condition struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Application          string    `json:"application"`
		DefinitionClass      string    `json:"definition_class"`
		Comparison           string    `json:"comparison"`
		Settings             string    `json:"settings"`
		DefinitionIdentifier string    `json:"definition_identifier"`
		Description          string    `json:"description"`
		CreatedAt            time.Time `json:"created_at"`
		UpdatedAt            time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships struct {
		CreatedBy pcgo.RelationshipNode `json:"created_by"`
	} `json:"relationships"`
}

type ConnectedPerson struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		GivenName        string `json:"given_name"`
		FirstName        string `json:"first_name"`
		Nickname         string `json:"nickname"`
		MiddleName       string `json:"middle_name"`
		LastName         string `json:"last_name"`
		Gender           string `json:"gender"`
		OrganizationName string `json:"organization_name"`
		OrganizationId   string `json:"organization_id"`
	} `json:"attributes"`

	Relationships struct {
		Organization pcgo.RelationshipNode `json:"organization"`
	} `json:"relationships"`
}

type Email struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Address   string    `json:"address"`
		Location  string    `json:"location"`
		Primary   bool      `json:"primary"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Blocked   bool      `json:"blocked"`
	} `json:"attributes"`

	Relationships struct {
		Person pcgo.RelationshipNode `json:"person"`
	} `json:"relationships"`
}

type FieldDatum struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Value           string      `json:"value"`
		File            interface{} `json:"file"`
		FileSize        int         `json:"file_size"`
		FileContentType string      `json:"file_content_type"`
		FileName        string      `json:"file_name"`
	} `json:"attributes"`

	Links map[string]string `json:"links"`

	Relationships struct {
		FieldDefinition pcgo.RelationshipNode `json:"field_definition"`
		Customizable    pcgo.RelationshipNode `json:"customizable"`
	} `json:"relationships"`
}

type FieldDefinition struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Links map[string]string `json:"links"`

	Attributes struct {
		DataType  string    `json:"data_type"`
		Name      string    `json:"name"`
		Sequence  int       `json:"sequence"`
		Slug      string    `json:"slug"`
		Config    string    `json:"config"`
		DeletedAt time.Time `json:"deleted_at"`
		// this is accurate to the api
		// why this isn't a string like the actual id is beyond be
		TabID int `json:"tab_id"`
	} `json:"attributes"`

	Relationships struct {
		Tab pcgo.RelationshipNode `json:"tab"`
	} `json:"relationships"`
}

type FieldOption struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Value    string `json:"value"`
		Sequence int    `json:"sequence"`
	} `json:"attributes"`

	Relationships struct {
		FieldDefinition pcgo.RelationshipNode `json:"field_definition"`
	} `json:"relationships"`
}

type Form struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name            string    `json:"name"`
		Description     string    `json:"description"`
		Active          bool      `json:"active"`
		ArchivedAt      time.Time `json:"archived_at"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
		DeletedAt       time.Time `json:"deleted_at"`
		SubmissionCount int       `json:"submission_count"`
		PublicUrl       string    `json:"public_url"`
		RecentlyViewed  bool      `json:"recently_viewed"`
		Archived        bool      `json:"archived"`
	} `json:"attributes"`

	Relationships struct {
		Campus       pcgo.RelationshipNode `json:"campus"`
		FormCategory pcgo.RelationshipNode `json:"form_category"`
	} `json:"relationships"`
}

type FormField struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		FieldType   string    `json:"field_type"`
		Label       string    `json:"label"`
		Description string    `json:"description"`
		Required    bool      `json:"required"`
		Settings    string    `json:"settings"`
		Sequence    int       `json:"sequence"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships struct {
		Form                pcgo.RelationshipNode `json:"form"`
		Fieldable           pcgo.RelationshipNode `json:"fieldable"`
		Options             pcgo.RelationshipNode `json:"options"`
		FormFieldConditions pcgo.RelationshipNode `json:"form_field_conditions"`
	} `json:"relationships"`
}

type FormFieldOption struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Label     string    `json:"label"`
		Sequence  int       `json:"sequence"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships struct {
		FormField  pcgo.RelationshipNode `json:"form_field,omitempty"`
		Optionable pcgo.RelationshipNode `json:"optionable,omitempty"`
	} `json:"relationships"`
}

type FormSubmission struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Verified             bool      `json:"verified"`
		RequiresVerification bool      `json:"requires_verification"`
		CreatedAt            time.Time `json:"created_at"`
	} `json:"attributes"`

	Relationships struct {
		Person pcgo.RelationshipNode `json:"person,omitempty"`
		Form   pcgo.RelationshipNode `json:"form,omitempty"`
	} `json:"relationships"`
}

type FormSubmissionValue struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		DisplayValue string      `json:"display_value"`
		Attachments  interface{} `json:"attachments"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type Household struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name               string    `json:"name"`
		MemberCount        int       `json:"member_count"`
		PrimaryContactName string    `json:"primary_contact_name"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		Avatar             string    `json:"avatar"`
		PrimaryContactID   string    `json:"primary_contact_id"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type HouseholdMembership struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		PersonName string `json:"person_name"`
		Pending    bool   `json:"pending"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type InactiveReason struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Value string `json:"value"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type List struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name                 string    `json:"name"`
		AutoRefresh          bool      `json:"auto_refresh"`
		Status               string    `json:"status"`
		HasInactiveResults   bool      `json:"has_inactive_results"`
		IncludeInactive      bool      `json:"include_inactive"`
		Returns              string    `json:"returns"`
		ReturnOriginalIfNone bool      `json:"return_original_if_none"`
		Subset               string    `json:"subset"`
		AutomationsActive    bool      `json:"automations_active"`
		AutomationsCount     int       `json:"automations_count"`
		Description          string    `json:"description"`
		Invalid              bool      `json:"invalid"`
		NameOrDescription    string    `json:"name_or_description"`
		RecentlyViewed       bool      `json:"recently_viewed"`
		RefreshedAt          time.Time `json:"refreshed_at"`
		Starred              bool      `json:"starred"`
		TotalPeople          int       `json:"total_people"`
		BatchCompletedAt     time.Time `json:"batch_completed_at"`
		CreatedAt            time.Time `json:"created_at"`
		UpdatedAt            time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type ListCategory struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name           string    `json:"name"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		OrganizationID string    `json:"organization_id"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type ListResult struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type ListShare struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Permission string    `json:"permission"`
		Group      string    `json:"group"`
		CreatedAt  time.Time `json:"created_at"`
		Name       string    `json:"name"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type ListStar struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt time.Time `json:"created_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type MailchimpSyncStatus struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Status      string    `json:"status"`
		Error       string    `json:"error"`
		Progress    int       `json:"progress"`
		CompletedAt time.Time `json:"completed_at"`
		SegmentID   int       `json:"segment_id"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type MaritalStatus struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Value string `json:"value"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Message struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Kind                        string    `json:"kind"`
		ToAddresses                 string    `json:"to_addresses"`
		Subject                     string    `json:"subject"`
		File                        string    `json:"file"`
		DeliveryStatus              string    `json:"delivery_status"`
		RejectReason                string    `json:"reject_reason"`
		CreatedAt                   time.Time `json:"created_at"`
		SentAt                      time.Time `json:"sent_at"`
		BouncedAt                   time.Time `json:"bounced_at"`
		RejectionNotificationSentAt time.Time `json:"rejection_notification_sent_at"`
		FromName                    string    `json:"from_name"`
		FromAddress                 string    `json:"from_address"`
		ReadAt                      time.Time `json:"read_at"`
		AppName                     string    `json:"app_name"`
		MessageType                 string    `json:"message_type"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type MessageGroup struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		UUID          string    `json:"uuid"`
		MessageType   string    `json:"message_type"`
		FromAddress   string    `json:"from_address"`
		Subject       string    `json:"subject"`
		MessageCount  int       `json:"message_count"`
		SystemMessage bool      `json:"system_message"`
		CreatedAt     time.Time `json:"created_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type NamePrefix struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Value string `json:"value"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type NameSuffix struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Value string `json:"value"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Note struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Note           string    `json:"note"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		DisplayDate    time.Time `json:"display_date"`
		NoteCategoryID string    `json:"note_category_id"`
		OrganizationID string    `json:"organization_id"`
		PersonID       string    `json:"person_id"`
		CreatedByID    string    `json:"created_by_id"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type NoteCategory struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name           string    `json:"name"`
		Locked         bool      `json:"locked"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		OrganizationID string    `json:"organization_id"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type NoteCategoryShare struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Group      string `json:"group"`
		Permission string `json:"permission"`
		PersonID   string `json:"person_id"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type NoteCategorySubscription struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type Organization struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name           string    `json:"name"`
		CountryCode    string    `json:"country_code"`
		DateFormat     int       `json:"date_format"`
		TimeZone       string    `json:"time_zone"`
		ContactWebsite string    `json:"contact_website"`
		CreatedAt      time.Time `json:"created_at"`
		AvatarUrl      string    `json:"avatar_url"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type OrganizationStatistics struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Age map[string]struct {
			Male       int `json:"male"`
			Female     int `json:"female"`
			Unassigned int `json:"unassigned"`
		} `json:"age"`

		Campuses []interface{} `json:"campuses"`

		CreatedLast30Days int  `json:"created_last_30_days"`
		Elasticsearch     bool `json:"elasticsearch"`
		Gender            struct {
			Male       int `json:"male"`
			Female     int `json:"female"`
			Unassigned int `json:"unassigned"`
		} `json:"gender"`

		Membership []interface{} `json:"membership"`

		Total int `json:"total"`
	} `json:"attributes"`

	Meta struct {
		Parent struct {
			ID   string `json:"id"`
			Type string `json:"type"`
		} `json:"parent"`
	} `json:"meta"`
}

type PeopleImport struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Attribs     string    `json:"attribs"`
		Status      string    `json:"status"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
		ProcessedAt time.Time `json:"processed_at"`
		UndoneAt    time.Time `json:"undone_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type PeopleImportConflict struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Kind               string    `json:"kind"`
		Name               string    `json:"name"`
		Message            string    `json:"message"`
		Data               string    `json:"data"`
		ConflictingChanges string    `json:"conflicting_changes"`
		Ignore             bool      `json:"ignore"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type PeopleImportHistory struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name               string      `json:"name"`
		CreatedAt          time.Time   `json:"created_at"`
		UpdatedAt          time.Time   `json:"updated_at"`
		ConflictingChanges interface{} `json:"conflicting_changes"`
		Kind               string      `json:"kind"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Person struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		GivenName  string `json:"given_name"`
		FirstName  string `json:"first_name"`
		Nickname   string `json:"nickname"`
		MiddleName string `json:"middle_name"`
		LastName   string `json:"last_name"`
		// The person's birthdate.
		// Because it's not a timestamp,
		// you have to parse it from a string
		Birthdate string `json:"birthdate"`
		// The person's anniversary.
		// Because it's not a timestamp,
		// you have to parse it from a string
		Anniversary string `json:"anniversary"`

		Gender                  string    `json:"gender"`
		Grade                   int       `json:"grade"`
		Child                   bool      `json:"child"`
		GraduationYear          int       `json:"graduation_year"`
		SiteAdministrator       bool      `json:"site_administrator"`
		AccountingAdministrator bool      `json:"accounting_administrator"`
		PeoplePermissions       string    `json:"people_permissions"`
		Membership              string    `json:"membership"`
		InactivatedAt           time.Time `json:"inactivated_at"`
		Status                  string    `json:"status"`
		MedicalNotes            string    `json:"medical_notes"`
		CreatedAt               time.Time `json:"created_at"`
		UpdatedAt               time.Time `json:"updated_at"`
		Avatar                  string    `json:"avatar"`
		Name                    string    `json:"name"`
		DemographicAvatarUrl    string    `json:"demographic_avatar_url"`
		DirectoryStatus         string    `json:"directory_status"`
		PassedBackgroundCheck   bool      `json:"passed_background_check"`
		CanCreateForms          bool      `json:"can_create]"`
		SchoolType              string    `json:"school_type"`
		RemoteID                int       `json:"remote_id"`
	} `json:"attributes"`

	// Relationships interface{} `json:"relationships"`

	// Relationships map[string]pcgo.RelationshipNode `json:"relationships"`

	// need to finish and add everything once we're ready
	Relationships struct {
		PrimaryCampus pcgo.RelationshipNode `json:"primary_campus"`

		Gender pcgo.RelationshipNode `json:"gender"`

		Emails struct {
			Data []struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
		} `json:"emails"`

		PhoneNumbers struct {
			Data []struct {
				Type string `json:"type"`
				ID   string `json:"id"`
			} `json:"data"`
		} `json:"phone_numbers"`
	} `json:"relationships"`
}

type PersonApp struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		AllowPCOLogin     bool   `json:"allow_pco_login"`
		PeoplePermissions string `json:"people_permissions"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type PersonMerger struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt        time.Time `json:"created_at"`
		PersonToKeepID   string    `json:"person_to_keep_id"`
		PersonToRemoveID string    `json:"person_to_remove_id"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type PhoneNumber struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Links map[string]string `json:"links"`

	Attributes struct {
		Number        string    `json:"number"`
		Carrier       string    `json:"carrier"`
		Location      string    `json:"location"`
		Primary       bool      `json:"primary"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		E164          string    `json:"e164"`
		International string    `json:"international"`
		National      string    `json:"national"`
		CountryCode   string    `json:"country_code"`
	} `json:"attributes"`

	Relationships struct {
		Person pcgo.RelationshipNode `json:"person"`
	} `json:"relationships"`
}

type PlatformNotification struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		HTML string `json:"html"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Report struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name      string    `json:"name"`
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Rule struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Subset    string    `json:"subset"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type SchoolOption struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Value          string        `json:"value"`
		Sequence       int           `json:"sequence"`
		BeginningGrade string        `json:"beginning_grade"`
		EndingGrade    string        `json:"ending_grade"`
		SchoolTypes    []interface{} `json:"school_types"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type ServiceTime struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		StartTime   int    `json:"start_time"`
		Day         string `json:"day"`
		Description string `json:"description"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type SocialProfile struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Site      string    `json:"site"`
		Url       string    `json:"url"`
		Verified  bool      `json:"verified"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Tab struct {
	ID   string `json:"id"`
	Type string `json:"type"`

	Links map[string]string `json:"links"`

	Attributes struct {
		Name     string `json:"name"`
		Sequence int    `json:"sequence"`
		Slug     string `json:"slug"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Workflow struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name                          string    `json:"name"`
		MyReadyCardCount              int       `json:"my_ready_card_count"`
		TotalReadyCardCount           int       `json:"total_ready_card_count"`
		CompletedCardCount            int       `json:"completed_card_count"`
		TotalCardsCount               int       `json:"total_cards_count"`
		TotalReadyAndSnoozedCardCount int       `json:"total_ready_and_snoozed_card_count"`
		CreatedAt                     time.Time `json:"created_at"`
		UpdatedAt                     time.Time `json:"updated_at"`
		DeletedAt                     time.Time `json:"deleted_at"`
		CampusID                      string    `json:"campus_id"`
		WorkflowCategoryID            string    `json:"workflow_category_id"`
		RecentlyViewed                bool      `json:"recently_viewed"`
	}

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type WorkflowCard struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Links map[string]string `json:"links"`

	Attributes struct {
		SnoozeUntil              time.Time `json:"snooze_until"`
		Overdue                  bool      `json:"overdue"`
		Stage                    string    `json:"stage"`
		CalculatedDueAtInDaysAgo int       `json:"calculated_due_at_in_days_ago"`
		StickyAssignment         bool      `json:"sticky_assignment"`
		CreatedAt                time.Time `json:"created_at"`
		UpdatedAt                time.Time `json:"updated_at"`
		CompletedAt              time.Time `json:"completed_at"`
		FlaggedForNotificationAt time.Time `json:"flagged_for_notification_at"`
		RemovedAt                time.Time `json:"removed_at"`
		MovedToStepAt            time.Time `json:"moved_to_step_at"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type WorkflowCardActivity struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Comment               string    `json:"comment"`
		Content               string    `json:"content"`
		FormSubmissionUrl     string    `json:"form_submission_url"`
		PersonAvatarUrl       string    `json:"person_avatar_url"`
		PersonName            string    `json:"person_name"`
		ReassignedToAvatarUrl string    `json:"reassigned_to_avatar_url"`
		ReassignedToName      string    `json:"reassigned_to_name"`
		Subject               string    `json:"subject"`
		Type                  string    `json:"type"`
		ContentIsHtml         bool      `json:"content_is_html"`
		CreatedAt             time.Time `json:"created_at"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type WorkflowCardNote struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Note      string    `json:"note"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type WorkflowCategory struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name      string    `json:"name"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type WorkflowShare struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Group      string `json:"group"`
		Permission string `json:"permission"`
		PersonID   string `json:"person_id"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type WorkflowStep struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name                       string    `json:"name"`
		CreatedAt                  time.Time `json:"created_at"`
		UpdatedAt                  time.Time `json:"updated_at"`
		Sequence                   int       `json:"sequence"`
		Description                string    `json:"description"`
		AutoSnoozeDays             int       `json:"auto_snooze_days"`
		AutoSnoozeValue            int       `json:"auto_snooze_value"`
		AutoSnoozeInterval         string    `json:"auto_snooze_interval"`
		ExpectedResponseTimeInDays int       `json:"expected_response_time_in_days"`
		MyReadyCardCount           int       `json:"my_ready_card_count"`
		TotalReadyCardCount        int       `json:"total_ready_card_count"`
		DefaultAssigneeID          string    `json:"default_assignee_id"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type WorkflowStepAssigneeSummary struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		ReadyCount   int `json:"ready_count"`
		SnoozedCount int `json:"snoozed_count"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}
