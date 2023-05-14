package api

type OauthApplication struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name        string `json:"name"`
		Url         string `json:"url"`
		Description string `json:"description"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type OauthApplicationMau struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Count int `json:"count"`
		Month int `json:"month"`
		Year  int `json:"year"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type Organization struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes interface{} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}

type PersonalAccessToken struct {
	Type string `json:"type"`
	ID   string `json:"id"`

	Attributes struct {
		Name string `json:"name"`
	} `json:"attributes"`

	Relationships interface{} `json:"relationships"`
}
