package health

import "time"

type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Uptime    time.Time `json:"uptime"`
	Memory    string    `json:"memory"`
	DBStatus  string    `json:"db_status"`
}
