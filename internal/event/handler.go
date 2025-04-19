package event

import (
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}
func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}
func (h *Handler) CreateEvent(c *gin.Context) {
	var event CreateEventRequest
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"code":    -1,
		})
		return
	}
	createdEvent, err := h.service.CreateEvent(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create event",
			"code":    -1,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Event created successfully",
		"code":    0,
		"event":   createdEvent,
	})
}
func (h *Handler) GetEventByID(c *gin.Context) {	
	id := c.GetUint("id")
	event, err := h.service.GetEventByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
			"code":    -1,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"code":    0,
		"event":   event,
	})
}
func (h *Handler) GetAllEvents(c *gin.Context) {
	events, err := h.service.GetAllEvents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve events",
			"code":    -1,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"code":    0,
		"events":  events,
	})
}
func (h *Handler) UpdateEvent(c *gin.Context) {
	id := c.GetUint("id")
	var event UpdateEventRequest
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"code":    -1,
		})
		return
	}
	updatedEvent, err := h.service.UpdateEvent(id, event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update event",
			"code":    -1,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Event updated successfully",
		"code":    0,
		"event":   updatedEvent,
	})
}
func (h *Handler) DeleteEvent(c *gin.Context) {
	id := c.GetUint("id")
	err := h.service.DeleteEvent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete event",
			"code":    -1,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Event deleted successfully",
		"code":    0,
	})
}
func (h *Handler) GetEventByDate(c *gin.Context) {
	dateStr := c.Query("date")
	if dateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Date query parameter is required",
			"code":    -1,
		})
		return
	}
	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid date format",
			"code":    -1,
		})
		return
	}
	event, err := h.service.GetEventsByDate(date)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Event not found",
			"code":    -1,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"code":    0,
		"event":   event,
	})
}