package server

import (
	"context"
	"net/http"

	"github.com/KonferCA/NoKap/db"
	"github.com/labstack/echo/v4"
)

func (s *Server) handleCreateTag(c echo.Context) error {
	var req CreateTagRequest
	if err := validateBody(c, &req); err != nil {
		return err
	}

	queries := db.New(s.DBPool)
	tag, err := queries.CreateTag(context.Background(), req.Name)
	if err != nil {
		return handleDBError(err, "create", "tag")
	}

	return c.JSON(http.StatusCreated, tag)
}

func (s *Server) handleGetTag(c echo.Context) error {
	tagID, err := validateUUID(c.Param("id"), "tag")
	if err != nil {
		return err
	}

	queries := db.New(s.DBPool)
	tag, err := queries.GetTag(context.Background(), tagID)
	if err != nil {
		return handleDBError(err, "fetch", "tag")
	}

	return c.JSON(http.StatusOK, tag)
}

func (s *Server) handleListTags(c echo.Context) error {
	queries := db.New(s.DBPool)
	tags, err := queries.ListTags(context.Background())
	if err != nil {
		return handleDBError(err, "fetch", "tags")
	}

	return c.JSON(http.StatusOK, tags)
}

func (s *Server) handleDeleteTag(c echo.Context) error {
	tagID, err := validateUUID(c.Param("id"), "tag")
	if err != nil {
		return err
	}

	queries := db.New(s.DBPool)
	err = queries.DeleteTag(context.Background(), tagID)
	if err != nil {
		return handleDBError(err, "delete", "tag")
	}

	return c.NoContent(http.StatusNoContent)
}
