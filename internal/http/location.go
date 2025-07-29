package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/youngprinnce/geolocation-service/internal/service/location"
	"gorm.io/gorm"
)

// LocationController handles HTTP requests for location endpoints
type LocationController struct {
	service location.LocationBC
}

// NewLocationController creates a new location controller
func NewLocationController(service location.LocationBC) *LocationController {
	return &LocationController{
		service: service,
	}
}

// CreateLocation handles POST /locations
func (h *LocationController) CreateLocation(c *gin.Context) {
	var req location.CreateLocationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.WithError(err).Error("Failed to bind location request")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate coordinates
	if err := location.ValidateCoordinates(req.Latitude, req.Longitude); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdLocation, err := h.service.CreateLocation(req)
	if err != nil {
		switch err.(type) {
		case *location.DuplicateNameError:
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		default:
			log.WithError(err).Error("Failed to create location")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create location"})
			return
		}
	}

	log.WithFields(log.Fields{
		"name":      createdLocation.Name,
		"latitude":  createdLocation.Latitude,
		"longitude": createdLocation.Longitude,
	}).Info("Location created successfully")

	c.JSON(http.StatusCreated, createdLocation)
}

// GetLocations handles GET /locations
func (h *LocationController) GetLocations(c *gin.Context) {
	locations, err := h.service.GetAllLocations()
	if err != nil {
		log.WithError(err).Error("Failed to get locations")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get locations"})
		return
	}

	c.JSON(http.StatusOK, locations)
}

// GetNearest handles GET /nearest?lat=LAT&lng=LNG
func (h *LocationController) GetNearest(c *gin.Context) {
	latStr := c.Query("lat")
	lngStr := c.Query("lng")

	if latStr == "" || lngStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "lat and lng query parameters are required"})
		return
	}

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid latitude value"})
		return
	}

	lng, err := strconv.ParseFloat(lngStr, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid longitude value"})
		return
	}

	// Validate coordinates
	if err := location.ValidateCoordinates(lat, lng); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	nearest, distance, err := h.service.FindNearestLocation(lat, lng)
	if err != nil {
		switch err.(type) {
		case *location.NoLocationsError:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		default:
			log.WithError(err).Error("Failed to find nearest location")
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find nearest location"})
			return
		}
	}

	response := gin.H{
		"location":    nearest,
		"distance_km": distance,
	}

	c.JSON(http.StatusOK, response)
}

// DeleteLocation handles DELETE /locations/{name}
func (h *LocationController) DeleteLocation(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Location name is required"})
		return
	}

	err := h.service.DeleteLocationByName(name)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
			return
		}
		log.WithError(err).Error("Failed to delete location")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete location"})
		return
	}

	log.WithField("name", name).Info("Location deleted successfully")
	c.JSON(http.StatusOK, gin.H{"message": "Location deleted successfully"})
}
