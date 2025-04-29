package config

import (
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "os"
    "Assessment/models"
    "log"
    "fmt"
)

func InitDB() (*gorm.DB, error) {
    dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )
    if dsn == "" {
        dsn = "host=localhost user=postgres password=love dbname=postgres port=5432 sslmode=disable"
    }
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }

        // Auto migrate tables
        err = db.AutoMigrate(
            &models.User{},
            &models.Vendor{},
            &models.Service{},
            &models.Booking{},
        )
        if err != nil {
            log.Fatal("Failed to auto-migrate database: ", err)
        }
    return db, nil
}
