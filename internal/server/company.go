package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonferCA/NoKap/db"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

func (s *Server) handleCreateCompany(c echo.Context) error {
	var req CreateCompanyRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body :(")
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var ownerUUID pgtype.UUID
	if err := ownerUUID.Scan(req.OwnerUserID); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid owner ID format :(")
	}

	queries := db.New(s.DBPool)
	params := db.CreateCompanyParams{
		OwnerUserID: ownerUUID,
		Name:        req.Name,
		Description: pgtype.Text{String: req.Description, Valid: req.Description != ""},
	}

	company, err := queries.CreateCompany(context.Background(), params)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to create company: %v", err))
	}

	return c.JSON(http.StatusCreated, company)
}

func (s *Server) handleGetCompany(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing company ID :(")
	}

	var companyID pgtype.UUID
	if err := companyID.Scan(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid company ID format :(")
	}

	queries := db.New(s.DBPool)
	company, err := queries.GetCompanyByID(context.Background(), companyID)
	if err != nil {
		if isNoRowsError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "Company not found :(")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch company :(")
	}

	return c.JSON(http.StatusOK, company)
}

func (s *Server) handleListCompanies(c echo.Context) error {
	queries := db.New(s.DBPool)

	companies, err := queries.ListCompanies(context.Background())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch companies :(")
	}

	return c.JSON(http.StatusOK, companies)
}

func (s *Server) handleDeleteCompany(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing company ID :(")
	}

	var companyID pgtype.UUID
	if err := companyID.Scan(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid company ID format :(")
	}

	queries := db.New(s.DBPool)
	_, err := queries.GetCompanyByID(context.Background(), companyID)
	if err != nil {
		if isNoRowsError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "Company not found :(")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to verify company :(")
	}

	err = queries.DeleteCompany(context.Background(), companyID)
	if err != nil {
		if isNoRowsError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "Company not found :(")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete company :(")
	}

	return c.NoContent(http.StatusNoContent)
}
