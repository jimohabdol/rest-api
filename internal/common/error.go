package common

import (
	"errors"
	"github.com/gin-gonic/gin"
	// "net/http"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")
	ErrEventNotFound      = errors.New("event not found")
	ErrBookingNotFound    = errors.New("booking not found")
	ErrUnauthorized       = errors.New("unauthorized")
	ErrEventFull          = errors.New("event is at full capacity")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidToken       = errors.New("invalid token")
	ErrTokenExpired       = errors.New("token expired")
	ErrInternalServer     = errors.New("internal server error")
	ErrBadRequest         = errors.New("bad request")
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func RespondWithError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, ErrorResponse{Error: err.Error()})
}
