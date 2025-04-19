package main

import (
	"log"
	"os"

	// "net/http"
	"github.com/gin-gonic/gin"
	"github.com/jimohabdol/rest-api/internal/auth"
	"github.com/jimohabdol/rest-api/internal/booking"
	"github.com/jimohabdol/rest-api/internal/common"
	"github.com/jimohabdol/rest-api/internal/event"
	"github.com/jimohabdol/rest-api/internal/user"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	// Load environment variables first
	if err := common.LoadEnv(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

	var err error
	db, err = common.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	if err := db.AutoMigrate(&user.User{}); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database and environment successfully initialized!")
}

func main() {
	log.Println(db)
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	eventRepo := event.NewRepository(db)
	eventService := event.NewService(eventRepo)
	eventHandler := event.NewHandler(eventService)
	bookingRepo := booking.NewRepository(db)
	bookingService := booking.NewService(bookingRepo)
	bookingHandler := booking.NewHandler(bookingService)
	authService := auth.NewService(os.Getenv("JWT_ACCESS_SECRET"), os.Getenv("JWT_REFESH_SECRET"), userRepo)
	authMiddleware := auth.NewMiddleware(authService)
	authHandler := auth.NewHandler(userService, authService)


	server := gin.Default()
	// server.Use(gin.Logger())
	contextPath := server.Group("/api/v1")

	// Auth routes
	contextPath.POST("/auth/register", authHandler.Register)
	contextPath.POST("/auth/login", authHandler.Login)
	contextPath.POST("/auth/refresh", authHandler.RefreshToken)

	// User routes
	userPath := contextPath.Group("/")
	userPath.Use(authMiddleware.AuthRequired())
	{
		userPath.POST("user", userHandler.CreateUser)
		userPath.GET("users", userHandler.GetAllUsers)
		userPath.DELETE("user/:id", userHandler.DeleteUser)
		userPath.GET("user/:id", userHandler.GetUserByID)
		userPath.PUT("user/:id", userHandler.UpdateUser)
		userPath.GET("user/email/:email", userHandler.GetUserByEmail)
	}

	// Event routes
	eventPath := contextPath.Group("/")
	eventPath.Use(authMiddleware.AuthRequired())
	{
		eventPath.POST("event", eventHandler.CreateEvent)
		eventPath.GET("events", eventHandler.GetAllEvents)
		eventPath.DELETE("event/:id", eventHandler.DeleteEvent)
		eventPath.GET("event/:id", eventHandler.GetEventByID)
		eventPath.PUT("event/:id", eventHandler.UpdateEvent)
		eventPath.GET("event/date/:date", eventHandler.GetEventByDate)
	}

	// Booking routes
	bookingPath := contextPath.Group("/")
	bookingPath.Use(authMiddleware.AuthRequired())
	{
		bookingPath.POST("booking", bookingHandler.CreateBooking)
		bookingPath.GET("bookings", bookingHandler.GetAllBookings)
		bookingPath.DELETE("booking/:id", bookingHandler.DeleteBooking)
		bookingPath.GET("booking/:id", bookingHandler.GetBookingByID)
		bookingPath.PUT("booking/:id", bookingHandler.UpdateBooking)
		bookingPath.GET("bookings/user/:user_id", bookingHandler.GetBookingsByUserID)
		bookingPath.GET("bookings/event/:event_id", bookingHandler.GetBookingsByEventID)
		bookingPath.GET("bookings/date/:date", bookingHandler.GetBookingsByDate)
		//bookingPath.GET("bookings/status/:status", bookingHandler.GetBookingsByStatus)
	}

	server.Run(":8080")
}