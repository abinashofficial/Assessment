package repository

import (
    "Assessment/models"
    "gorm.io/gorm"
    "time"
)

type Repository struct {
    DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
    return &Repository{DB: db}
}

// Services
func (r *Repository) CreateService(s *models.Service) error {
    return r.DB.Create(s).Error
}

func (r *Repository) UpdateService(s *models.Service) error {

    var req models.Service
    if err := r.DB.First(&req, s.ID).Error; err != nil {
        return err
    }
    s.Active = req.Active
    return r.DB.Save(&s).Error
}

func (r *Repository) ToggleService(id uint) error {
    var s models.Service
    if err := r.DB.First(&s, id).Error; err != nil {
        return err
    }
    s.Active = !s.Active
    return r.DB.Save(&s).Error
}

// Bookings
func (r *Repository) CreateBooking(b *models.Booking) error {
    return r.DB.Create(b).Error
}

func (r *Repository) GetBookingsByVendorAndTime(vendorID uint, start, end time.Time) ([]models.Booking, error) {
    var bookings []models.Booking
    err := r.DB.Where("vendor_id = ? AND start_time < ? AND end_time > ?", vendorID, end, start).
        Find(&bookings).Error
    return bookings, err
}

func (r *Repository) ListCustomerBookings(customerID uint, offset, limit int) ([]models.Booking, error) {
    var bookings []models.Booking

    // err := r.DB.Where("customer_id = ?", customerID).Find(&bookings).Error
    result := r.DB.Where("customer_id = ?", customerID).
    Offset(offset).
    Limit(limit).
    Find(&bookings)
    return bookings, result.Error
}

// Summary
func (r *Repository) VendorSummary(vendorID uint) (int64, map[string]int64, error) {
    var total int64
    counts := make(map[string]int64)

    if err := r.DB.Model(&models.Booking{}).Where("vendor_id = ?", vendorID).Count(&total).Error; err != nil {
        return 0, nil, err
    }

    statuses := []string{"pending", "confirmed", "completed"}
    for _, status := range statuses {
        var count int64
        if err := r.DB.Model(&models.Booking{}).
            Where("vendor_id = ? AND status = ?", vendorID, status).
            Count(&count).Error; err != nil {
            return 0, nil, err
        }
        counts[status] = count
    }

    return total, counts, nil
}
