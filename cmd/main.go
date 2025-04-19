package main

import (
	"log"
	"os"
	// "net/http"
	"github.com/gin-gonic/gin"
	"github.com/jimohabdol/rest-api/internal/auth"
	"github.com/jimohabdol/rest-api/internal/common"
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
	}

	server.Run(":8080")
}

// postgresql://test_mg7t_user:7Nq5gF5hOXOrflibp8ugGVs1KDgRnaWa@dpg-cvvfforuibrs73bdogkg-a.oregon-postgres.render.com/test_mg7t
