package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KonferCA/NoKap/db"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

func (s *Server) handleCreateProject(c echo.Context) error {
	var req CreateProjectRequest
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

	fmt.Printf("Creating project with params: company_id=%v, title=%s, status=%s\n",
		companyID, req.Title, req.Status)

	params := db.CreateProjectParams{
		CompanyID: companyID,
		Title:     req.Title,
		Description: pgtype.Text{
			String: req.Description,
			Valid:  req.Description != "",
		},
		Status: req.Status,
	}

	project, err := queries.CreateProject(context.Background(), params)
	if err != nil {
		fmt.Printf("Error creating project: %v\n", err)
		return handleDBError(err, "create", "project")
	}

	fmt.Printf("Created project: %+v\n", project)

	return c.JSON(http.StatusCreated, project)
}

func (s *Server) handleGetProject(c echo.Context) error {
	projectID, err := validateUUID(c.Param("id"), "project")
	if err != nil {
		return err
	}

	queries := db.New(s.DBPool)
	project, err := queries.GetProject(context.Background(), projectID)
	if err != nil {
		return handleDBError(err, "fetch", "project")
	}

	return c.JSON(http.StatusOK, project)
}

func (s *Server) handleListProjects(c echo.Context) error {
	queries := db.New(s.DBPool)
	companyID := c.QueryParam("company_id")

	if companyID != "" {
		companyUUID, err := validateUUID(companyID, "company")
		if err != nil {
			return err
		}
		projects, err := queries.ListProjectsByCompany(context.Background(), companyUUID)
		if err != nil {
			return handleDBError(err, "fetch", "projects")
		}
		return c.JSON(http.StatusOK, projects)
	}

	projects, err := queries.ListProjects(context.Background())
	if err != nil {
		return handleDBError(err, "fetch", "projects")
	}
	return c.JSON(http.StatusOK, projects)
}

func (s *Server) handleDeleteProject(c echo.Context) error {
	projectID, err := validateUUID(c.Param("id"), "project")
	if err != nil {
		return err
	}

	queries := db.New(s.DBPool)
	_, err = queries.GetProject(context.Background(), projectID)
	if err != nil {
		return handleDBError(err, "verify", "project")
	}

	err = queries.DeleteProject(context.Background(), projectID)
	if err != nil {
		return handleDBError(err, "delete", "project")
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Server) handleCreateProjectFile(c echo.Context) error {
	projectID, err := validateUUID(c.Param("project_id"), "project")
	if err != nil {
		return err
	}

	var req CreateProjectFileRequest
	if err := validateBody(c, &req); err != nil {
		return err
	}

	queries := db.New(s.DBPool)
	_, err = queries.GetProject(context.Background(), projectID)
	if err != nil {
		return handleDBError(err, "verify", "project")
	}

	params := db.CreateProjectFileParams{
		ProjectID: projectID,
		FileType:  req.FileType,
		FileUrl:   &req.FileURL,
	}

	file, err := queries.CreateProjectFile(context.Background(), params)
	if err != nil {
		return handleDBError(err, "create", "project file")
	}

	return c.JSON(http.StatusCreated, file)
}

func (s *Server) handleListProjectFiles(c echo.Context) error {
	projectID, err := validateUUID(c.Param("project_id"), "project")
	if err != nil {
		return err
	}

	queries := db.New(s.DBPool)
	files, err := queries.ListProjectFiles(context.Background(), projectID)
	if err != nil {
		return handleDBError(err, "fetch", "project files")
	}

	return c.JSON(http.StatusOK, files)
}

func (s *Server) handleDeleteProjectFile(c echo.Context) error {
	fileID, err := validateUUID(c.Param("id"), "file")
	if err != nil {
		return err
	}

	queries := db.New(s.DBPool)
	err = queries.DeleteProjectFile(context.Background(), fileID)
	if err != nil {
		return handleDBError(err, "delete", "project file")
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Server) handleCreateProjectComment(c echo.Context) error {
	projectID, err := validateUUID(c.Param("project_id"), "project")
	if err != nil {
		return err
	}

	var req CreateProjectCommentRequest
	if err := validateBody(c, &req); err != nil {
		return err
	}

	userID, err := validateUUID(req.UserID, "user")
	if err != nil {
		return err
	}

	queries := db.New(s.DBPool)
	_, err = queries.GetProject(context.Background(), projectID)
	if err != nil {
		return handleDBError(err, "verify", "project")
	}

	params := db.CreateProjectCommentParams{
		ProjectID: projectID,
		UserID:    userID,
		Comment:   &req.Comment,
	}

	comment, err := queries.CreateProjectComment(context.Background(), params)
	if err != nil {
		return handleDBError(err, "create", "project comment")
	}

	return c.JSON(http.StatusCreated, comment)
}

func (s *Server) handleListProjectComments(c echo.Context) error {
	projectID, err := validateUUID(c.Param("project_id"), "project")
	if err != nil {
		return err
	}

	queries := db.New(s.DBPool)
	comments, err := queries.GetProjectComments(context.Background(), projectID)
	if err != nil {
		return handleDBError(err, "fetch", "project comments")
	}

	return c.JSON(http.StatusOK, comments)
}

func (s *Server) handleDeleteProjectComment(c echo.Context) error {
	commentID, err := validateUUID(c.Param("id"), "comment")
	if err != nil {
		return err
	}

	queries := db.New(s.DBPool)
	err = queries.DeleteProjectComment(context.Background(), commentID)
	if err != nil {
		return handleDBError(err, "delete", "project comment")
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Server) handleCreateProjectLink(c echo.Context) error {
	projectID, err := validateUUID(c.Param("project_id"), "project")
	if err != nil {
		return err
	}

	var req CreateProjectLinkRequest
	if err := validateBody(c, &req); err != nil {
		return err
	}

	queries := db.New(s.DBPool)
	_, err = queries.GetProject(context.Background(), projectID)
	if err != nil {
		return handleDBError(err, "verify", "project")
	}

	params := db.CreateProjectLinkParams{
		ProjectID: projectID,
		LinkType:  req.LinkType,
		Url:       &req.URL,
	}

	link, err := queries.CreateProjectLink(context.Background(), params)
	if err != nil {
		return handleDBError(err, "create", "project link")
	}

	return c.JSON(http.StatusCreated, link)
}

func (s *Server) handleListProjectLinks(c echo.Context) error {
	projectID, err := validateUUID(c.Param("project_id"), "project")
	if err != nil {
		return err
	}

	queries := db.New(s.DBPool)
	links, err := queries.ListProjectLinks(context.Background(), projectID)
	if err != nil {
		return handleDBError(err, "fetch", "project links")
	}

	return c.JSON(http.StatusOK, links)
}

func (s *Server) handleDeleteProjectLink(c echo.Context) error {
	linkID, err := validateUUID(c.Param("id"), "link")
	if err != nil {
		return err
	}

	queries := db.New(s.DBPool)
	err = queries.DeleteProjectLink(context.Background(), linkID)
	if err != nil {
		return handleDBError(err, "delete", "project link")
	}

	return c.NoContent(http.StatusNoContent)
}

func (s *Server) handleAddProjectTag(c echo.Context) error {
	projectID, err := validateUUID(c.Param("project_id"), "project")
	if err != nil {
		return err
	}

	var req AddProjectTagRequest
	if err := validateBody(c, &req); err != nil {
		return err
	}

	tagID, err := validateUUID(req.TagID, "tag")
	if err != nil {
		return err
	}

	queries := db.New(s.DBPool)

	_, err = queries.GetProject(context.Background(), projectID)
	if err != nil {
		return handleDBError(err, "verify", "project")
	}

	params := db.AddProjectTagParams{
		ProjectID: projectID,
		TagID:     tagID,
	}

	projectTag, err := queries.AddProjectTag(context.Background(), params)
	if err != nil {
		return handleDBError(err, "create", "project tag")
	}

	return c.JSON(http.StatusCreated, projectTag)
}

func (s *Server) handleListProjectTags(c echo.Context) error {
	projectID, err := validateUUID(c.Param("project_id"), "project")
	if err != nil {
		return err
	}

	queries := db.New(s.DBPool)
	tags, err := queries.ListProjectTags(context.Background(), projectID)
	if err != nil {
		return handleDBError(err, "fetch", "project tags")
	}

	return c.JSON(http.StatusOK, tags)
}

func (s *Server) handleDeleteProjectTag(c echo.Context) error {
	projectID, err := validateUUID(c.Param("project_id"), "project")
	if err != nil {
		return err
	}

	tagID, err := validateUUID(c.Param("tag_id"), "tag")
	if err != nil {
		return err
	}

	queries := db.New(s.DBPool)
	params := db.DeleteProjectTagParams{
		ProjectID: projectID,
		TagID:     tagID,
	}

	err = queries.DeleteProjectTag(context.Background(), params)
	if err != nil {
		return handleDBError(err, "delete", "project tag")
	}

	return c.NoContent(http.StatusNoContent)
}
