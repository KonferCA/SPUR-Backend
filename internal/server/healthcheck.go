package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// TODO: Update when DB working + connected
func (s *Server) handleHealthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, HealthReport{
		Database: true,
		Version:  "1.0.0",
	})
}
