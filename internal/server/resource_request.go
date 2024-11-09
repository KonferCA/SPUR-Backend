package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonferCA/NoKap/db"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

func (s *Server) handleCreateResourceRequest(c echo.Context) error {
	var req CreateResourceRequestRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body :(")
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var companyID pgtype.UUID
	if err := companyID.Scan(req.CompanyID); err != nil {
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

	params := db.CreateResourceRequestParams{
		CompanyID:    companyID,
		ResourceType: req.ResourceType,
		Description:  pgtype.Text{String: req.Description, Valid: req.Description != ""},
		Status:       req.Status,
	}

	request, err := queries.CreateResourceRequest(context.Background(), params)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to create resource request: %v", err))
	}

	return c.JSON(http.StatusCreated, request)
}

func (s *Server) handleGetResourceRequest(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing resource request ID :(")
	}

	var requestID pgtype.UUID
	if err := requestID.Scan(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid resource request ID format :(")
	}

	queries := db.New(s.DBPool)
	request, err := queries.GetResourceRequestByID(context.Background(), requestID)
	if err != nil {
		if isNoRowsError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "Resource request not found :(")
		}
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch resource request :(")
	}

	return c.JSON(http.StatusOK, request)
}

func (s *Server) handleListResourceRequests(c echo.Context) error {
	companyID := c.QueryParam("company_id")
	queries := db.New(s.DBPool)

	if companyID != "" {
		var companyUUID pgtype.UUID
		if err := companyUUID.Scan(companyID); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid company ID format :(")
		}

		_, err := queries.GetCompanyByID(context.Background(), companyUUID)
		if err != nil {
			if isNoRowsError(err) {
				return echo.NewHTTPError(http.StatusNotFound, "Company not found :(")
			}

			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to verify company :(")
		}

		requests, err := queries.ListResourceRequestsByCompany(context.Background(), companyUUID)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch resource requests :(")
		}

		return c.JSON(http.StatusOK, requests)
	}

	requests, err := queries.ListResourceRequests(context.Background())
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to fetch resource requests :(")
	}

	return c.JSON(http.StatusOK, requests)
}

func (s *Server) handleUpdateResourceRequestStatus(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing resource request ID :(")
	}

	var requestID pgtype.UUID
	if err := requestID.Scan(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid resource request ID format :(")
	}

	var status struct {
		Status string `json:"status" validate:"required"`
	}
	if err := c.Bind(&status); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body :(")
	}
	if err := c.Validate(status); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	queries := db.New(s.DBPool)
	_, err := queries.GetResourceRequestByID(context.Background(), requestID)
	if err != nil {
		if isNoRowsError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "Resource request not found :(")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to verify resource request :(")
	}

	request, err := queries.UpdateResourceRequestStatus(context.Background(), db.UpdateResourceRequestStatusParams{
		ID:     requestID,
		Status: status.Status,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to update resource request status :(")
	}

	return c.JSON(http.StatusOK, request)
}

func (s *Server) handleDeleteResourceRequest(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Missing resource request ID :(")
	}

	var requestID pgtype.UUID
	if err := requestID.Scan(id); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid resource request ID format :(")
	}

	queries := db.New(s.DBPool)
	_, err := queries.GetResourceRequestByID(context.Background(), requestID)
	if err != nil {
		if isNoRowsError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "Resource request not found :(")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to verify resource request :(")
	}

	err = queries.DeleteResourceRequest(context.Background(), requestID)
	if err != nil {
		if isNoRowsError(err) {
			return echo.NewHTTPError(http.StatusNotFound, "Resource request not found :(")
		}

		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete resource request :(")
	}

	return c.NoContent(http.StatusNoContent)
}
