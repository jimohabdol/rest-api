package booking

import (
	"errors"
	"strings"
	"time"

	"github.com/jimohabdol/rest-api/internal/common"
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

func NewService(repo Repository) Service {
	return &service{repo: repo}
}
func (s *service) CreateBooking(booking CreateBookingRequest) (BookingResponse, error) {
	newBooking := Booking{
		EventID:       booking.EventID,
		UserID:        booking.UserID,
		BookingDate:   booking.BookingDate,
		BookingStatus: booking.BookingStatus,
		PaymentStatus: booking.PaymentStatus,
		Tickets:       booking.Tickets,
	}

	if booking.Tickets <= 0 {
		return BookingResponse{}, errors.New("tickets must be greater than zero")
	}

	if booking.BookingDate.IsZero() || booking.BookingDate.Before(time.Now()) {
		return BookingResponse{}, errors.New("booking date must be in the future")
	}

	if strings.TrimSpace(booking.BookingStatus) == "" {
		return BookingResponse{}, errors.New("booking status cannot be empty")
	} else if common.IsValidStatus(booking.BookingStatus) == false {
		return BookingResponse{}, errors.New("invalid booking status")
	}

	createdBooking, err := s.repo.CreateBooking(newBooking)
	if err != nil {
		return BookingResponse{}, err
	}

	return ToBookingResponse(createdBooking), nil
}

func (s *service) GetBookingByID(id uint) (BookingResponse, error) {
	booking, err := s.repo.GetBookingByID(id)
	if err != nil {
		return BookingResponse{}, err
	}

	return ToBookingResponse(booking), nil
}
func (s *service) GetAllBookings() ([]BookingResponse, error) {
	bookings, err := s.repo.GetAllBookings()
	if err != nil {
		return nil, err
	}

	var bookingResponses []BookingResponse
	for _, booking := range bookings {
		bookingResponses = append(bookingResponses, ToBookingResponse(booking))
	}

	return bookingResponses, nil
}
func (s *service) UpdateBooking(id uint, booking UpdateBookingRequest) (BookingResponse, error) {
	existingBooking, err := s.repo.GetBookingByID(id)
	if err != nil {
		return BookingResponse{}, err
	}

	if booking.EventID != 0 {
		existingBooking.EventID = booking.EventID
	}
	if booking.UserID != 0 {
		existingBooking.UserID = booking.UserID
	}
	if booking.BookingDate.IsZero() == false {
		existingBooking.BookingDate = booking.BookingDate
	}
	if strings.TrimSpace(booking.BookingStatus) != "" {
		existingBooking.BookingStatus = booking.BookingStatus
	}
	if strings.TrimSpace(booking.PaymentStatus) != "" {
		existingBooking.PaymentStatus = booking.PaymentStatus
	}
	if booking.Tickets > 0 {
		existingBooking.Tickets = booking.Tickets
	}

	updatedBooking, err := s.repo.UpdateBooking(id, existingBooking)
	if err != nil {
		return BookingResponse{}, err
	}

	return ToBookingResponse(updatedBooking), nil
}
func (s *service) DeleteBooking(id uint) error {
	err := s.repo.DeleteBooking(id)
	if err != nil {
		return err
	}

	return nil
}
func (s *service) GetBookingsByUserID(userID uint) ([]BookingResponse, error) {
	bookings, err := s.repo.GetBookingsByUserID(userID)
	if err != nil {
		return nil, err
	}

	var bookingResponses []BookingResponse
	for _, booking := range bookings {
		bookingResponses = append(bookingResponses, ToBookingResponse(booking))
	}

	return bookingResponses, nil
}
func (s *service) GetBookingsByEventID(eventID uint) ([]BookingResponse, error) {
	bookings, err := s.repo.GetBookingsByEventID(eventID)
	if err != nil {
		return nil, err
	}

	var bookingResponses []BookingResponse
	for _, booking := range bookings {
		bookingResponses = append(bookingResponses, ToBookingResponse(booking))
	}

	return bookingResponses, nil
}
func (s *service) GetBookingsByDate(date time.Time) ([]BookingResponse, error) {
	bookings, err := s.repo.GetBookingsByDate(date)
	if err != nil {
		return nil, err
	}

	var bookingResponses []BookingResponse
	for _, booking := range bookings {
		bookingResponses = append(bookingResponses, ToBookingResponse(booking))
	}

	return bookingResponses, nil
}
func (s *service) GetBookingsByStatus(status string) ([]BookingResponse, error) {
	bookings, err := s.repo.GetBookingsByStatus(status)
	if err != nil {
		return nil, err
	}

	var bookingResponses []BookingResponse
	for _, booking := range bookings {
		bookingResponses = append(bookingResponses, ToBookingResponse(booking))
	}

	return bookingResponses, nil
}
