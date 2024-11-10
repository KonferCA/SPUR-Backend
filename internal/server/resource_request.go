package server

import (
	"context"
	"net/http"

	"github.com/KonferCA/NoKap/db"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

func (s *Server) handleCreateResourceRequest(c echo.Context) error {
	var req CreateResourceRequestRequest
	if err := validateBody(c, &req); err != nil {
		return err
	}

	companyID, err := validateUUID(req.CompanyID, "company")
	if err != nil {
		return err
	}

	queries := db.New(s.DBPool)
	_, err = queries.GetCompanyByID(context.Background(), companyID)
	if err != nil {
		return handleDBError(err, "verify", "company")
	}

	params := db.CreateResourceRequestParams{
		CompanyID:    companyID,
		ResourceType: req.ResourceType,
		Description:  pgtype.Text{String: req.Description, Valid: req.Description != ""},
		Status:       req.Status,
	}

	request, err := queries.CreateResourceRequest(context.Background(), params)
	if err != nil {
		return handleDBError(err, "create", "resource request")
	}

	return c.JSON(http.StatusCreated, request)
}

func (s *Server) handleGetResourceRequest(c echo.Context) error {
	requestID, err := validateUUID(c.Param("id"), "resource request")
	if err != nil {
		return err
	}

	queries := db.New(s.DBPool)
	request, err := queries.GetResourceRequestByID(context.Background(), requestID)
	if err != nil {
		return handleDBError(err, "fetch", "resource request")
	}

	return c.JSON(http.StatusOK, request)
}

func (s *Server) handleListResourceRequests(c echo.Context) error {
	companyID := c.QueryParam("company_id")
	queries := db.New(s.DBPool)

	if companyID != "" {
		companyUUID, err := validateUUID(companyID, "company")
		if err != nil {
			return err
		}

		_, err = queries.GetCompanyByID(context.Background(), companyUUID)
		if err != nil {
			return handleDBError(err, "verify", "company")
		}

		requests, err := queries.ListResourceRequestsByCompany(context.Background(), companyUUID)
		if err != nil {
			return handleDBError(err, "fetch", "resource requests")
		}

		return c.JSON(http.StatusOK, requests)
	}

	requests, err := queries.ListResourceRequests(context.Background())
	if err != nil {
		return handleDBError(err, "fetch", "resource requests")
	}

	return c.JSON(http.StatusOK, requests)
}

func (s *Server) handleUpdateResourceRequestStatus(c echo.Context) error {
	requestID, err := validateUUID(c.Param("id"), "resource request")
	if err != nil {
		return err
	}

	var status struct {
		Status string `json:"status" validate:"required"`
	}
	if err := validateBody(c, &status); err != nil {
		return err
	}

	queries := db.New(s.DBPool)
	_, err = queries.GetResourceRequestByID(context.Background(), requestID)
	if err != nil {
		return handleDBError(err, "verify", "resource request")
	}

	request, err := queries.UpdateResourceRequestStatus(context.Background(), db.UpdateResourceRequestStatusParams{
		ID:     requestID,
		Status: status.Status,
	})
	if err != nil {
		return handleDBError(err, "update", "resource request status")
	}

	return c.JSON(http.StatusOK, request)
}

func (s *Server) handleDeleteResourceRequest(c echo.Context) error {
	requestID, err := validateUUID(c.Param("id"), "resource request")
	if err != nil {
		return err
	}

	queries := db.New(s.DBPool)
	_, err = queries.GetResourceRequestByID(context.Background(), requestID)
	if err != nil {
		return handleDBError(err, "verify", "resource request")
	}

	err = queries.DeleteResourceRequest(context.Background(), requestID)
	if err != nil {
		handleDBError(err, "delete", "resource request")
	}

	return c.NoContent(http.StatusNoContent)
}