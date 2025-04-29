package models

import "time"

type User struct {
    ID   uint   `gorm:"primaryKey"`
    Name string
    Role string // admin or customer
}

type Vendor struct {
    ID   uint   `gorm:"primaryKey"`
    Name     string    `json:"name"`
	Services []Service `gorm:"foreignKey:VendorID"`
}

// type Vendor struct {
// 	Name     string    `json:"name"`
// 	Services []Service `gorm:"foreignKey:VendorID"`
// }

type Service struct {
    ID       uint   `gorm:"primaryKey"`
    Name     string `json:"name"`
    VendorID []Vendor `json:"vendors"`
    Active   bool
    
}

type Booking struct {
    ID         uint      `gorm:"primaryKey"`
    CustomerID uint
    ServiceID  uint
    VendorID   uint
    StartTime  time.Time
    EndTime    time.Time
    Status     string // pending, confirmed, completed
}
