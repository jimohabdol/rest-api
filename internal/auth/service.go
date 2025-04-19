package auth

import (
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/jimohabdol/rest-api/internal/user"
)

type Service interface {
	GenerateToken(user user.UserResponse) (accessToken, refreshToken string, err error)
	ValidateToken(token string) (uint, error)
	RefreshToken(token string) (newAccessToken, newRefreshToken string, err error)
}

type service struct {
	jwtSecretKey          string
	refreshTokenSecretKey string
	userRepo              user.Repository
}

type TokenClaims struct {
	UserID  uint   `json:"user_id"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

type RefeshTokenClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

func NewService(jwtSecretKey, refreshTokenSecretKey string, userRepo user.Repository) Service {
	return &service{
		jwtSecretKey:          jwtSecretKey,
		refreshTokenSecretKey: refreshTokenSecretKey,
		userRepo:              userRepo,
	}
}
func (s *service) GenerateToken(user user.UserResponse) (string, string, error) {
	// Generate JWT token using user information and secret key
	// Return the generated token
	accessTokenId, err := uuid.NewRandom()
	if err != nil {
		log.Printf("Error generating token ID: %v", err)
		return "", "", err
	}
	claims := TokenClaims{
		UserID:  user.ID,
		Email:   user.Email,
		IsAdmin: user.IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        accessTokenId.String(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
			Subject:   user.Email,
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err := accessToken.SignedString([]byte(s.jwtSecretKey))
	if err != nil {
		log.Printf("Error signing token: %v", err)
		return "", "", err
	}

	refreshTokenId, err := uuid.NewRandom()
	if err != nil {
		log.Printf("Error generating refresh token ID: %v", err)
		return "", "", err
	}
	refreshClaims := RefeshTokenClaims{
		UserID: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        refreshTokenId.String(),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(s.refreshTokenSecretKey))
	if err != nil {
		log.Printf("Error signing refresh token: %v", err)
		return "", "", err
	}
	return "Bearer " + accessTokenString, "Bearer " + refreshTokenString, nil
}
func (s *service) RefreshToken(token string) (string, string, error) {
	if !strings.HasPrefix(token, "Bearer ") {
		log.Printf("Invalid token format")
		return "", "", jwt.NewValidationError("invalid token format", jwt.ValidationErrorMalformed)
	}
	token = strings.TrimPrefix(token, "Bearer ")

	parsedToken, err := jwt.ParseWithClaims(
		token,
		&RefeshTokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
			}
			return []byte(s.refreshTokenSecretKey), nil
		},
	)
	if err != nil {
		log.Printf("Error parsing refresh token: %v", err)
		return "", "", err
	}
	if claims, ok := parsedToken.Claims.(*RefeshTokenClaims); ok && parsedToken.Valid {
		usr, err := s.userRepo.GetUserByID(claims.UserID)
		if err != nil {
			log.Printf("Error getting user by ID: %v", err)
			return "", "", err
		}
		return s.GenerateToken(user.ToUserResponse(usr))
	}
	log.Printf("Invalid token claims")
	return "", "", jwt.NewValidationError("invalid token claims", jwt.ValidationErrorClaimsInvalid)

}
func (s *service) ValidateToken(token string) (uint, error) {
	if !strings.HasPrefix(token, "Bearer ") {
		log.Printf("Invalid token format")
		return 0, jwt.NewValidationError("invalid token format", jwt.ValidationErrorMalformed)
	}
	token = strings.TrimPrefix(token, "Bearer ")

	parsedToken, err := jwt.ParseWithClaims(
		token,
		&TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
			}
			return []byte(s.jwtSecretKey), nil
		},
	)
	if err != nil {
		log.Printf("Error parsing token: %v", err)
		return 0, err
	}
	if claims, ok := parsedToken.Claims.(*TokenClaims); ok && parsedToken.Valid {
		return claims.UserID, nil
	}
	log.Printf("Invalid token claims")
	return 0, jwt.NewValidationError("invalid token claims", jwt.ValidationErrorClaimsInvalid)
}
