package user

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var user CreateUserRequest
	log.Println(c.GetString("is_admin"))
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"code":    -1,
			"error":   err.Error(),
		})
		return
	}
	createdUser, err := h.service.CreateUser(user)
	if err != nil {
		log.Println("Error creating user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create user",
			"code":    -1,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"code":    0,
		"user":    createdUser,
	})
}

func (h *Handler) GetUserByID(c *gin.Context) {
	id := c.GetUint("id")
	user, err := h.service.GetUserByID(id)
	if err != nil {
		c.JSON((http.StatusNotFound), gin.H{
			"message": `User with ID ${id} not found`,
			"code":    -1,
			"error":   err.Error(),
		})
		return
	}
	c.JSON((http.StatusFound), gin.H{
		"message": "success",
		"code":    0,
		"user":    user,
	})
}

func (h *Handler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := h.service.GetUserByEmail(email)
	if err != nil {
		c.JSON((http.StatusNotFound), gin.H{
			"message": `User with email ${email} not found`,
			"code":    -1,
			"error":   err.Error(),
		})
		return
	}
	c.JSON((http.StatusFound), gin.H{
		"message": "success",
		"code":    0,
		"user":    user,
	})
}

func (h *Handler) UpdateUser(c *gin.Context) {
	id := c.GetUint("id")
	var user UpdateUserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request",
			"code":    -1,
			"error":   err.Error(),
		})
		return
	}
	updatedUser, err := h.service.UpdateUser(id, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to update user",
			"code":    -1,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"code":    0,
		"user":    updatedUser,
	})
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get users",
			"code":    -1,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"code":    0,
		"users":   users,
	})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.GetUint("id")
	err := h.service.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to delete user",
			"code":    -1,
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
		"code":    0,
	})
}
