package location

type LocationBC interface {
	CreateLocation(req CreateLocationRequest) (*Location, error)
	GetAllLocations() ([]Location, error)
	FindNearestLocation(lat, lng float64) (*Location, float64, error)
	DeleteLocationByName(name string) error
}

// Service handles location-related business logic
type LocationService struct {
	repo       LocationStore
	Calculator *DistanceCalculator
}

// NewLocationService creates a new location service
func NewLocationService(repo LocationStore, calculator *DistanceCalculator) LocationBC {
	return &LocationService{
		repo:       repo,
		Calculator: calculator,
	}
}

// CreateLocation handles the business logic for creating a location
func (s *LocationService) CreateLocation(req CreateLocationRequest) (*Location, error) {
	// Check if name already exists
	exists, err := s.repo.NameExists(req.Name)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, &DuplicateNameError{Name: req.Name}
	}

	location := &Location{
		Name:      req.Name,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}

	if err := s.repo.Create(location); err != nil {
		return nil, err
	}

	return location, nil
}

// GetAllLocations returns all locations
func (s *LocationService) GetAllLocations() ([]Location, error) {
	return s.repo.GetAll()
}

// FindNearestLocation finds the nearest location to given coordinates
func (s *LocationService) FindNearestLocation(lat, lng float64) (*Location, float64, error) {
	locations, err := s.repo.GetAll()
	if err != nil {
		return nil, 0, err
	}

	if len(locations) == 0 {
		return nil, 0, &NoLocationsError{}
	}

	// Find the nearest location
	var nearest *Location
	var minDistance float64

	for i, location := range locations {
		distance := s.Calculator.HaversineDistance(lat, lng, location.Latitude, location.Longitude)
		if i == 0 || distance < minDistance {
			minDistance = distance
			nearest = &locations[i]
		}
	}

	return nearest, minDistance, nil
}

// DeleteLocationByName deletes a location by name
func (s *LocationService) DeleteLocationByName(name string) error {
	// Check if location exists
	_, err := s.repo.GetByName(name)
	if err != nil {
		return err
	}

	return s.repo.DeleteByName(name)
}

// Custom error types for better error handling
type DuplicateNameError struct {
	Name string
}

func (e *DuplicateNameError) Error() string {
	return "Location name already exists: " + e.Name
}

type NoLocationsError struct{}

func (e *NoLocationsError) Error() string {
	return "No locations found"
}

// ValidationError for coordinate validation
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return e.Field + ": " + e.Message
}

// ValidateCoordinates validates latitude and longitude
func ValidateCoordinates(lat, lng float64) error {
	if lat < -90 || lat > 90 {
		return &ValidationError{Field: "latitude", Message: "must be between -90 and 90"}
	}
	if lng < -180 || lng > 180 {
		return &ValidationError{Field: "longitude", Message: "must be between -180 and 180"}
	}
	return nil
}
