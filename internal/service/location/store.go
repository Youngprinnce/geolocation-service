package location

import (
	"gorm.io/gorm"
)

// LocationStore defines the interface for location data access
type LocationStore interface {
	Create(location *Location) error
	GetAll() ([]Location, error)
	GetByName(name string) (*Location, error)
	DeleteByName(name string) error
	NameExists(name string) (bool, error)
}

// LocationRepo provides data access methods for locations
type LocationRepo struct {
	db *gorm.DB
}

// NewLocationRepo creates a new location repository
func NewLocationRepo(db *gorm.DB) LocationStore {
	return &LocationRepo{
		db: db,
	}
}

// Create creates a new location in the database
func (s *LocationRepo) Create(location *Location) error {
	return s.db.Create(location).Error
}

// GetAll retrieves all locations from the database
func (s *LocationRepo) GetAll() ([]Location, error) {
	var locations []Location
	err := s.db.Find(&locations).Error
	return locations, err
}

// GetByName retrieves a location by name
func (s *LocationRepo) GetByName(name string) (*Location, error) {
	var location Location
	err := s.db.Where("name = ?", name).First(&location).Error
	if err != nil {
		return nil, err
	}
	return &location, nil
}

// DeleteByName deletes a location by name
func (s *LocationRepo) DeleteByName(name string) error {
	return s.db.Where("name = ?", name).Delete(&Location{}).Error
}

// NameExists checks if a location with the given name exists
func (s *LocationRepo) NameExists(name string) (bool, error) {
	var count int64
	err := s.db.Model(&Location{}).Where("name = ?", name).Count(&count).Error
	return count > 0, err
}
