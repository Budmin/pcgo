package checkins

import (
	"time"

	"github.com/Budmin/pcgo"
)

// TODO:
// when parsing the JSON, you should use a custom time.Time unmarshaler function or use a library like
// https://github.com/golang/time that can parse timestamps in ISO format.

type AttendanceType struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name      string    `json:"name"`
		Color     string    `json:"color"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Limit     int       `json:"limit"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type CheckIn struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		FirstName                   string    `json:"first_name"`
		LastName                    string    `json:"last_name"`
		MedicalNotes                string    `json:"medical_notes"`
		Number                      int       `json:"number"`
		SecurityCode                string    `json:"security_code"`
		CreatedAt                   time.Time `json:"created_at"`
		UpdatedAt                   time.Time `json:"updated_at"`
		CheckedOutAt                time.Time `json:"checked_out_at"`
		ConfirmedAt                 time.Time `json:"confirmed_at"`
		EmergencyContactName        string    `json:"emergency_contact_name"`
		EmergencyContactPhoneNumber string    `json:"emergency_contact_phone_number"`
		OneTimeGuest                bool      `json:"one_time_guest"`
		Kind                        string    `json:"kind"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type CheckInGroup struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		NameLabelsCount     int       `json:"name_labels_count"`
		SecurityLabelsCount int       `json:"security_labels_count"`
		CheckInsCount       int       `json:"check_ins_count"`
		PrintStatus         string    `json:"print_status"`
		CreatedAt           time.Time `json:"created_at"`
		UpdatedAt           time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type CheckInTime struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Kind               string   `json:"kind"`
		HasValidated       bool     `json:"has_validated"`
		Errors             []string `json:"errors"`
		ServicesIntegrated bool     `json:"services_integrated"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type Event struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name                      string    `json:"name"`
		Frequency                 string    `json:"frequency"`
		EnableServicesIntegration bool      `json:"enable_services_integration"`
		CreatedAt                 time.Time `json:"created_at"`
		UpdatedAt                 time.Time `json:"updated_at"`
		ArchivedAt                time.Time `json:"archived_at"`
		IntegrationKey            string    `json:"integration_key"`
		LocationTimesEnabled      bool      `json:"location_times_enabled"`
		PreSelectEnabled          bool      `json:"pre_select_enabled"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type EventLabel struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Quantity     int       `json:"quantity"`
		ForRegular   bool      `json:"for_regular"`
		ForGuest     bool      `json:"for_guest"`
		ForVolunteer bool      `json:"for_volunteer"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type EventPeriod struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		StartsAt       time.Time `json:"starts_at"`
		EndsAt         time.Time `json:"ends_at"`
		RegularCount   int       `json:"regular_count"`
		GuestCount     int       `json:"guest_count"`
		VolunteerCount int       `json:"volunteer_count"`
		Note           string    `json:"note"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type EventTime struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		TotalCount     int       `json:"total_count"`
		StartsAt       time.Time `json:"starts_at"`
		ShowsAt        time.Time `json:"shows_at"`
		HidesAt        time.Time `json:"hides_at"`
		RegularCount   int       `json:"regular_count"`
		GuestCount     int       `json:"guest_count"`
		VolunteerCount int       `json:"volunteer_count"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
		Name           string    `json:"name"`
		Hour           int       `json:"hour"`
		Minute         int       `json:"minute"`
		DayOfWeek      int       `json:"day_of_week"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type Headcount struct {
	ID   string `json:"id"`
	Type string `json:"type"`

	Attributes struct {
		Total     int       `json:"total"`
		UpdatedAt time.Time `json:"updated_at"`
		CreatedAt time.Time `json:"created_at"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type Label struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	Attributes struct {
		Name      string    `json:"name"`
		Xml       string    `json:"xml"`
		PrintsFor string    `json:"prints_for"`
		Roll      string    `json:"roll"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Location struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes struct {
		Name                  string    `json:"name"`
		Kind                  string    `json:"kind"`
		Opened                bool      `json:"opened"`
		Questions             string    `json:"questions"`
		AgeMinInMonths        int       `json:"age_min_in_months"`
		AgeMaxInMonths        int       `json:"age_max_in_months"`
		AgeRangeBy            string    `json:"age_range_by"`
		AgeOn                 string    `json:"age_on"`
		ChildOrAdult          string    `json:"child_or_adult"`
		EffectiveDate         string    `json:"effective_date"`
		Gender                string    `json:"gender"`
		GradeMin              int       `json:"grade_min"`
		GradeMax              int       `json:"grade_max"`
		MaxOccupancy          int       `json:"max_occupancy"`
		MinVolunteers         int       `json:"min_volunteers"`
		AttendeesPerVolunteer int       `json:"attendees_per_volunteer"`
		Position              int       `json:"position"`
		UpdatedAt             time.Time `json:"updated_at"`
		CreatedAt             time.Time `json:"created_at"`
	} `json:"attributes"`

	Relationships map[string]pcgo.RelationshipNode `json:"relationships"`
}

type LocationEventPeriod struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		RegularCount   int       `json:"regular_count"`
		GuestCount     int       `json:"guest_count"`
		VolunteerCount int       `json:"volunteer_count"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type LocationEventTime struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		RegularCount   int       `json:"regular_count"`
		GuestCount     int       `json:"guest_count"`
		VolunteerCount int       `json:"volunteer_count"`
		CreatedAt      time.Time `json:"created_at"`
		UpdatedAt      time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type LocationLabel struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Quantity     int       `json:"quantity"`
		ForRegular   bool      `json:"for_regular"`
		ForGuest     bool      `json:"for_guest"`
		ForVolunteer bool      `json:"for_volunteer"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Option struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Body      string    `json:"body"`
		Quantity  int       `json:"quantity"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Organization struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		DateFormatPattern string    `json:"date_format_pattern"`
		TimeZoneOlson     string    `json:"time_zone_olson"`
		Name              string    `json:"name"`
		DailyCheckIns     int       `json:"daily_check_ins"`
		TimeZone          string    `json:"time_zone"`
		AvatarUrl         string    `json:"avatar_url"`
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Pass struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Code      string    `json:"code"`
		Kind      string    `json:"kind"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Person struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Addresses             []interface{} `json:"addresses"`
		EmailAddresses        []interface{} `json:"email_addresses"`
		PhoneNumbers          []interface{} `json:"phone_numbers"`
		AvatarURL             string        `json:"avatar_url"`
		NamePrefix            string        `json:"name_prefix"`
		FirstName             string        `json:"first_name"`
		MiddleName            string        `json:"middle_name"`
		LastName              string        `json:"last_name"`
		NameSuffix            string        `json:"name_suffix"`
		Birthdate             string        `json:"birthdate"`
		Grade                 int           `json:"grade"`
		Gender                string        `json:"gender"`
		MedicalNotes          string        `json:"medical_notes"`
		Child                 bool          `json:"child"`
		Permission            string        `json:"permission"`
		Headcounter           bool          `json:"headcounter"`
		LastCheckedInAt       time.Time     `json:"last_checked_in_at"`
		CheckInCount          int           `json:"check_in_count"`
		CreatedAt             time.Time     `json:"created_at"`
		UpdatedAt             time.Time     `json:"updated_at"`
		PassedBackgroundCheck bool          `json:"passed_background_check"`
		DemographicAvatarUrl  string        `json:"demographic_avatar_url"`
		Name                  string        `json:"name"`
		TopPermission         string        `json:"top_permission"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type PersonEvent struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		CheckInCount int       `json:"check_in_count"`
		UpdatedAt    time.Time `json:"updated_at"`
		CreatedAt    time.Time `json:"created_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type RosterListPerson struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		FirstName            string    `json:"first_name"`
		LastName             string    `json:"last_name"`
		Name                 string    `json:"name"`
		DemographicAvatarUrl string    `json:"demographic_avatar_url"`
		Grade                string    `json:"grade"`
		Gender               string    `json:"gender"`
		MedicalNotes         string    `json:"medical_notes"`
		Birthdate            time.Time `json:"birthdate"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Station struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Online           bool      `json:"online"`
		Mode             int       `json:"mode"`
		Name             string    `json:"name"`
		TimeoutSeconds   int       `json:"timeout_seconds"`
		InputType        string    `json:"input_type"`
		InputTypeOptions string    `json:"input_type_options"`
		CreatedAt        time.Time `json:"created_at"`
		UpdatedAt        time.Time `json:"updated_at"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Theme struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		ImageThumbnail  string    `json:"image_thumbnail"`
		Name            string    `json:"name"`
		Color           string    `json:"color"`
		TextColor       string    `json:"text_color"`
		Image           string    `json:"image"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
		BackgroundColor string    `json:"background_color"`
		Mode            string    `json:"mode"`
	}

	Relationships interface{} `json:"relationships"`
}
