package groups

import "time"

type Attendance struct {
	Type string `json:"attendance"`
	ID   string `json:"id"`

	Attributes struct {
		Attended bool   `json:"attended"`
		Role     string `json:"role"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Event struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		AttendanceRequestsEnabled bool      `json:"attendance_requests_enabled"`
		AutomatedReminderEnabled  bool      `json:"automated_reminder_enabled"`
		Canceled                  bool      `json:"canceled"`
		CanceledAt                time.Time `json:"canceled_at"`
		Description               string    `json:"description"`
		EndsAt                    time.Time `json:"ends_at"`
		LocationTypePreference    string    `json:"location_type_preference"`
		MultiDay                  bool      `json:"multi_day"`
		Name                      string    `json:"name"`
		RemindersSent             bool      `json:"reminders_sent"`
		RemindersSentAt           time.Time `json:"reminders_sent_at"`
		Repeating                 bool      `json:"repeating"`
		StartsAt                  time.Time `json:"starts_at"`
		VirtualLocationUrl        string    `json:"virtual_location_url"`
		VisitorsCount             int       `json:"visitors_count"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type EventNote struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Body string `json:"body"`
	}

	Relationships interface{} `json:"relationships"`
}

type Group struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		ArchivedAt               time.Time   `json:"archived_at"`
		ContactEmail             string      `json:"contact_email"`
		CreatedAt                time.Time   `json:"created_at"`
		Description              string      `json:"description"`
		EnrollmentOpen           bool        `json:"enrollment_open"`
		EnrollmentStrategy       string      `json:"enrollment_strategy"`
		EventsVisibility         string      `json:"events_visibility"`
		HeaderImage              interface{} `json:"header_image"`
		LocationTypePreference   string      `json:"location_type_preference"`
		MembershipsCount         int         `json:"memberships_count"`
		Name                     string      `json:"name"`
		PublicChurchCenterWebUrl string      `json:"public_church_center_web_url"`
		Schedule                 string      `json:"schedule"`
		VirtualLocationUrl       string      `json:"virtual_location_url"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type GroupType struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		ChurchCenterVisible    bool   `json:"church_center_visible"`
		ChurchCenterMapVisible bool   `json:"church_center_map_visible"`
		Color                  string `json:"color"`
		DefaultGroupSettings   string `json:"default_group_settings"`
		Name                   string `json:"name"`
		Position               int    `json:"position"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Location struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		DisplayPreference    string  `json:"display_preference"`
		FullFormattedAddress string  `json:"full_formatted_address"`
		Latitude             float64 `json:"latitude"`
		Longitude            float64 `json:"longitude"`
		Name                 string  `json:"name"`
		Radius               string  `json:"radius"`
		Strategy             string  `json:"strategy"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Membership struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		AccountCenterIdentifier string    `json:"account_center_identifier"`
		AvatarUrl               string    `json:"avatar_url"`
		ColorIdentifier         string    `json:"color_identifier"`
		EmailAddress            string    `json:"email_address"`
		FirstName               string    `json:"first_name"`
		JoinedAt                time.Time `json:"joined_at"`
		LastName                string    `json:"last_name"`
		PhoneNumber             string    `json:"phone_number"`
		Role                    string    `json:"role"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Organization struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name string `json:"name"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Owner struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		AvatarUrl string `json:"avatar_url"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Person struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Addresses      []string  `json:"addresses"`
		AvatarUrl      string    `json:"avatar_url"`
		CreatedAt      time.Time `json:"created_at"`
		EmailAddresses []string  `json:"email_addresses"`
		FirstName      string    `json:"first_name"`
		LastName       string    `json:"last_name"`
		Permissions    string    `json:"permissions"`
		PhoneNumbers   []string  `json:"phone_number"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Resource struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Description string    `json:"description"`
		LastUpdated time.Time `json:"last_updated"`
		Name        string    `json:"name"`
		Type        string    `json:"type"`
		Visibility  string    `json:"visibility"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Tag struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name     string `json:"name"`
		Position int    `json:"position"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type TagGroup struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		DisplayPublicly        bool   `json:"display_publicly"`
		MultipleOptionsEnabled bool   `json:"multiple_options_enabled"`
		Name                   string `json:"name"`
		Position               int    `json:"position"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}
