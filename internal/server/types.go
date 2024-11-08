package server

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type HealthReport struct {
	Database bool   `json:"database"`
	Version  string `json:"version"`
}

type CreateStartupRequest struct {
	Name    string `json:"name" validate:"required"`
	OwnerID int32  `json:"owner_id" validate:"required"`
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
