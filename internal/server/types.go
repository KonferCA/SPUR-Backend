package server

import (
	"fmt"
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
	OwnerUserID string `json:"owner_user_id" validate:"required,uuid"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
}

type CreateResourceRequestRequest struct {
	CompanyID    string `json:"company_id" validate:"required,uuid"`
	ResourceType string `json:"resource_type" validate:"required"`
	Description  string `json:"description"`
	Status       string `json:"status" validate:"required"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func NewCustomValidator() *CustomValidator {
	v := validator.New()
	return &CustomValidator{validator: v}
}

func (cv *CustomValidator) Validate(i interface{}) error {
	fmt.Printf("Validating struct: %+v\n", i)
	if err := cv.validator.Struct(i); err != nil {
		fmt.Printf("Validation error: %v\n", err)
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

type SignupRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Role      string `json:"role" validate:"required,oneof=startup_owner admin investor"`
}

type SigninRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	User         User   `json:"user"`
}

type User struct {
	ID            string  `json:"id"`
	Email         string  `json:"email"`
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	Role          string  `json:"role"`
	WalletAddress *string `json:"wallet_address,omitempty"`
}
