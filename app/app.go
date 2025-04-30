package app

import(
	"Assessment/config"
	"github.com/gin-gonic/gin"
	book_handlers "Assessment/handlers/booking"
	service_handlers "Assessment/handlers/service"
	vendor_handlers "Assessment/handlers/vendor"
	 "Assessment/services"
	 "Assessment/handlers"
	 book_services "Assessment/services/booking"
	 service_service "Assessment/services/service"
	 vendor_service "Assessment/services/vendor"
	 repo "Assessment/store"
	 repo_booking "Assessment/store/booking"
	 repo_service "Assessment/store/service"
	 repo_vendor "Assessment/store/vendor"
	 "gorm.io/gorm"

)

var service services.Store
var h handlers.Store
var r repo.Store


func setupHandlers() {
	h = handlers.Store{
		BookingHandler: book_handlers.NewBookingHandler(service.BookingService),
		ServiceHandler: service_handlers.NewServiceHandler(service.ServiceService),	
		VendorHandler:  vendor_handlers.NewVendorHandler(service.VendorService),
	}
}

func setupService(){
	service = services.Store{
		BookingService: book_services.NewBookingService(r.BookingRepo),
		ServiceService: service_service.NewServiceService(r.ServiceRepo),
		VendorService:  vendor_service.NewVendorService(r.VendorRepo),
	}
}

func setupRepo(db *gorm.DB) {
	r = repo.Store{
		BookingRepo: repo_booking.NewBookingRepository(db),
		ServiceRepo: repo_service.NewServiceRepository(db),
		VendorRepo:  repo_vendor.NewVendorRepository(db),
	}
}


func Start() {
	db := config.SetupDatabase()
	setupRepo(db)
	setupService()
	setupHandlers()


	// Setup Gin
	g := gin.Default()
	runServer(g, db, h)
}