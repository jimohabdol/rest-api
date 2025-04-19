package auth

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	authService Service
}

func NewMiddleware(authService Service) *Middleware {
	return &Middleware{
		authService: authService,
	}
}
func (m *Middleware) AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			log.Println("Missing token")
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Missing or invalid token",
				"code":    -1,
			})
			c.Abort()
			return
		}

		userID, err := m.authService.ValidateToken(token)
		if err != nil {
			log.Println("Invalid token:", err)
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Invalid token",
				"code":    -1,
			})
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Next()
	}
}
