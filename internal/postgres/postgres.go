package postgres

import (
	"fmt"

	"github.com/youngprinnce/geolocation-service/config"
	"github.com/youngprinnce/geolocation-service/internal/logger"
	"github.com/youngprinnce/geolocation-service/internal/service/location"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var session *gorm.DB

func GetSession() *gorm.DB {
	return session
}

func GetDB() *gorm.DB {
	return session
}

func Load(config *config.Config) error {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Database.Host,
		config.Database.Port,
		config.Database.User,
		config.Database.Password,
		config.Database.DbName)

	// Connect to the database
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	// Run automatic migrations
	logger.Info("Running database auto-migrations...")
	if err := db.AutoMigrate(&location.Location{}); err != nil {
		return fmt.Errorf("failed to run auto-migrations: %w", err)
	}
	logger.Info("Database auto-migrations completed successfully")

	session = db.Session(&gorm.Session{})

	logger.Info("Successfully initialized Postgres")
	return nil
}
