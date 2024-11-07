package server

import "github.com/labstack/echo/v4"

type RouterGroup interface {
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

type BasicResponse struct {
	Message   string `json:"message"`
	RequestID string `json:"request_id,omitempty"`
}

type HealthReport struct {
	Database bool   `json:"database"`
	Version  string `json:"version"`
}
