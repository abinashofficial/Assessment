package booking

import(
	"github.com/gin-gonic/gin"
)

type Handler interface {
	CreateBooking(c *gin.Context)
	ListBookings(c *gin.Context)
	VendorSummary(c *gin.Context)
}