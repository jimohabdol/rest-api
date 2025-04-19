package booking

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}
func (h *Handler) CreateBooking(c *gin.Context) {
	var booking CreateBookingRequest
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"code":    -1,
			"error":   err.Error(),
		})
		return
	}
	createdBooking, err := h.service.CreateBooking(booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create booking",
			"code":    -1,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Booking created successfully",
		"code":    0,
		"booking": createdBooking,
	})
}

func (h *Handler) GetBookingByID(c *gin.Context) {
	id := c.GetUint("id")
	booking, err := h.service.GetBookingByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Booking not found",
			"code":    -1,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"code":    0,
		"booking": booking,
	})
}
func (h *Handler) GetAllBookings(c *gin.Context) {
	bookings, err := h.service.GetAllBookings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve bookings",
			"code":    -1,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "success",
		"code":     0,
		"bookings": bookings,
	})
}
func (h *Handler) UpdateBooking(c *gin.Context) {
	id := c.GetUint("id")
	var booking UpdateBookingRequest
	if err := c.ShouldBindJSON(&booking); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"code":    -1,
			"error":   err.Error(),
		})
		return
	}
	updatedBooking, err := h.service.UpdateBooking(id, booking)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update booking",
			"code":    -1,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Booking updated successfully",
		"code":    0,
		"booking": updatedBooking,
	})
}
func (h *Handler) DeleteBooking(c *gin.Context) {
	id := c.GetUint("id")
	err := h.service.DeleteBooking(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete booking",
			"code":    -1,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Booking deleted successfully",
		"code":    0,
	})
}
func (h *Handler) GetBookingsByUserID(c *gin.Context) {
	userID := c.GetUint("user_id")
	bookings, err := h.service.GetBookingsByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve bookings",
			"code":    -1,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "success",
		"code":     0,
		"bookings": bookings,
	})
}
func (h *Handler) GetBookingsByEventID(c *gin.Context) {
	eventID := c.GetUint("event_id")
	bookings, err := h.service.GetBookingsByEventID(eventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve bookings",
			"code":    -1,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "success",
		"code":     0,
		"bookings": bookings,
	})
}
func (h *Handler) GetBookingsByDate(c *gin.Context) {
	date := c.Query("date")
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid date format",
			"code":    -1,
			"error":   err.Error(),
		})
		return
	}
	bookings, err := h.service.GetBookingsByDate(parsedDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to retrieve bookings",
			"code":    -1,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":  "success",
		"code":     0,
		"bookings": bookings,
	})
}
