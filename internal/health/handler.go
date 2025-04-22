package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}	

func (h *Handler) GetHealthCheck(c *gin.Context) {
	health := h.service.GetHealthCheck()
	statusCode := http.StatusOK
	if health.Status == StatusDown {
		statusCode = http.StatusServiceUnavailable
	}
	c.JSON(statusCode, health)
}

func (h *Handler) GetInfo(c *gin.Context) {
	info := h.service.GetInfo()
	c.JSON(http.StatusOK, info)
}