package vendor

import (
	"Assessment/model"
	"Assessment/services/vendor"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VendorHandler struct {
	service vendor.VendorService
}

func NewVendorHandler(s vendor.VendorService) *VendorHandler {
	return &VendorHandler{service: s}
}

func (h *VendorHandler) RegisterRoutes(r *gin.RouterGroup) {
	vendor := r.Group("/vendors")
	{
		vendor.POST("/", h.CreateVendor)
	}
}

func (h *VendorHandler) CreateVendor(c *gin.Context) {
	var vendor model.Vendor
	if err := c.ShouldBindJSON(&vendor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateVendor(&vendor); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create vendor"})
		return
	}
	c.JSON(http.StatusCreated, vendor)
}
