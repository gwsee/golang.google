package model

type Paging struct {
	Cursors Cursors `json:"cursors"`
	Next    string  `json:"next"`
}

type Cursors struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

type ListAll struct {
	Data   []interface{} `json:"data"`
	Paging Paging        `json:"paging"`
}

type Auth struct {
	RedirectURL         string
	DeveloperToken      string
	AdwordsClientID     string
	AdwordsClientSecret string
	GaClientID          string
	GaClientSecret      string
	OauthURI            string
}
