package server

import (
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type DatabaseInfo struct {
	Connected       bool    `json:"connected"`
	LatencyMs       float64 `json:"latency_ms"`
	PostgresVersion string  `json:"postgres_version,omitempty"`
	Error           string  `json:"error,omitempty"`
}

type SystemInfo struct {
	Version      string  `json:"version"`
	GoVersion    string  `json:"go_version"`
	NumGoRoutine int     `json:"num_goroutines"`
	MemoryUsage  float64 `json:"memory_usage"`
}

type HealthReport struct {
	Status    string       `json:"status"`
	Timestamp time.Time    `json:"timestamp"`
	Database  DatabaseInfo `json:"database"`
	System    SystemInfo   `json:"system"`
}

type CreateCompanyRequest struct {
	OwnerUserID string `json:"owner_id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}
