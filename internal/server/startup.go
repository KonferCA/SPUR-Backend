package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonferCA/NoKap/db"
	"github.com/labstack/echo/v4"
)

func (s *Server) handleCreateStartup(c echo.Context) error {
	var req CreateStartupRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body :(")
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	queries := db.New(s.DBPool)
	params := db.CreateStartupParams{
		OwnerID: req.OwnerID,
		Name:    req.Name,
		Status:  "active",
	}

	startup, err := queries.CreateStartup(context.Background(), params)
	if err != nil {
		fmt.Printf("Database error: %v\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to create startup: %v", err))
	}

	return c.JSON(http.StatusCreated, startup)
}

func (s *Server) handleGetStartup(c echo.Context) error {
	queries := db.New(s.DBPool)

	startups, err := queries.GetStartups(context.Background())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch startups :(")
	}

	return c.JSON(http.StatusOK, startups)
}
