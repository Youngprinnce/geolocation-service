package location

import (
	"math"
	"time"
)

// Location represents a geographical location with coordinates
type Location struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"unique;not null" binding:"required"`
	Latitude  float64   `json:"latitude" gorm:"not null" binding:"required,min=-90,max=90"`
	Longitude float64   `json:"longitude" gorm:"not null" binding:"required,min=-180,max=180"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateLocationRequest represents the request body for creating a location
type CreateLocationRequest struct {
	Name      string  `json:"name" binding:"required"`
	Latitude  float64 `json:"latitude" binding:"required,min=-90,max=90"`
	Longitude float64 `json:"longitude" binding:"required,min=-180,max=180"`
}

// DistanceCalculator provides methods for calculating distances between coordinates
type DistanceCalculator struct{}

// HaversineDistance calculates the distance between two points using the Haversine formula
// Returns distance in kilometers
func (dc *DistanceCalculator) HaversineDistance(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371 // Earth's radius in kilometers

	// Convert degrees to radians
	lat1Rad := lat1 * math.Pi / 180
	lon1Rad := lon1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lon2Rad := lon2 * math.Pi / 180

	// Calculate differences
	dlat := lat2Rad - lat1Rad
	dlon := lon2Rad - lon1Rad

	// Haversine formula
	a := math.Sin(dlat/2)*math.Sin(dlat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(dlon/2)*math.Sin(dlon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := R * c

	return distance
}
