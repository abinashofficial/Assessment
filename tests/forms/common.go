package forms

import (
	"Assessment/model"
	"Assessment/tapcontext"
	"fmt"
	"github.com/stretchr/testify/mock"
)

type FormTestCase struct {
	Case           string
	Input          map[string]string
	MockFunctions  []func() *mock.Call
	ExpectedOutPut model.ConvertedRequest
	ExpectedError  error
}

var ctx = tapcontext.NewTapContext()

var CreateInput = map[string]string{
	"ev":  "contact_form_submitted",
	"et":  "form_submit",
	"id":  "cl_app_id_001",
	"uid": "cl_app_id_001-uid-001",
	"mid": "cl_app_id_001-uid-001",
	"t":   "Vegefoods - Free Bootstrap 4 Template by Colorlib",
	"p":   "http://shielded-eyrie-45679.herokuapp.com/contact-us",
	"l":   "en-US",

	"sc":     "1920 x 1080",
	"atrk1":  "form_varient",
	"atrv1":  "red_top",
	"atrt1":  "string",
	"atrk2":  "ref",
	"atrv2":  "XPOWJRICW993LKJD",
	"atrt2":  "string",
	"uatrk1": "name",
	"uatrv1": "iron man",
	"uatrt1": "string",
	"uatrk2": "email",
	"uatrv2": "ironman@avengers.com",
	"uatrt2": "string",
	"uatrk3": "age",
	"uatrv3": "32",
	"uatrt3": "integer",
}

var CreateOutput = model.ConvertedRequest{
	Event:       "contact_form_submitted",
	EventType:   "form_submit",
	AppID:       "cl_app_id_001",
	UserID:      "cl_app_id_001-uid-001",
	MessageID:   "cl_app_id_001-uid-001",
	PageTitle:   "Vegefoods - Free Bootstrap 4 Template by Colorlib",
	PageURL:     "http://shielded-eyrie-45679.herokuapp.com/contact-us",
	BrowserLang: "en-US",
	ScreenSize:  "1920 x 1080",
	Attributes: map[string]model.FormValues{
		"form_varient": {
			FormType:  "red_top",
			FormValue: "string",
		},
		"ref": {
			FormType:  "XPOWJRICW993LKJD",
			FormValue: "string",
		},
	},
	UserTraits: map[string]model.FormValues{
		"name": {
			FormValue: "iron man",
			FormType:  "string",
		},
		"email": {
			FormValue: "ironman@avengers.com",
			FormType:  "string",
		},
		"age": {
			FormValue: "32",
			FormType:  "integer",
		},
	},
}

var OutputError = fmt.Errorf("empty map")
