package user

import (
	"github.com/jimohabdol/rest-api/internal/common"
	"gorm.io/gorm"

	"errors"
	"log"
)

type Repository interface {
	CreateUser(user User) (User, error)
	GetUserByID(id uint) (User, error)
	GetUserByEmail(email string) (User, error)
	UpdateUser(user User) (User, error)
	DeleteUser(id uint) error
	GetAllUsers() ([]User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(user User) (User, error) {
	hashPassword, err := common.HashPassword(user.Password)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return User{}, err
	}

	var existingUser User
	if err := r.db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		log.Printf("Email already exists: %s", user.Email)
		return User{}, common.ErrEmailAlreadyExists
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Printf("Database error while checking email: %v", err)
		return User{}, err
	}

	user.Password = string(hashPassword)
	if err := r.db.Create(&user).Error; err != nil {
		log.Printf("Failed to create user: %v", err)
		return User{}, err
	}
	return user, nil
}

func (r *repository) GetUserByID(id uint) (User, error) {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		log.Printf("Failed to get user by ID: %v", err)
		return User{}, err
	}
	return user, nil
}

func (r *repository) GetUserByEmail(email string) (User, error) {
	var user User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		log.Printf("Failed to get user by email: %v", err)
		return User{}, err
	}

	return user, nil
}

func (r *repository) UpdateUser(user User) (User, error) {
	if err := r.db.Save(&user).Error; err != nil {
		log.Printf("Failed to update user: %v", err)
		return User{}, err
	}
	return user, nil
}
func (r *repository) DeleteUser(id uint) error {
	if err := r.db.Delete(&User{}, id).Error; err != nil {
		log.Printf("Failed to delete user: %v", err)
		return err
	}
	return nil
}

func (r *repository) GetAllUsers() ([]User, error) {
	var users []User
	err := r.db.FindInBatches(&users, 100, func(tx *gorm.DB, batch int) error {
		return nil
	}).Error

	if err != nil {
		log.Printf("Failed to get all users: %v", err)
		return nil, err
	}

	return users, nil
}
