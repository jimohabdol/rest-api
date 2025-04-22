package router

import (
	"github.com/gin-gonic/gin"
	"github.com/jimohabdol/rest-api/internal/auth"
	"github.com/jimohabdol/rest-api/internal/booking"
	"github.com/jimohabdol/rest-api/internal/event"
	"github.com/jimohabdol/rest-api/internal/health"
	"github.com/jimohabdol/rest-api/internal/user"
)

func AuthRouter(r *gin.RouterGroup, h *auth.Handler) {
	AuthRoutes := r.Group("/auth")
	{
		AuthRoutes.POST("/auth/register", h.Register)
		AuthRoutes.POST("/auth/login", h.Login)
		AuthRoutes.POST("/auth/refresh", h.RefreshToken)
	}
}

func UserRouter(r *gin.RouterGroup, h *user.Handler, authMiddleWare *auth.Middleware) {
	userRoutes := r.Group("/users")
	userRoutes.Use(authMiddleWare.AuthRequired())
	{
		userRoutes.POST("", h.CreateUser)
		userRoutes.GET("/:id", h.GetUserByID)
		userRoutes.PUT("/:id", h.UpdateUser)
		userRoutes.DELETE("/:id", h.DeleteUser)
		userRoutes.GET("", h.GetAllUsers)
		userRoutes.GET("/email/:email", h.GetUserByEmail)
	}
}

func EvenRouter(r *gin.RouterGroup, h *event.Handler, authMiddleWare *auth.Middleware) {
	eventRoutes := r.Group("/envents")
	eventRoutes.Use(authMiddleWare.AuthRequired())
	{
		eventRoutes.POST("event", h.CreateEvent)
		eventRoutes.GET("events", h.GetAllEvents)
		eventRoutes.DELETE("event/:id", h.DeleteEvent)
		eventRoutes.GET("event/:id", h.GetEventByID)
		eventRoutes.PUT("event/:id", h.UpdateEvent)
		eventRoutes.GET("event/date/:date", h.GetEventByDate)
	}
}

func BookingRouter(r *gin.RouterGroup, h *booking.Handler, authMiddleWare *auth.Middleware) {
	bookingRoutes := r.Group("/bookings")
	bookingRoutes.Use(authMiddleWare.AuthRequired())
	{
		bookingRoutes.POST("booking", h.CreateBooking)
		bookingRoutes.GET("bookings", h.GetAllBookings)
		bookingRoutes.DELETE("booking/:id", h.DeleteBooking)
		bookingRoutes.GET("booking/:id", h.GetBookingByID)
		bookingRoutes.PUT("booking/:id", h.UpdateBooking)
		bookingRoutes.GET("bookings/user/:user_id", h.GetBookingsByUserID)
		bookingRoutes.GET("bookings/event/:event_id", h.GetBookingsByEventID)
		bookingRoutes.GET("bookings/date/:date", h.GetBookingsByDate)
		//bookingPath.GET("bookings/status/:status", h.GetBookingsByStatus)
	}
}

func HealthCheckerRouter(r *gin.RouterGroup, h *health.Handler) {
	healthRoutes := r.Group("/actuator")
	
	healthRoutes.GET("/health", h.GetHealthCheck)
	healthRoutes.GET("/info", h.GetInfo)
}
