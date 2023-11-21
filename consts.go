package main

// ConvertedRequest represents the structure of the converted request
type ConvertedRequest struct {
	Event       string                `json:"event"`
	EventType   string                `json:"event_type"`
	AppID       string                `json:"app_id"`
	UserID      string                `json:"user_id"`
	MessageID   string                `json:"message_id"`
	PageTitle   string                `json:"page_title"`
	PageURL     string                `json:"page_url"`
	BrowserLang string                `json:"browser_language"`
	ScreenSize  string                `json:"screen_size"`
	Attributes  map[string]formValues `json:"attributes"`
	UserTraits  map[string]formValues `json:"traits"`
}

var basicForm1 = []string{
	"uatrk",
	"uatrv",
	"uatrt",
}
var basicForm2 = []string{
	"atrk",
	"atrv",
	"atrt",
}

var basicForm3 = []string{
	"ev",
	"et",
	"id",
	"uid",
	"mid",
	"t",
	"p",
	"l",
	"sc",
}

type formValues struct {
	FormType  string `json:"type"`
	FormValue string `json:"value"`
}
