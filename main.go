// @title BRB Mid Service Platform API
// @version 1.0
// @description This is a microservice for managing services, vendors, and bookings.

// @contact.name API Support
// @contact.url http://www.brb.com/support
// @contact.email support@brb.com

// @host localhost:8080
// @BasePath /
package main

import (
    "Assessment/config"
    "Assessment/handlers"
    "Assessment/middleware"
    "Assessment/repository"
    "Assessment/service"
    "log"

    "github.com/gin-gonic/gin"
    _ "Assessment/docs" // swagger docs

    ginSwagger "github.com/swaggo/gin-swagger"
    swaggerFiles "github.com/swaggo/files"
)

func main() {
    db, err := config.InitDB()
    if err != nil {
        log.Fatal("Failed to connect to DB:", err)
    }

    repo := repository.NewRepository(db)
    svc := service.NewService(repo)
    handler := handlers.NewHandler(svc)

    r := gin.Default()

    r.GET("/health", handler.HealthCheck)

    api := r.Group("/api", middleware.RoleMiddleware())
    {
        // Admin routes
        api.POST("/services", handler.CreateService)
        api.PATCH("/services/:id/toggle", handler.ToggleService)

        // Customer routes
        api.POST("/bookings", handler.CreateBooking)
        api.GET("/bookings", handler.ListBookings)

        // Admin Summary
        api.GET("/summary/vendor/:id", handler.VendorSummary)
    }

    r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
    handler.StartBookingRateLimitCleaner()

    r.Run(":8080")
}
