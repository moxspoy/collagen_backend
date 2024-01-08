package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SetupTestDB creates a new in-memory database for testing purposes
func SetupTestDB() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
}
