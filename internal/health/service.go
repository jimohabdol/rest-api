package health

import (
	"runtime"
	"time"
)

type Service interface {
	GetHealthCheck() HealthStatus
	GetInfo() map[string]interface{}
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) GetHealthCheck() HealthStatus {
	err := s.repo.GetHealthCheck()
	healthStatus := HealthStatus{
		Status:  func() string {
			if err != nil {
				return StatusDegraded
			}
			return StatusUp
		}(),
		Checks:  make(map[string]Check),
		Info:    make(map[string]interface{}),
	}

	if err == nil {
		healthStatus.Checks["database"] = Check{
			Message: "Database connection is healthy",
			Status:  StatusUp,
		}
	} else {
		healthStatus.Checks["database"] = Check{
			Message: "Database connection is unhealthy",
			Status:  StatusDown,
			Error:   err.Error(),
		}
	}
	healthStatus.Info = s.GetInfo()

	return healthStatus
}

func (s *service) GetInfo() map[string]interface{} {
	info := make(map[string]interface{})
	info["service"] = map[string]any{
		"name":    "Event Booking Service",
		"version": "1.0.0",
		"description": "This is a API for managing events and booking.",
		"author":  "Abdulrahman Jimoh",
	}
	info["metrics"] = map[string]any{
		"uptime":     time.Since(time.Now()).String(),
		"timestamp":  time.Now().Format(time.RFC3339),
		"memory_usage": runtime.MemStats{},
		"cpu_usage":  runtime.NumCPU(),
	}
	return info
}