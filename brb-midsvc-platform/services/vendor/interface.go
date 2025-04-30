package vendor

import (
	"Assessment/model"
)

type VendorService interface {
	CreateVendor(vendor *model.Vendor) error
}