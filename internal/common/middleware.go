package common

import (
	"fmt"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func LatencyLogMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Generate UUID and set start time
		uuid := uuid.New()
		startTime := time.Now()
		c.Set("uuid", uuid)
		
		// Log request start with method and path
		fmt.Printf("[%s] %s %s started at %v\n", 
			uuid, 
			c.Request.Method, 
			c.Request.URL.Path, 
			startTime.Format(time.RFC3339))
		
		c.Next()
		
		// Calculate latency and log completion
		latency := time.Since(startTime).Milliseconds()
		fmt.Printf("[%s] %s %s completed in %dms | Status: %d | Completed at %v\n", 
			uuid,
			c.Request.Method,
			c.Request.URL.Path,
			latency,
			c.Writer.Status(),
			time.Now().Format(time.RFC3339))
			
	}
}

func GetHttpClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	return &http.Client{
		Transport: tr,
	}
}