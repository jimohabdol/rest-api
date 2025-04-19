package booking

import (
	"github.com/jimohabdol/rest-api/internal/event"
	"github.com/jimohabdol/rest-api/internal/user"
	"gorm.io/gorm"
	"time"
)

type Booking struct {
	gorm.Model
	EventID       uint        `json:"event_id" gorm:"not null"`
	Event         event.Event `json:"event" gorm:"foreignKey:EventID"`
	UserID        uint        `json:"user_id" gorm:"not null"`
	User          user.User   `json:"user" gorm:"foreignKey:UserID"`
	BookingDate   time.Time   `json:"booking_date" gorm:"not null"`
	BookingStatus string      `json:"booking_status" gorm:"not null"`
	PaymentStatus string      `json:"payment_status" gorm:"not null"`
	Tickets       int         `json:"tickets" gorm:"not null;check: tickets > 0"`
}

type BookingResponse struct {
	ID            uint                `json:"id"`
	EventID       uint                `json:"event_id"`
	Event         event.EventResponse `json:"event"`
	UserID        uint                `json:"user_id"`
	User          user.UserResponse   `json:"user"`
	BookingDate   time.Time           `json:"booking_date"`
	BookingStatus string              `json:"booking_status"`
	PaymentStatus string              `json:"payment_status"`
	Tickets       int                 `json:"tickets"`
	CreatedAt     time.Time           `json:"created_at"`
	UpdatedAt     time.Time           `json:"updated_at"`
	DeletedAt     gorm.DeletedAt      `json:"deleted_at" gorm:"index"`
}
type CreateBookingRequest struct {
	EventID       uint      `json:"event_id" binding:"required"`
	UserID        uint      `json:"user_id" binding:"required"`
	BookingDate   time.Time `json:"booking_date" binding:"required"`
	BookingStatus string    `json:"booking_status" binding:"required"`
	PaymentStatus string    `json:"payment_status" binding:"required"`
	Tickets       int       `json:"tickets" binding:"required"`
}
type UpdateBookingRequest struct {
	EventID       uint      `json:"event_id"`
	UserID        uint      `json:"user_id"`
	BookingDate   time.Time `json:"booking_date"`
	BookingStatus string    `json:"booking_status"`
	PaymentStatus string    `json:"payment_status"`
	Tickets       int       `json:"tickets"`
}
