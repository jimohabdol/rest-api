package health

import "time"

type HealthStatus struct {
	Status      string    `json:"status"`
	Checks map[string]Check `json:"checks"`
	Info map[string]interface{} `json:"info"`
	Timestamp time.Time `json:"timestamp"`
}
	
type Check struct {
	Message    string `json:"Message"`
	Status  string `json:"status"`
	Error string `json:"error"`
}

const (
	StatusUp   = "UP"
	StatusDown = "DOWN"
	StatusUnknown = "UNKNOWN"
	StatusDegraded = "DEGRADED"
)