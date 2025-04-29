package utils

import (
    "errors"
    "fmt"
    "time"
	"Assessment/models"
)

// Retry attempts to run a function multiple times until it succeeds or max retries reached
func Retry(attempts int, sleep time.Duration, fn func() error) error {
    var err error
    for i := 0; i < attempts; i++ {
        err = fn()
        if err == nil {
            return nil
        }
        fmt.Printf("[Retry] Attempt %d failed: %s\n", i+1, err.Error())
        time.Sleep(sleep)
    }
    return fmt.Errorf("after %d attempts, last error: %s", attempts, err)
}

func SimulateSendNotification(booking *models.Booking) error {
    // Fake random failure simulation:
    if time.Now().UnixNano()%2 == 0 { // randomly fail
        return errors.New("temporary notification failure")
    }
    
    fmt.Printf("[Notification] Booking Created Successfully: BookingID=%d, VendorID=%d\n", booking.ID, booking.VendorID)
    return nil
}
