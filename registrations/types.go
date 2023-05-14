package registrations

import (
	"fmt"
	"pcgo"
	"time"
)

type Organization struct {
	ID   string `json:"id"`
	Type string `json:"type"`

	Links map[string]string `json:"links"`

	Attributes struct {
		CreatedAt  time.Time `json:"created_at"`
		DateFormat string    `json:"date_format"`
		Name       string    `json:"name"`
		UpdatedAt  time.Time `json:"updated_at"`
	} `json:"attributes"`
}

type AttendeeType struct {
	ID   string `json:"id"`
	Type string `json:"type"`

	Links         map[string]string         `json:"links"`
	Relationships map[string]map[string]any `json:"relationships"`

	Attributes struct {
		ActiveAttendeesOnWaitlistCount int       `json:"active_attendees_on_waitlist_count"`
		AnyActiveAttendeesOnWaitlist   bool      `json:"any_active_attendees_on_waitlist"`
		AtMaximumCapacity              bool      `json:"at_maximum_capacity"`
		AvailableCapacity              int       `json:"available_capacity"`
		ConstrainedCapacity            int       `json:"constrained_capacity"`
		CreatedAt                      time.Time `json:"created_at"`
		MaximumCapacity                int       `json:"maximum_capacity"`
		Name                           string    `json:"name"`
		Position                       int       `json:"position"`
		PriceCents                     int       `json:"price_cents"`
		PriceString                    string    `json:"price_string"`
		UpdatedAt                      time.Time `json:"updated_at"`
		Waitlist                       bool      `json:"waitlist"`
	} `json:"attributes"`
}

type Attendee struct {
	ID    string            `json:"id"`
	Type  string            `json:"type"`
	Links map[string]string `json:"links"`

	Relationships map[string]map[string]any `json:"relationships"`

	Attributes struct {
		CanceledOn                  time.Time `json:"canceled_on"`
		CreatedAt                   time.Time `json:"created_at"`
		DeletedAt                   time.Time `json:"deleted_at"`
		EmergencyContactName        string    `json:"emergency_contact_name"`
		EmergencyContactPhoneNumber string    `json:"emergency_contact_phone_number"`
		FinalizationState           string    `json:"finalization_state"`
		MissingPersonalInfo         bool      `json:"missing_personal_info"`
		Name                        string    `json:"name"`
		UpdatedAt                   time.Time `json:"updated_at"`
	} `json:"attributes"`
}

type Event struct {
	ID    string            `json:"id"`
	Type  string            `json:"type"`
	Links map[string]string `json:"links"`

	Relationships struct {
		Contact pcgo.RelationshipNode `json:"contact"`
	} `json:"relationships"`
	// Relationships map[string]map[string]any `json:"relationships"`

	Attributes struct {
		ActiveAttendeesCount int       `json:"active_attendees_count"`
		ArchivedOn           time.Time `json:"archived_on"`
		AtMaximumCapacity    bool      `json:"at_maximum_capacity"`
		CloseAt              time.Time `json:"close_at"`
		Closed               bool      `json:"closed"`
		CreatedAt            time.Time `json:"created_at"`
		CurrentVisibility    string    `json:"current_visibility"`
		Description          string    `json:"description"`
		DirectLinkOnly       bool      `json:"direct_link_only"`
		EventTimeSummary     string    `json:"event_time_summary"`
		Featured             bool      `json:"featured"`
		HideAt               time.Time `json:"hide_at"`
		LinkOnly             bool      `json:"link_only"`
		Listed               bool      `json:"listed"`
		LogoUrl              string    `json:"logo_url"`
		MaximumCapacity      int       `json:"maximum_capacity"`
		Name                 string    `json:"name"`
		Open                 bool      `json:"open"`
		OpenAt               time.Time `json:"open_at"`
		PublicUrl            string    `json:"public_url"`
		RegistrationType     string    `json:"registration_type"`
		ShowAt               time.Time `json:"show_at"`
		UpdatedAt            time.Time `json:"updated_at"`
		Visibility           string    `json:"visibility"`
	} `json:"attributes"`
}

func (e Event) String() string {
	return fmt.Sprintf(
		`
    ID: %s
    Type: %s
    Links: %v

    Relationships:
      Contact:
        Data:
          ID: %s
          Type: %s

    Attributes:
      ActiveAttendeesCount: %v
      ArchivedOn: %v
      AtMaximumCapacity: %v
      CloseAt: %v
      Closed: %v
      CreatedAt: %v
      CurrentVisibility: %v
      Description: %v
      DirectLinkOnly: %v
      EventTimeSummary: %v
      Featured: %v
      HideAt: %v
      LinkOnly: %v
      Listed: %v
      LogoUrl: %v
      MaximumCapacity: %v
      Name: %v
      Open: %v
      OpenAt: %v
      PublicUrl: %v
      RegistrationType: %v
      ShowAt: %v
      UpdatedAt: %v
      Visibility: %v

		`,
		e.ID,
		e.Type,
		e.Links,
		e.Relationships.Contact.Data.ID,
		e.Relationships.Contact.Data.Type,

		e.Attributes.ActiveAttendeesCount,
		e.Attributes.ArchivedOn,
		e.Attributes.AtMaximumCapacity,
		e.Attributes.CloseAt,
		e.Attributes.Closed,
		e.Attributes.CreatedAt,
		e.Attributes.CurrentVisibility,
		e.Attributes.Description,
		e.Attributes.DirectLinkOnly,
		e.Attributes.EventTimeSummary,
		e.Attributes.Featured,
		e.Attributes.HideAt,
		e.Attributes.LinkOnly,
		e.Attributes.Listed,
		e.Attributes.LogoUrl,
		e.Attributes.MaximumCapacity,
		e.Attributes.Name,
		e.Attributes.Open,
		e.Attributes.OpenAt,
		e.Attributes.PublicUrl,
		e.Attributes.RegistrationType,
		e.Attributes.ShowAt,
		e.Attributes.UpdatedAt,
		e.Attributes.Visibility,
	)
}

type Person struct {
	ID    string            `json:"id"`
	Type  string            `json:"type"`
	Links map[string]string `json:"links"`

	Attributes struct {
		AvatarUrl   string    `json:"avatar_url"`
		CreatedAt   time.Time `json:"created_at"`
		FirstName   string    `json:"first_name"`
		LastName    string    `json:"last_name"`
		Name        string    `json:"name"`
		Permissions string    `json:"permissions"`
	} `json:"attributes"`
}

type Registration struct {
	ID    string            `json:"id"`
	Type  string            `json:"type"`
	Links map[string]string `json:"links"`

	Relationships map[string]map[string]any `json:"relationships"`

	Attributes struct {
		AvailableCount    string      `json:"available_count"`
		BalanceDueCents   int         `json:"balance_due_cents"`
		CanceledOn        time.Time   `json:"canceled_on"`
		CompletionStatus  string      `json:"completion_status"`
		CreatedAt         time.Time   `json:"created_at"`
		CreatedBy         interface{} `json:"created_by"`
		CreationChannel   string      `json:"creation_channel"`
		DeletedAt         time.Time   `json:"deleted_at"`
		FinalizationState string      `json:"finalization_state"`
		MissingFormsCount int         `json:"missing_forms_count"`
		UpdatedAt         time.Time   `json:"updated_at"`
		WaitlistCount     int         `json:"waitlist_count"`
	} `json:"attributes"`
}
