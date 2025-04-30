package store

import(
	"Assessment/store/booking"
	"Assessment/store/service"
	"Assessment/store/vendor"
)

	type Store struct {
		BookingRepo booking.BookingRepository
		ServiceRepo service.ServiceRepository
		VendorRepo  vendor.VendorRepository
}