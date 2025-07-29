package manualwire

import (
	"github.com/youngprinnce/geolocation-service/internal/http"
	"github.com/youngprinnce/geolocation-service/internal/postgres"
	"github.com/youngprinnce/geolocation-service/internal/service/location"
)

func GetLocationRepository() location.LocationStore {
	session := postgres.GetSession()
	return location.NewLocationRepo(session)
}

func GetLocationService(repo location.LocationStore, calculator *location.DistanceCalculator) location.LocationBC {
	return location.NewLocationService(repo, calculator)
}

func GetLocationController() *http.LocationController {
	repo := GetLocationRepository()
	service := GetLocationService(repo, GetLocationDistanceCalculator())
	return http.NewLocationController(service)
}

func GetLocationDistanceCalculator() *location.DistanceCalculator {
	return &location.DistanceCalculator{}
}
