package common

import (
	"golang.org/x/crypto/bcrypt"
	"net/mail"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func IsValidStatus(status string) bool {
	validStatus := map[string]bool{
		"pending":  true,
		"confirmed": true,
		"cancelled": true,
	}
	return validStatus[status]
}

func IsValidEmail(email string) bool {
    _, err := mail.ParseAddress(email)
    return err == nil // If there's no error, email is valid
}

func IsFIeldEmpty(field string) bool {
    return len(field) == 0
}