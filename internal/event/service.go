package event

import "time"

type Service interface {
	CreateEvent(event CreateEventRequest) (EventResponse, error)
	GetEventByID(id uint) (EventResponse, error)
	GetAllEvents() ([]EventResponse, error)
	UpdateEvent(id uint, event UpdateEventRequest) (EventResponse, error)
	DeleteEvent(id uint) error
	GetEventsByUserID(userID uint) ([]EventResponse, error)
	GetEventsByDate(date time.Time) ([]EventResponse, error)
	GetEventsByLocation(location string) ([]EventResponse, error)
	ValidateEvent(eventID uint) (EventResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}
func (s *service) ValidateEvent(eventID uint) (EventResponse, error) {
	event, err := s.repo.GetEventByID(eventID)
	if err != nil {
		return EventResponse{}, err
	}
	if event.ID == 0 {
		return EventResponse{}, err
	}
	return ToEventResponse(event), nil
}
func (s *service) CreateEvent(event CreateEventRequest) (EventResponse, error) {
	newEvent := Event{
		Title:       event.Title,
		Description: event.Description,
		Location:    event.Location,
		StartTime:   event.StartTime,
		EndTime:     event.EndTime,
		Capacity:    event.Capacity,
		UserID:      event.UserID,
	}

	createdEvent, err := s.repo.CreateEvent(newEvent)
	if err != nil {
		return EventResponse{}, err
	}

	return ToEventResponse(createdEvent), nil
}
func (s *service) GetEventByID(id uint) (EventResponse, error) {
	event, err := s.repo.GetEventByID(id)
	if err != nil {
		return EventResponse{}, err
	}

	return ToEventResponse(event), nil
}
func (s *service) GetAllEvents() ([]EventResponse, error) {
	events, err := s.repo.GetAllEvents()
	if err != nil {
		return nil, err
	}

	var eventResponses []EventResponse
	for _, event := range events {
		eventResponses = append(eventResponses, ToEventResponse(event))
	}

	return eventResponses, nil
}
func (s *service) UpdateEvent(id uint, event UpdateEventRequest) (EventResponse, error) {
	existingEvent, err := s.repo.GetEventByID(id)
	if err != nil {
		return EventResponse{}, err
	}

	if event.Title != "" {
		existingEvent.Title = event.Title
	}
	if event.Description != "" {
		existingEvent.Description = event.Description
	}
	if event.Location != "" {
		existingEvent.Location = event.Location
	}
	if !event.StartTime.IsZero() {
		existingEvent.StartTime = event.StartTime
	}
	if !event.EndTime.IsZero() {
		existingEvent.EndTime = event.EndTime
	}
	if event.Capacity != 0 {
		existingEvent.Capacity = event.Capacity
	}

	updatedEvent, err := s.repo.UpdateEvent(id, existingEvent)
	if err != nil {
		return EventResponse{}, err
	}

	return ToEventResponse(updatedEvent), nil
}
func (s *service) DeleteEvent(id uint) error {
	err := s.repo.DeleteEvent(id)
	if err != nil {
		return err
	}

	return nil
}
func (s *service) GetEventsByUserID(userID uint) ([]EventResponse, error) {
	events, err := s.repo.GetEventsByUserID(userID)
	if err != nil {
		return nil, err
	}

	var eventResponses []EventResponse
	for _, event := range events {
		eventResponses = append(eventResponses, ToEventResponse(event))
	}

	return eventResponses, nil
}
func (s *service) GetEventsByDate(date time.Time) ([]EventResponse, error) {
	events, err := s.repo.GetEventsByDate(date)
	if err != nil {
		return nil, err
	}

	var eventResponses []EventResponse
	for _, event := range events {
		eventResponses = append(eventResponses, ToEventResponse(event))
	}

	return eventResponses, nil
}
func (s *service) GetEventsByLocation(location string) ([]EventResponse, error) {
	events, err := s.repo.GetEventsByLocation(location)
	if err != nil {
		return nil, err
	}

	var eventResponses []EventResponse
	for _, event := range events {
		eventResponses = append(eventResponses, ToEventResponse(event))
	}

	return eventResponses, nil
}
