package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimohabdol/rest-api/internal/user"
)

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthResponse struct {
	Token string            `json:"token"`
	User  user.UserResponse `json:"user"`
}

type Handler struct {
	userService user.Service
	authService Service
}

func NewHandler(userService user.Service, authService Service) *Handler {
	return &Handler{
		userService: userService,
		authService: authService,
	}
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"code":    -1,
		})
		return
	}
	user, err := h.userService.ValidateUser(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid email or password",
			"code":    -1,
		})
		return
	}
	token, refreshToken, err := h.authService.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate token",
			"code":    -1,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message":       "Login successful",
		"code":          0,
		"token":         token,
		"refresh_token": refreshToken,
		"user":          user,
	})
}

func (h *Handler) Register(c *gin.Context) {
	var userReq user.CreateUserRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		log.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"code":    -1,
		})
		return
	}
	newUser, err := h.userService.CreateUser(userReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create user",
			"code":    -1,
			"error":   err.Error(),
		})
		return
	}
	token, refreshToken, err := h.authService.GenerateToken(newUser)
	if err != nil {
		log.Println("Error generating token:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to generate token",
			"code":    -1,
		})
		return
	}
	c.Header("Authorization", token)
	c.JSON(http.StatusCreated, gin.H{
		"message":       "User created successfully",
		"code":          0,
		"user":          newUser,
		"refresh_token": refreshToken,
		"token":         token,
	})
}

func (h *Handler) RefreshToken(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Missing or invalid token",
			"code":    -1,
		})
		return
	}
	newToken, refreshToken, err := h.authService.RefreshToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Invalid token",
			"code":    -1,
		})
		return
	}
	c.Header("Authorization", newToken)
	c.JSON(http.StatusOK, gin.H{
		"message":       "Token refreshed successfully",
		"code":          0,
		"token":         newToken,
		"refresh_token": refreshToken,
	})
}
