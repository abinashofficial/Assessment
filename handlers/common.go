package handlers

import (
	"Assessment/handlers/booking"
	"Assessment/handlers/service"
	"Assessment/handlers/vendor"
)

type Store  struct{
	BookingHandler booking.Handler
	ServiceHandler service.Handler
	VendorHandler  vendor.Handler
}