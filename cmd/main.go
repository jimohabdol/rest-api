package main

import (
	"log"
	// "net/http"
	"github.com/gin-gonic/gin"
	"github.com/jimohabdol/rest-api/internal/auth"
	"github.com/jimohabdol/rest-api/internal/common"
	"github.com/jimohabdol/rest-api/internal/user"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	// Initialize any necessary configurations or dependencies here
	// For example, you might want to set up a database connection or load environment variables.
	var err error
	db, err = common.InitDB()
	if err != nil {
		log.Fatal("failed to connect database")
	}
	db.AutoMigrate(&user.User{})
	// , &event.Event{}, &booking.Booking{})
}

func main() {
	log.Println(db)
	userRepo := user.NewRepository(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)
	authService := auth.NewService("test", "test", userRepo)
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
	}

	server.Run(":8080")
}

// postgresql://test_mg7t_user:7Nq5gF5hOXOrflibp8ugGVs1KDgRnaWa@dpg-cvvfforuibrs73bdogkg-a.oregon-postgres.render.com/test_mg7t
