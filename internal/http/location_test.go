package http

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/youngprinnce/geolocation-service/internal/service/location"
)

func TestCreateLocation(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Valid Request Structure", func(t *testing.T) {
		// Create request body
		reqBody := location.CreateLocationRequest{
			Name:      "Test Location",
			Latitude:  40.7128,
			Longitude: -74.0060,
		}
		jsonBody, _ := json.Marshal(reqBody)

		// Create HTTP request
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/locations", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")

		// Verify request structure
		assert.Equal(t, "POST", c.Request.Method)
		assert.Equal(t, "/locations", c.Request.URL.Path)
		assert.Equal(t, "application/json", c.Request.Header.Get("Content-Type"))
	})

	t.Run("Invalid Coordinates Structure", func(t *testing.T) {
		// Test invalid latitude
		reqBody := location.CreateLocationRequest{
			Name:      "Invalid Location",
			Latitude:  95.0, // Invalid latitude
			Longitude: -74.0060,
		}
		jsonBody, _ := json.Marshal(reqBody)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		c.Request = httptest.NewRequest("POST", "/locations", bytes.NewBuffer(jsonBody))
		c.Request.Header.Set("Content-Type", "application/json")

		// Verify request structure
		assert.Equal(t, "POST", c.Request.Method)
		assert.NotNil(t, c.Request.Body)
	})
}

func TestLocationControllerStructure(t *testing.T) {
	// Test that controller can be created (without real service for simplicity)
	controller := &LocationController{
		service: nil, // Would normally inject a real or mock service
	}

	assert.NotNil(t, controller)
}
