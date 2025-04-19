package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `json:"email" gorm:"unique; not null"`
	Password  string `json:"password" gorm:"not null"`
	FirstName string `json:"name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`
	IsAdmin   bool   `json:"is_admin" gorm:"not null"`
}
type UserResponse struct {
	ID        uint           `json:"id"`
	Email     string         `json:"email"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	CreatedAt time.Time      `json:"created_at"`
	IsAdmin   bool           `json:"is_admin"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
type CreateUserRequest struct {
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	IsAdmin   bool   `json:"is_admin" binding:"required"`
}
type UpdateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// func ToUserResponse(user User) UserResponse {
// 	return UserResponse{
// 		ID: user.ID,
// 		Email: user.Email,
// 		FirstName: user.FirstName,
// 		LastName: user.LastName,
// 		CreatedAt: user.CreatedAt,
// 		IsAdmin: user.IsAdmin,
// 		UpdatedAt: user.UpdatedAt,
// 		DeletedAt: user.DeletedAt,
// 	}
// }
