package server

type HealthReport struct {
	Database bool   `json:"database"`
	Version  string `json:"version"`
}
