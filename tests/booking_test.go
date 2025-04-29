package tests

import (
    "Assessment/models"
    "Assessment/repository"
    "Assessment/service"
    "testing"
    "time"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "github.com/stretchr/testify/assert"
)

func setupTestDB() (*gorm.DB, error) {
    db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
    // if err != nil {
    //     return nil, err
    // }
    db.AutoMigrate(&models.User{}, &models.Vendor{}, &models.Service{}, &models.Booking{})
    return db, nil
}

func TestCreateBooking(t *testing.T) {
    db, _ := setupTestDB()
    repo := repository.NewRepository(db)
    svc := service.NewService(repo)

    vendor := models.Vendor{Name: "Test Vendor"}
    // db.Create(&vendor)

    booking := &models.Booking{
        CustomerID: 1,
        VendorID:   vendor.ID,
        StartTime:  time.Date(2025, 4, 28, 10, 0, 0, 0, time.UTC),
        EndTime:    time.Date(2025, 4, 28, 11, 0, 0, 0, time.UTC),
    }

    err := svc.CreateBooking(booking)

    assert.Nil(t, err)
    assert.Equal(t, "pending", booking.Status)
}
