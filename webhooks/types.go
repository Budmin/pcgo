package webhooks

import "time"

type AvailableEvent struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name     string `json:"name"`
		App      string `json:"app"`
		Version  string `json:"version"`
		Type     string `json:"type"`
		Resource string `json:"resource"`
		Action   string `json:"action"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Delivery struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Status          int       `json:"status"`
		RequestHeaders  string    `json:"request_headers"`
		RequestBody     string    `json:"request_body"`
		ResponseHeaders string    `json:"response_headers"`
		ResponseBody    string    `json:"response_body"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
		Timing          float32   `json:"timing"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Event struct {
	Type string `json:"Type"`
	ID   string `json:"id"`

	Attributes struct {
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		UUID      string    `json:"uuid"`
		Payload   string    `json:"payload"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Organization struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes interface{} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type WebhookSubscription struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name               string    `json:"name"`
		Url                string    `json:"url"`
		Active             bool      `json:"active"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		AuthenticitySecret string    `json:"authenticity_secret"`
		ApplicationId      string    `json:"application_id"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}
