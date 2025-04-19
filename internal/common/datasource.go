package common

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

// PGPASSWORD=7Nq5gF5hOXOrflibp8ugGVs1KDgRnaWa psql -h dpg-cvvfforuibrs73bdogkg-a.oregon-postgres.render.com -U test_mg7t_user test_mg7t
func InitDB() (*gorm.DB, error) {
	dsn := "host=dpg-cvvfforuibrs73bdogkg-a.oregon-postgres.render.com user=test_mg7t_user password=7Nq5gF5hOXOrflibp8ugGVs1KDgRnaWa dbname=test_mg7t port=5432 sslmode=require TimeZone=Africa/Lagos"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	// Configure connection pool
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(10)           // Maximum number of open connections
	sqlDB.SetMaxIdleConns(5)            // Maximum number of idle connections
	sqlDB.SetConnMaxLifetime(time.Hour) // Maximum lifetime of a connection

	return db, nil
}
