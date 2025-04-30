package booking

import(
	"Assessment/model"
)

type BookingService interface {
	CreateBooking(booking *model.Booking) (*model.Booking, error)
	ListBookings(page, limit int) ([]model.Booking, error)
	GetVendorSummary(vendorID string) (int64, map[string]int64, error)
}