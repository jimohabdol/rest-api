package booking

import (
	"time"
	"gorm.io/gorm"
)

type Repository interface {
	CreateBooking(booking Booking) (Booking, error)
	GetBookingByID(id uint) (Booking, error)
	GetAllBookings() ([]Booking, error)
	UpdateBooking(id uint, booking Booking) (Booking, error)
	DeleteBooking(id uint) error
	GetBookingsByUserID(userID uint) ([]Booking, error)
	GetBookingsByEventID(eventID uint) ([]Booking, error)
	GetBookingsByDate(date time.Time) ([]Booking, error)
	GetBookingsByStatus(status string) ([]Booking, error)
	GetBookingsByPaymentStatus(status string) ([]Booking, error)
	GetBookingsByDateRange(startDate, endDate time.Time) ([]Booking, error)
}

type repository struct {
	db *gorm.DB
}
func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}
func (r *repository) CreateBooking(booking Booking) (Booking, error) {
	if err := r.db.Create(&booking).Error; err != nil {
		return Booking{}, err
	}
	return booking, nil
}
func (r *repository) GetBookingByID(id uint) (Booking, error) {
	var booking Booking
	if err := r.db.First(&booking, id).Error; err != nil {
		return Booking{}, err
	}
	return booking, nil
}
func (r *repository) GetAllBookings() ([]Booking, error) {
	var bookings []Booking
	if err := r.db.Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}
func (r *repository) UpdateBooking(id uint, booking Booking) (Booking, error) {
	if err := r.db.Model(&Booking{}).Where("id = ?", id).Updates(booking).Error; err != nil {
		return Booking{}, err
	}
	return booking, nil
}
func (r *repository) DeleteBooking(id uint) error {
	if err := r.db.Delete(&Booking{}, id).Error; err != nil {
		return err
	}
	return nil
}
func (r *repository) GetBookingsByUserID(userID uint) ([]Booking, error) {
	var bookings []Booking
	if err := r.db.Where("user_id = ?", userID).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}
func (r *repository) GetBookingsByEventID(eventID uint) ([]Booking, error) {
	var bookings []Booking
	if err := r.db.Where("event_id = ?", eventID).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}
func (r *repository) GetBookingsByDate(date time.Time) ([]Booking, error) {
	var bookings []Booking
	if err := r.db.Where("booking_date = ?", date).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}
func (r *repository) GetBookingsByStatus(status string) ([]Booking, error) {
	var bookings []Booking
	if err := r.db.Where("booking_status = ?", status).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}
func (r *repository) GetBookingsByPaymentStatus(status string) ([]Booking, error) {
	var bookings []Booking
	if err := r.db.Where("payment_status = ?", status).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}
func (r *repository) GetBookingsByDateRange(startDate, endDate time.Time) ([]Booking, error) {
	var bookings []Booking
	if err := r.db.Where("booking_date BETWEEN ? AND ?", startDate, endDate).Find(&bookings).Error; err != nil {
		return nil, err
	}
	return bookings, nil
}