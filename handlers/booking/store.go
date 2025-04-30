package booking

import(
	"Assessment/model"
	"Assessment/services/booking"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)



type BookingHandler struct {
	service booking.BookingService
}

func NewBookingHandler(s booking.BookingService) *BookingHandler {
	return &BookingHandler{service: s}
}

func (h *BookingHandler) CreateBooking(c *gin.Context) {
	var booking model.Booking
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.CreateBooking(&booking)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (h *BookingHandler) ListBookings(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	bookings, err := h.service.ListBookings(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch bookings"})
		return
	}

	c.JSON(http.StatusOK, bookings)
}

func (h *BookingHandler) VendorSummary(c *gin.Context) {
	id := c.Param("id")
	total, statusCounts, err := h.service.GetVendorSummary(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch summary"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"total_bookings": total,
		"status_counts":  statusCounts,
	})
}

