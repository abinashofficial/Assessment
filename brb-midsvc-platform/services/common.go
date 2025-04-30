package services

import(
	"Assessment/services/booking"
	"Assessment/services/service"
	"Assessment/services/vendor"
)

type Store struct {
	BookingService booking.BookingService
	ServiceService service.ServiceService	
	VendorService  vendor.VendorService
}