package booking

import (
	"Assessment/model"
	"time"
)

type BookingRepository interface {
	CreateBooking(booking *model.Booking) error
	FindOverlappingBooking(vendorID uint, bookingTime time.Time) (*model.Booking, error)
	ListBookings(offset, limit int) ([]model.Booking, error)
	GetBookingSummary(vendorID string) (int64, map[string]int64, error)
}
