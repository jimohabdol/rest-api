package booking

import (
	"time"
)

type Service interface {
	CreateBooking(booking CreateBookingRequest) (BookingResponse, error)
	GetBookingByID(id uint) (BookingResponse, error)
	GetAllBookings() ([]BookingResponse, error)
	UpdateBooking(id uint, booking UpdateBookingRequest) (BookingResponse, error)
	DeleteBooking(id uint) error
	GetBookingsByUserID(userID uint) ([]BookingResponse, error)
	GetBookingsByEventID(eventID uint) ([]BookingResponse, error)
	GetBookingsByDate(date time.Time) ([]BookingResponse, error)
	GetBookingsByStatus(status string) ([]BookingResponse, error)
}
type service struct {
	repo Repository
}