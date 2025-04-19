package event

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	CreateEvent(event Event) (Event, error)
	GetEventByID(id uint) (Event, error)
	GetAllEvents() ([]Event, error)
	UpdateEvent(id uint, event Event) (Event, error)
	DeleteEvent(id uint) error
	GetEventsByUserID(userID uint) ([]Event, error)
	GetEventsByDate(date time.Time) ([]Event, error)
	GetEventsByLocation(location string) ([]Event, error)
	CreateBulkEvents(events []Event) ([]Event, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateEvent(event Event) (Event, error) {
	if err := r.db.Create(&event).Error; err != nil {
		log.Printf("Failed to create event: %v", err)
		return Event{}, err
	}
	return event, nil
}
func (r *repository) GetEventByID(id uint) (Event, error) {
	var event Event
	if err := r.db.First(&event, id).Error; err != nil {
		log.Printf("Failed to get event by ID: %v", err)
		return Event{}, err
	}
	return event, nil
}
func (r *repository) GetAllEvents() ([]Event, error) {
	var events []Event
	if err := r.db.Find(&events).Error; err != nil {
		log.Printf("Failed to get all events: %v", err)
		return nil, err
	}
	return events, nil
}
func (r *repository) UpdateEvent(id uint, event Event) (Event, error) {
	if err := r.db.Model(&Event{}).Where("id = ?", id).Updates(event).Error; err != nil {
		log.Printf("Failed to update event: %v", err)
		return Event{}, err
	}
	return event, nil
}
func (r *repository) DeleteEvent(id uint) error {
	if err := r.db.Delete(&Event{}, id).Error; err != nil {
		log.Printf("Failed to delete event: %v", err)
		return err
	}
	return nil
}
func (r *repository) GetEventsByUserID(userID uint) ([]Event, error) {
	var events []Event
	if err := r.db.Where("user_id = ?", userID).Find(&events).Error; err != nil {
		log.Printf("Failed to get events by user ID: %v", err)
		return nil, err
	}
	return events, nil
}
func (r *repository) GetEventsByDate(date time.Time) ([]Event, error) {
	var events []Event
	if err := r.db.Where("DATE(start_time) = ?", date).Find(&events).Error; err != nil {
		log.Printf("Failed to get events by date: %v", err)
		return nil, err
	}
	return events, nil
}
func (r *repository) GetEventsByLocation(location string) ([]Event, error) {
	var events []Event
	if err := r.db.Where("location = ?", location).Find(&events).Error; err != nil {
		log.Printf("Failed to get events by location: %v", err)
		return nil, err
	}
	return events, nil
}
func (r *repository) CreateBulkEvents(events []Event) ([]Event, error) {
	if err := r.db.CreateInBatches(events, 100).Error; err != nil {
		log.Printf("Failed to create bulk events: %v", err)
		return nil, err
	}
	return events, nil
}
