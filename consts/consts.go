package consts

import "errors"

var NoReplyEmail = ""

var (
	WebUrl                    = ""
	PublicCommunicationWebUrl = ""
	PortalLink                = ""
	FeedbackBaseUrl           = ""
	EmailAttachmentsBucket    = ""
	AWSRegion                 = ""
)

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
	Attributes  map[string]FormValues `json:"attributes"`
	UserTraits  map[string]FormValues `json:"traits"`
}

type FormValues struct {
	FormType  string `json:"type"`
	FormValue string `json:"value"`
}

const TraceID = "traceid"

const (
	ErrUserEmailNotInHeader         = "ErrUserEmailNotInHeader"
	CorsOriginMatchAll       string = "*"
	SupportPortalDB                 = "support_portal"
	FormsColl                       = "forms"
	ErrMongoDisconnect              = "ErrMongoDisconnect"
	ErrMongoConn                    = "ErrMongoConn"
	ErrMySQLConnectionFailed        = "ErrMySQLConnectionFailed"
	ErrTicketIDCreation             = "ErrTicketIDCreation"
	ErrGetFailed                    = "ErrGetFailed"
)

// for env based initializations
func EnvSetup(portalLink, feedbackUrl string) error {

	if portalLink == "" {
		return errors.New("Support Portal Link is not set")
	}

	PortalLink = portalLink

	//if feedbackUrl == "" {
	//	return errors.New("FEEDBACK_URL is not set")
	//}

	FeedbackBaseUrl = feedbackUrl

	return nil
}
