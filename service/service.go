package service

import (
    "Assessment/models"
    "Assessment/repository"
    "errors"
    "log"
)

type Service struct {
    Repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
    return &Service{Repo: repo}
}

// Service Management
func (s *Service) CreateService(service *models.Service) error {
    return s.Repo.CreateService(service)
}

func (s *Service) UpdateService(service *models.Service) error {
    return s.Repo.UpdateService(service)
}

func (s *Service) ToggleService(id uint) error {
    return s.Repo.ToggleService(id)
}

// Booking
func (s *Service) CreateBooking(booking *models.Booking) error {
    // Booking must be between 9 AM - 5 PM
    hour := booking.StartTime.Hour()
    if hour < 9 || hour > 16 {
        return errors.New("booking must be between 9 AM to 5 PM")
    }

    // Check for conflicts
    existing, err := s.Repo.GetBookingsByVendorAndTime(booking.VendorID, booking.StartTime, booking.EndTime)
    if err != nil {
        return err
    }
    if len(existing) > 0 {
        return errors.New("time slot already booked")
    }

    booking.Status = "pending"
    err = s.Repo.CreateBooking(booking)
    if err == nil {
        log.Printf("Booking created: %+v", booking)
    }
    return err
}

func (s *Service) ListBookings(customerID uint, offset, limit int) ([]models.Booking, error) {
    return s.Repo.ListCustomerBookings(customerID, offset, limit)
}

func (s *Service) VendorSummary(vendorID uint) (int64, map[string]int64, error) {
    return s.Repo.VendorSummary(vendorID)
}
