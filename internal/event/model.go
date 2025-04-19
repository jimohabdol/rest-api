package event

import (
	"github.com/jimohabdol/rest-api/internal/user"
	"gorm.io/gorm"
	"time"
)

type Event struct {
	gorm.Model
	Title       string    `json:"title" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	Location    string    `json:"location" gorm:"not null"`
	StartTime   time.Time `json:"start_time" gorm:"not null"`
	EndTime     time.Time `json:"end_time" gorm:"not null"`
	Capacity    int       `json:"capacity" gorm:"not null"`
	UserID      uint      `json:"user_id" gorm:"not null"`
	User        user.User `json:"user" gorm:"foreignKey:UserID"`
}

type EventResponse struct {
	ID          uint              `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Location    string            `json:"location"`
	StartTime   time.Time         `json:"start_time"`
	EndTime     time.Time         `json:"end_time"`
	Capacity    int               `json:"capacity"`
	User        user.UserResponse `json:"user"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

type CreateEventRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	StartTime   time.Time `json:"start_time" binding:"required"`
	EndTime     time.Time `json:"end_time" binding:"required"`
	Capacity    int       `json:"capacity" binding:"required"`
	UserID      uint      `json:"user_id" binding:"required"`
}
type UpdateEventRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Capacity    int       `json:"capacity"`
	UserID      uint      `json:"user_id"`
}

func ToEventResponse(event Event) EventResponse {
	return EventResponse{
		ID:          event.ID,
		Title:       event.Title,
		Description: event.Description,
		Location:    event.Location,
		StartTime:   event.StartTime,
		EndTime:     event.EndTime,
		Capacity:    event.Capacity,
		User:        user.ToUserResponse(event.User),
		CreatedAt:   event.CreatedAt,
		UpdatedAt:   event.UpdatedAt,
	}
}
func ToEventResponses(events []Event) []EventResponse {
	var eventResponses []EventResponse
	for _, event := range events {
		eventResponses = append(eventResponses, ToEventResponse(event))
	}
	return eventResponses
}
