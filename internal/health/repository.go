package health

import (
	"gorm.io/gorm"
)

type Repository interface {
	GetHealthCheck() error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) GetHealthCheck() error {
	db, err := r.db.DB()
	if err != nil {
		return err
	}
	if err := db.Ping(); err != nil {
		return err
	}
	return nil
}
