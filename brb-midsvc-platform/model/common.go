package model

import ("gorm.io/gorm"
"time")


type Service struct {
	gorm.Model
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       float64 `json:"price"`
	Active      bool    `json:"active"`
	VendorID    uint    `json:"vendor_id"`
}

type Vendor struct {
	gorm.Model
	Name     string    `json:"name"`
	Services []Service `gorm:"foreignKey:VendorID"`
}

type Booking struct {
	gorm.Model
	CustomerName string    `json:"customer_name"`
	ServiceID    uint      `json:"service_id"`
	VendorID     uint      `json:"vendor_id"`
	BookingTime  time.Time `json:"booking_time"`
	Status       string    `json:"status"` // pending, confirmed, completed
}

// TapContext contains context of client
type TapContext struct {
	DealerID       string              // Dealer ID of particular Tenant
	RoleID         string              // RoleID of client
	UserEmail      string              // Email of the user
	RequestID      string              // RequestID - used to track logs across a request-response cycle
	PermissionsMap map[string][]string // this map will help in flagging
	TapApiToken    string              // TapApiToken - used to authenticate the session/request
	Application    string              // application for dynamic application auth
	Locale         string              // Locale for language
}