# BRB Mid Service Platform

This project is a GoLang microservice for managing service listings, vendor bookings, and basic user flows (admin/customer), built for Beauty Right Back (BRB).

It covers:
- Service and vendor management
- Basic booking with non-overlapping rules
- Notification simulation
- Simple user roles
- Booking summary per vendor
- Swagger documentation
- Basic retry handler and rate limiter
- Pagination for customer booking lists
- Docker-based deployment

---

## 🚀 Setup Instructions

### 1. Clone the repository
```bash
git clone https://github.com/yourusername/brb-midsvc-platform.git
cd brb-midsvc-platform


2. Run with Docker Compose

docker-compose up --build

This will spin up:
    PostgreSQL database
    GoLang API server on localhost:8080


📖 Swagger/OpenAPI Documentation
Once the service is running:

👉 Access Swagger UI here: http://localhost:8080/swagger/index.html

It contains:
    All API endpoints
    Request/response models
    Example payloads


⚙️ API Endpoints Overview

Method  | Endpoint                                  | Description

POST    | /services                                 | Create new service
PUT     | /services/:id                             | Update service
POST    | /vendors/:vendor_id/services/:service_id  | Link service to vendor
PATCH   | /services/:id/availability                | Toggle service availability
POST    | /bookings                                 | Book a service
GET     | /bookings/customer/:customer_id           | Get bookings (paginated)
GET     | /summary/vendor/:vendor_id                | Booking summary by vendor
GET     | /health                                   | Health check

