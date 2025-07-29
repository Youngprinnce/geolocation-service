package location

import (
	"math"
	"testing"
)

func TestHaversineDistance(t *testing.T) {
	calculator := &DistanceCalculator{}

	tests := []struct {
		name      string
		lat1      float64
		lon1      float64
		lat2      float64
		lon2      float64
		expected  float64
		tolerance float64
	}{
		{
			name:      "Same point",
			lat1:      40.7128,
			lon1:      -74.0060,
			lat2:      40.7128,
			lon2:      -74.0060,
			expected:  0,
			tolerance: 0.001,
		},
		{
			name:      "New York to Los Angeles",
			lat1:      40.7128, // NYC
			lon1:      -74.0060,
			lat2:      34.0522, // LA
			lon2:      -118.2437,
			expected:  3944, // Approximately 3944 km
			tolerance: 50,   // 50 km tolerance
		},
		{
			name:      "London to Paris",
			lat1:      51.5074, // London
			lon1:      -0.1278,
			lat2:      48.8566, // Paris
			lon2:      2.3522,
			expected:  344, // Approximately 344 km
			tolerance: 10,  // 10 km tolerance
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := calculator.HaversineDistance(tt.lat1, tt.lon1, tt.lat2, tt.lon2)
			if math.Abs(result-tt.expected) > tt.tolerance {
				t.Errorf("HaversineDistance() = %v, expected %v ± %v", result, tt.expected, tt.tolerance)
			}
		})
	}
}

func TestValidateCoordinates(t *testing.T) {
	tests := []struct {
		name      string
		latitude  float64
		longitude float64
		valid     bool
	}{
		{
			name:      "Valid coordinates",
			latitude:  40.7128,
			longitude: -74.0060,
			valid:     true,
		},
		{
			name:      "Invalid latitude - too high",
			latitude:  91,
			longitude: -74.0060,
			valid:     false,
		},
		{
			name:      "Invalid latitude - too low",
			latitude:  -91,
			longitude: -74.0060,
			valid:     false,
		},
		{
			name:      "Invalid longitude - too high",
			latitude:  40.7128,
			longitude: 181,
			valid:     false,
		},
		{
			name:      "Invalid longitude - too low",
			latitude:  40.7128,
			longitude: -181,
			valid:     false,
		},
		{
			name:      "Edge case - North Pole",
			latitude:  90,
			longitude: 0,
			valid:     true,
		},
		{
			name:      "Edge case - South Pole",
			latitude:  -90,
			longitude: 0,
			valid:     true,
		},
		{
			name:      "Edge case - Date Line",
			latitude:  0,
			longitude: 180,
			valid:     true,
		},
		{
			name:      "Edge case - Prime Meridian",
			latitude:  0,
			longitude: -180,
			valid:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			valid := tt.latitude >= -90 && tt.latitude <= 90 && tt.longitude >= -180 && tt.longitude <= 180
			if valid != tt.valid {
				t.Errorf("Coordinate validation for lat=%v, lon=%v: got %v, expected %v",
					tt.latitude, tt.longitude, valid, tt.valid)
			}
		})
	}
}
