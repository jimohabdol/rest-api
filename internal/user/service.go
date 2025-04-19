package user

import (
	"errors"
	"github.com/jimohabdol/rest-api/internal/common"
)

type Service interface {
	CreateUser(user CreateUserRequest) (UserResponse, error)
	GetUserByID(id uint) (UserResponse, error)
	GetUserByEmail(email string) (UserResponse, error)
	UpdateUser(id uint, user UpdateUserRequest) (UserResponse, error)
	DeleteUser(id uint) error
	GetAllUsers() ([]UserResponse, error)
	ValidateUser(email, password string) (UserResponse, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) CreateUser(user CreateUserRequest) (UserResponse, error) {
	newUser := User{
		Email:     user.Email,
		Password:  user.Password,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		IsAdmin:   user.IsAdmin,
	}

	createdUser, err := s.repo.CreateUser(newUser)
	if err != nil {
		if err.Error() == "email already exists" {
			return UserResponse{}, errors.New("email already exists")
		}
		return UserResponse{}, err
	}

	return ToUserResponse(createdUser), nil
}
func (s *service) GetUserByID(id uint) (UserResponse, error) {
	user, err := s.repo.GetUserByID(id)
	if err != nil {
		return UserResponse{}, err
	}

	return ToUserResponse(user), nil
}
func (s *service) GetUserByEmail(email string) (UserResponse, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return UserResponse{}, err
	}

	return ToUserResponse(user), nil
}
func (s *service) UpdateUser(id uint, user UpdateUserRequest) (UserResponse, error) {
	existingUser, err := s.repo.GetUserByID(id)
	if err != nil {
		return UserResponse{}, err
	}

	if user.Email != "" {
		existingUser.Email = user.Email
	}
	if user.Password != "" {
		existingUser.Password = user.Password
	}

	updatedUser, err := s.repo.UpdateUser(existingUser)
	if err != nil {
		return UserResponse{}, err
	}

	return ToUserResponse(updatedUser), nil
}
func (s *service) DeleteUser(id uint) error {
	err := s.repo.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}
func (s *service) GetAllUsers() ([]UserResponse, error) {
	users, err := s.repo.GetAllUsers()
	if err != nil {
		return nil, err
	}

	var userResponses []UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses, nil
}

func (s *service) ValidateUser(email, password string) (UserResponse, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return UserResponse{}, err
	}

	if common.CheckPasswordHash(password, user.Password) != nil {
		return UserResponse{}, err
	}

	return ToUserResponse(user), nil
}

func ToUserResponse(user User) UserResponse {
	return UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
		IsAdmin:   user.IsAdmin,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}
