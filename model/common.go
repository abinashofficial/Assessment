package model

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoClient struct {
	PrimaryClient   *mongo.Client
	SecondaryClient *mongo.Client
}

type MongoError struct {
	PrimaryClientError   error
	SecondaryClientError error
}

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

var BasicForm1 = []string{
	"uatrk",
	"uatrv",
	"uatrt",
}
var BasicForm2 = []string{
	"atrk",
	"atrv",
	"atrt",
}

var BasicForm3 = []string{
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

type FormValues struct {
	FormType  string `json:"type"`
	FormValue string `json:"value"`
}

var RequestChannel = make(chan map[string]string)
