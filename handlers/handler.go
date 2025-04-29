// CreateService godoc
// @Summary Create a new service
// @Description Create a new service listing.
// @Tags services
// @Accept json
// @Produce json
// @Param service body models.Service true "Service data"
// @Success 200 {object} models.Service
// @Failure 400 {object} ErrorResponse
// @Router /services [post]

package handlers

import (
	"Assessment/models"
	"Assessment/service"
	"net/http"
	"strconv"
	"time"
"fmt"
	"github.com/gin-gonic/gin"
    "Assessment/utils"
)

type Handler struct {
	Svc *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{Svc: svc}
}

// Health check
func (h *Handler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

// Create Service
func (h *Handler) CreateService(c *gin.Context) {
	var svc models.Service
	if err := c.ShouldBindJSON(&svc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.Svc.CreateService(&svc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, svc)
}

// update Service
func (h *Handler) UpdateService(c *gin.Context) {
	var svc models.Service
	if err := c.ShouldBindJSON(&svc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
    id, _ := strconv.Atoi(c.Param("id"))
    svc.ID = uint(id)
	if err := h.Svc.UpdateService(&svc); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, svc)
}

// Toggle Service
func (h *Handler) ToggleService(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.Svc.ToggleService(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "service toggled"})
}

var bookingRateLimit = make(map[uint]time.Time)

func (h *Handler)StartBookingRateLimitCleaner() {
    ticker := time.NewTicker(10 * time.Minute)
    go func() {
        for range ticker.C {
            now := time.Now()
            for customerID, lastRequest := range bookingRateLimit {
                if now.Sub(lastRequest) > 10*time.Minute {
                    delete(bookingRateLimit, customerID)
                }
            }
            fmt.Println("[BookingRateLimit] Cleanup completed")
        }
    }()
}

// Create Booking
func (h *Handler) CreateBooking(c *gin.Context) {
	var req struct {
		CustomerID uint      `json:"customer_id"`
		ServiceID  uint      `json:"service_id"`
		VendorID   uint      `json:"vendor_id"`
		StartTime  time.Time `json:"start_time"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    lastRequest, exists := bookingRateLimit[req.CustomerID]
    if exists && time.Since(lastRequest) < 5*time.Second {
        c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Rate limit exceeded. Please wait."})
        return
    }
    bookingRateLimit[req.CustomerID] = time.Now()

	booking := &models.Booking{
		CustomerID: req.CustomerID,
		ServiceID:  req.ServiceID,
		VendorID:   req.VendorID,
		StartTime:  req.StartTime,
		EndTime:    req.StartTime.Add(time.Hour),
	}

	if err := h.Svc.CreateBooking(booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

        // After saving booking to DB
        err := utils.Retry(3, 2*time.Second, func() error {
            return utils.SimulateSendNotification(booking)
        })
        if err != nil {
            fmt.Println("[Notification] Failed after retries:", err)
        }
	c.JSON(http.StatusCreated, booking)
}

// List Bookings
func (h *Handler) ListBookings(c *gin.Context) {
	customerID, _ := strconv.Atoi(c.Query("customer_id"))
    pageStr := c.DefaultQuery("page", "1")
    limitStr := c.DefaultQuery("limit", "10")
    page, _ := strconv.Atoi(pageStr)
    limit, _ := strconv.Atoi(limitStr)
    offset := (page - 1) * limit
	bookings, err := h.Svc.ListBookings(uint(customerID), offset, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, bookings)
}

// Vendor Summary
func (h *Handler) VendorSummary(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	total, counts, err := h.Svc.VendorSummary(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total": total, "status_counts": counts})
}
