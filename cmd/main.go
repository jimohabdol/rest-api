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
	"github.com/jimohabdol/rest-api/internal/health"
	"github.com/jimohabdol/rest-api/internal/router"
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

	// if err := db.AutoMigrate(
	// 	&user.User{},
	// 	&event.Event{},
	// 	&booking.Booking{},
	// ); err != nil {
	// 	log.Fatalf("Failed to run migrations: %v", err)
	// }

	log.Println("Database and environment successfully initialized!")
}

func main() {
	log.Println(db)
	healthRepo := health.NewRepository(db)
	healthService := health.NewService(healthRepo)
	healthHandler := health.NewHandler(healthService)
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	eventRepo := event.NewRepository(db)
	eventService := event.NewService(eventRepo)
	eventHandler := event.NewHandler(eventService)
	bookingRepo := booking.NewRepository(db)
	bookingService := booking.NewService(bookingRepo)
	bookingHandler := booking.NewHandler(bookingService)
	authService := auth.NewService(
		os.Getenv("JWT_ACCESS_SECRET"),
		os.Getenv("JWT_REFESH_SECRET"),
		userRepo,
	)
	authMiddleware := auth.NewMiddleware(authService)
	authHandler := auth.NewHandler(userService, authService)

	server := gin.Default()
	server.Use(common.LatencyLogMiddleWare())
	contextPath := server.Group("/api/v1")

	// Health routes
	router.HealthCheckerRouter(contextPath, healthHandler)

	// Auth routes
	router.AuthRouter(contextPath, authHandler)
	// User routes
	router.UserRouter(contextPath, userHandler, authMiddleware)
	// Event routes
	router.EvenRouter(contextPath, eventHandler, authMiddleware)
	// Booking routes
	router.BookingRouter(contextPath, bookingHandler, authMiddleware)

	server.Run(":8080")
}
