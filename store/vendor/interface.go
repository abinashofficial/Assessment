package vendor

import "Assessment/model"

type VendorRepository interface {
	Create(vendor *model.Vendor) error
}
