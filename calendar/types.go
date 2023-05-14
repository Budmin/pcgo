package calendar

import (
	"pcgo"
	"time"
)

type Attachment struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		ContentType string    `json:"content-type"`
		CreatedAt   time.Time `json:"created_at"`
		Description string    `json:"description"`
		FileSize    int       `json:"file_size"`
		Name        string    `json:"name"`
		UpdatedAt   time.Time `json:"updated_at"`
		Url         string    `json:"url"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type Conflict struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt  time.Time `json:"created_at"`
		Note       string    `json:"note"`
		ResolvedAt time.Time `json:"resolved_at"`
		UpdatedAt  time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type Event struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		ApprovalStatus string    `json:"approval_status"`
		CreatedAt      time.Time `json:"created_at"`
		Description    string    `json:"description"`
		ImageUrl       string    `json:"image_url"`
		Name           string    `json:"name"`
		// TODO: are these floats or ints?
		// according to the documentation they're ints
		PercentApproved       int    `json:"percent_approved"`
		PercentRejected       int    `json:"percent_rejected"`
		RegistrationUrl       string `json:"registration_url"`
		Summary               string `json:"summary"`
		VisibleInChurchCenter bool   `json:"visible_in_church_center"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type EventConnection struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		ConnectedToId   string `json:"connected_to_id"`
		ConnectedToType string `json:"connected_to_type"`
		ProductName     string `json:"product_name"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type EventInstance struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		AllDayEvent                  bool      `json:"all_day_event"`
		CompactRecurrenceDescription string    `json:"compact_recurrence_description"`
		CreatedAt                    time.Time `json:"created_at"`
		EndsAt                       time.Time `json:"ends_at"`
		Location                     string    `json:"location"`
		Recurrence                   string    `json:"recurrence"`
		RecurrenceDescription        string    `json:"recurrence_description"`
		StartsAt                     time.Time `json:"starts_at"`
		UpdatedAt                    time.Time `json:"updated_at"`
		ChurchCenterUrl              string    `json:"church_center_url"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type EventResourceRequest struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		ApprovalSent   bool      `json:"approval_sent"`
		ApprovalStatus string    `json:"approval_status"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		Notes          string    `json:"notes"`
		Quantity       int       `json:"quantity"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type EventTime struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		EndsAt                 time.Time `json:"ends_at"`
		StartsAt               time.Time `json:"starts_at"`
		Name                   time.Time `json:"name"`
		VisibleOnKiosks        bool      `json:"visible_on_kiosks"`
		VisibleOnWidgetAndIcal bool      `json:"visible_on_widget_and_ical"`
	}

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type Feed struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		DefaultChurchCenterVisibility string    `json:"default_church_center_visibility"`
		FeedType                      string    `json:"feed_type"`
		Name                          string    `json:"name"`
		ImportedAt                    time.Time `json:"imported_at"`
		CanDelete                     bool      `json:"can_delete"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type Organization struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name               string `json:"name"`
		TimeZone           string `json:"time_zone"`
		TwentyFourHourTime bool   `json:"twenty_four_hour_time"`
		DateFormat         string `json:"date_format"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Person struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt                time.Time `json:"created_at"`
		FirstName                string    `json:"first_name"`
		LastName                 string    `json:"last_name"`
		MiddleName               string    `json:"middle_name"`
		UpdatedAt                time.Time `json:"updated_at"`
		AvatarUrl                string    `json:"avatar_url"`
		Child                    bool      `json:"child"`
		ContactData              string    `json:"contact_data"`
		Gender                   string    `json:"gender"`
		HasAccess                bool      `json:"has_access"`
		NamePrefix               string    `json:"name_prefix"`
		NameSuffix               string    `json:"name_suffix"`
		PendingRequestCount      int       `json:"pending_request_count"`
		Permissions              int       `json:"permissions"`
		ResolvesConflicts        bool      `json:"resolves_conflicts"`
		SiteAdministrator        bool      `json:"site_administrator"`
		Status                   string    `json:"status"`
		EventPermissionsType     string    `json:"event_permissions_type"`
		PeoplePermissionsType    string    `json:"people_permissions_type"`
		RoomPermissionsType      string    `json:"room_permissions_type"`
		ResourcesPermissionsType string    `json:"resources_permissions_type"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type ReportTemplate struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		Body        string    `json:"body"`
		CreatedAt   time.Time `json:"created_at"`
		Description string    `json:"description"`
		Title       string    `json:"title"`
		UpdatedAt   time.Time `json:"updated_at"`
	} `json:"attributes"`
	Relationships interface{} `json:"relationships"`
}

type RequiredApproval struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes interface{} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Resource struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt    time.Time `json:"created_at"`
		Kind         string    `json:"kind"`
		Name         string    `json:"name"`
		SerialNumber string    `json:"serial_number"`
		UpdatedAt    time.Time `json:"updated_at"`
		Description  string    `json:"description"`
		ExpiresAt    time.Time `json:"expires_at"`
		HomeLocation string    `json:"home_location"`
		Image        string    `json:"image"`
		Quantity     int       `json:"quantity"`
		PathName     string    `json:"path_name"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type ResourceApprovalGroup struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt time.Time `json:"created_at"`
		Name      string    `json:"name"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type ResourceBooking struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt time.Time `json:"created_at"`
		EndsAt    time.Time `json:"ends_at"`
		StartsAt  time.Time `json:"starts_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Quantity  int       `json:"quantity"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type ResourceFolder struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt time.Time `json:"created_at"`
		Name      string    `json:"name"`
		UpdatedAt time.Time `json:"updated_at"`
		Ancestry  string    `json:"ancestry"`
		Kind      string    `json:"kind"`
		PathName  string    `json:"path_name"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type ResourceQuestion struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt      time.Time `json:"created_at"`
		Kind           string    `json:"kind"`
		UpdatedAt      time.Time `json:"updated_at"`
		Choices        string    `json:"choices"`
		Description    string    `json:"description"`
		MultipleSelect bool      `json:"multiple_select"`
		Optional       bool      `json:"optional"`
		Position       int       `json:"position"`
		Question       string    `json:"question"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type ResourceSuggestion struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt time.Time `json:"created_at"`
		Quantity  int       `json:"quantity"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type RoomSetup struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt           time.Time `json:"created_at"`
		Name                string    `json:"name"`
		UpdatedAt           time.Time `json:"updated_at"`
		Description         string    `json:"description"`
		Diagram             string    `json:"diagram"`
		DiagramURL          string    `json:"diagram_url"`
		DiagramThumbnailURL string    `json:"diagram_thumbnail_url"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Tag struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Color     string    `json:"color"`
		CreatedAt time.Time `json:"created_at"`
		Name      string    `json:"name"`
		Position  float32   `json:"position"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type TagGroup struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt time.Time `json:"created_at"`
		Name      string    `json:"name"`
		UpdatedAt time.Time `json:"updated_at"`
		Required  bool      `json:"required"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}
