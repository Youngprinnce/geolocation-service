# Geolocation Service API

A RESTful API in Golang that allows users to register geolocated stations and query for the nearest station to a given point.

## 🚀 Features

- **POST /locations** - Register new geolocated stations
- **GET /locations** - Get all registered locations
- **GET /locations/nearest?lat=LAT&lng=LNG** - Find nearest station to given coordinates
- **DELETE /locations/{name}** - Delete station by name
- Uses **Haversine formula** for accurate distance calculations
- **PostgreSQL** database for persistence
- Comprehensive input validation and error handling
- Clean architecture with proper separation of concerns

## 📋 Requirements Met

### ✅ Functionality
- [x] POST /locations with validation (lat: -90 to 90, lng: -180 to 180, unique names)
- [x] GET /nearest with Haversine formula distance calculation
- [x] GET /locations returning all registered locations
- [x] DELETE /locations/{name} for removing stations by name

### ✅ Technical Requirements
- [x] Idiomatic Golang code with clean architecture
- [x] PostgreSQL database for persistence with GORM ORM
- [x] Standard Go folder structure
- [x] Appropriate HTTP status codes (200, 201, 400, 404, 409, 500)
- [x] Input validation and clean error handling

### ✅ Testing
- [x] Unit tests for distance calculation logic
- [x] Unit tests for coordinate validation
- [x] HTTP handler tests with httptest
- [x] Service layer tests with business logic validation

### ✅ Bonus Features
- [x] Docker/Docker Compose for easy setup
- [x] Structured logging with logrus (info/debug/error levels)
- [x] Clean architecture with separated concerns

## 🛠 Setup Instructions

### Prerequisites
- Go 1.23+ installed
- Docker and Docker Compose installed
- Git installed

### 🎯 Quick Start (Recommended for Hiring Teams)

**One command to rule them all:**

```bash
git clone <repository-url> && cd geolocation-service && make docker-up
```

That's it! The application will:
- ✅ Start PostgreSQL database
- ✅ Automatically run database migrations
- ✅ Start the API server on http://localhost:8080

**Test that it's working:**
```bash
curl http://localhost:8080/
# Should return: Hello!
```

### 🔧 Alternative Setup Methods

#### Option 1: Local Development
```bash
# 1. Set up PostgreSQL database locally:
#    - Install PostgreSQL
#    - Create database: createdb geolocation_db
#    - Update config-local.yaml with your database credentials

# 2. Run the application (auto-migrations will run)
make run
```

#### Option 2: Manual Docker Build
```bash
# Build and start manually
make docker-up
```

### 📋 Configuration

The app uses `config.yaml` for Docker Compose (no changes needed).
For local development, update `config-local.yaml` with your database details:

```yaml
database:
  host: "localhost"
  port: 5432
  user: "your_postgres_user"
  password: "your_postgres_password"
  db_name: "geolocation_db"
```

## 📚 API Usage Examples

### 1. Register a New Location

```bash
curl -X POST http://localhost:8080/locations \
  -H "Content-Type: application/json" \
  -d '{
    "name": "CentralStation",
    "latitude": 40.7128,
    "longitude": -74.0060
  }'
```

**Response (201 Created):**

```json
{
  "id": 1,
  "name": "CentralStation",
  "latitude": 40.7128,
  "longitude": -74.0060,
  "created_at": "2025-01-28T10:00:00Z",
  "updated_at": "2025-01-28T10:00:00Z"
}
```

### 2. Get All Locations

```bash
curl http://localhost:8080/locations
```

**Response (200 OK):**

```json
[
  {
    "id": 1,
    "name": "CentralStation",
    "latitude": 40.7128,
    "longitude": -74.0060,
    "created_at": "2025-01-28T10:00:00Z",
    "updated_at": "2025-01-28T10:00:00Z"
  }
]
```

### 3. Find Nearest Location

```bash
curl "http://localhost:8080/locations/nearest?lat=40.7589&lng=-73.9851"
```

**Response (200 OK):**

```json
{
  "location": {
    "id": 1,
    "name": "CentralStation",
    "latitude": 40.7128,
    "longitude": -74.0060,
    "created_at": "2025-01-28T10:00:00Z",
    "updated_at": "2025-01-28T10:00:00Z"
  },
  "distance_km": 2.84
}
```

### 4. Delete a Location

```bash
curl -X DELETE "http://localhost:8080/locations/CentralStation"
```

**Response (200 OK):**

```json
{
  "message": "Location deleted successfully"
}
```

## 🧪 Testing

### Run All Tests

```bash
# Run all tests
make test

# Or directly
go test ./... -v
```

## 🎯 Architecture

This project follows clean architecture principles:

### Layers

- **Route Layer** (`cmd/server/app.go`): Configures routes, middleware, and HTTP setup
- **Controller Layer** (`internal/http/`): Handles HTTP requests, validation, and responses  
- **Service Layer** (`internal/service/location/`): Contains business logic and rules
- **Data Layer** (`internal/service/location/store.go`): Database operations and data access

### Key Components

- **Route Configuration**: Centralized in `cmd/server/app.go` for all API endpoints
- **HTTP Controllers**: Pure request/response handlers in `internal/http/`
- **Business Services**: Domain logic implementation in `internal/service/`
- **Data Stores**: Database access layer with interface abstraction
- **Models**: Data structures and domain entities

## 📊 Distance Calculation

The service uses the **Haversine formula** to calculate distances between coordinates:

```text
distance = 2 * R * arcsin(sqrt(sin²((lat2-lat1)/2) + cos(lat1) * cos(lat2) * sin²((lng2-lng1)/2)))
```

Where R is the Earth's radius (6,371 km).

## 🗄 Database Schema

```sql
CREATE TABLE locations (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    latitude DECIMAL(10, 8) NOT NULL,
    longitude DECIMAL(11, 8) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```
