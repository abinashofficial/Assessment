package vendor

import (
	"Assessment/model"
	"gorm.io/gorm"
)

type vendorRepository struct {
	db *gorm.DB
}

func NewVendorRepository(db *gorm.DB) VendorRepository {
	return &vendorRepository{db}
}

func (r *vendorRepository) Create(vendor *model.Vendor) error {
	return r.db.Create(vendor).Error
}
