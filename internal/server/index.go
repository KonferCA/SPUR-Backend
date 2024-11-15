package server

import (
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/KonferCA/NoKap/db"
	"github.com/KonferCA/NoKap/internal/middleware"
)

type Server struct {
	DBPool       *pgxpool.Pool
	queries      *db.Queries
	echoInstance *echo.Echo
	apiV1        *echo.Group
}

// Create a new Server instance and registers all routes and middlewares.
// Initialize database pool connection.
func New() (*Server, error) {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	)
	pool, err := db.NewPool(connStr)
	if err != nil {
		return nil, err
	}

	// Initialize queries
	queries := db.New(pool)

	e := echo.New()

	e.Debug = true
	e.Use(middleware.Logger())
	e.Use(echoMiddleware.Recover())

	customValidator := NewCustomValidator()
	fmt.Printf("Initializing validator: %+v\n", customValidator)
	e.Validator = customValidator

	server := &Server{
		DBPool:       pool,
		queries:      queries,
		echoInstance: e,
	}

	// setup api routes
	server.setupV1Routes()
	server.setupAuthRoutes()
	server.setupCompanyRoutes()
	server.setupResourceRequestRoutes()
	server.setupCompanyFinancialRoutes()
	server.setupEmployeeRoutes()
	server.setupCompanyDocumentRoutes()
	server.setupCompanyQuestionsAnswersRoutes()
	server.setupProjectRoutes()
	server.setupTagRoutes()
	server.setupFundingTransactionRoutes()
	server.setupMeetingRoutes()
	server.setupHealthRoutes()

	// setup static routes
	server.setupStaticRoutes()

	return server, nil
}

func (s *Server) setupV1Routes() {
	s.apiV1 = s.echoInstance.Group("/api/v1")

	s.echoInstance.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "Server do be running...",
		})
	})

	for _, route := range s.echoInstance.Routes() {
		s.echoInstance.Logger.Printf("Route: %s %s", route.Method, route.Path)
	}
}

func (s *Server) setupCompanyRoutes() {
	s.apiV1.POST("/companies", s.handleCreateCompany)
	s.apiV1.GET("/companies/:id", s.handleGetCompany)
	s.apiV1.GET("/companies", s.handleListCompanies)
	s.apiV1.DELETE("/companies/:id", s.handleDeleteCompany)
}

func (s *Server) setupResourceRequestRoutes() {
	s.apiV1.POST("/resource-requests", s.handleCreateResourceRequest)
	s.apiV1.GET("/resource-requests/:id", s.handleGetResourceRequest)
	s.apiV1.GET("/resource-requests", s.handleListResourceRequests)
	s.apiV1.PUT("/resource-requests/:id/status", s.handleUpdateResourceRequestStatus)
	s.apiV1.DELETE("/resource-requests/:id", s.handleDeleteResourceRequest)
}

func (s *Server) setupCompanyFinancialRoutes() {
	s.apiV1.POST("/companies/:id/financials", s.handleCreateCompanyFinancials)
	s.apiV1.GET("/companies/:id/financials", s.handleGetCompanyFinancials)
	s.apiV1.PUT("/companies/:id/financials", s.handleUpdateCompanyFinancials)
	s.apiV1.DELETE("/companies/:id/financials", s.handleDeleteCompanyFinancials)
	s.apiV1.GET("/companies/:id/financials/latest", s.handleGetLatestCompanyFinancials)
}

func (s *Server) setupEmployeeRoutes() {
	s.apiV1.POST("/employees", s.handleCreateEmployee)
	s.apiV1.GET("/employees", s.handleListEmployees)
	s.apiV1.GET("/employees/:id", s.handleGetEmployee)
	s.apiV1.PUT("/employees/:id", s.handleUpdateEmployee)
	s.apiV1.DELETE("/employees/:id", s.handleDeleteEmployee)
}

func (s *Server) setupCompanyDocumentRoutes() {
	s.apiV1.POST("/companies/:id/documents", s.handleCreateCompanyDocument)
	s.apiV1.GET("/companies/:id/documents", s.handleListCompanyDocuments)
	s.apiV1.GET("/documents/:id", s.handleGetCompanyDocument)
	s.apiV1.PUT("/documents/:id", s.handleUpdateCompanyDocument)
	s.apiV1.DELETE("/documents/:id", s.handleDeleteCompanyDocument)
}

func (s *Server) setupCompanyQuestionsAnswersRoutes() {
	s.apiV1.POST("/questions", s.handleCreateQuestion)
	s.apiV1.GET("/questions", s.handleListQuestions)
	s.apiV1.GET("/questions/:id", s.handleGetQuestion)

	s.apiV1.POST("/companies/:id/answers", s.handleCreateCompanyAnswer)
	s.apiV1.GET("/companies/:id/answers", s.handleListCompanyAnswers)
	s.apiV1.GET("/companies/:company_id/answers/:question_id", s.handleGetCompanyAnswer)
	s.apiV1.PUT("/companies/:company_id/answers/:question_id", s.handleUpdateCompanyAnswer)
	s.apiV1.DELETE("/companies/:company_id/answers/:question_id", s.handleDeleteCompanyAnswer)
}

func (s *Server) setupProjectRoutes() {
	s.apiV1.POST("/projects", s.handleCreateProject)
	s.apiV1.GET("/projects/:id", s.handleGetProject)
	s.apiV1.GET("/projects", s.handleListProjects)
	s.apiV1.DELETE("/projects/:id", s.handleDeleteProject)

	s.apiV1.POST("/projects/:project_id/files", s.handleCreateProjectFile)
	s.apiV1.GET("/projects/:project_id/files", s.handleListProjectFiles)
	s.apiV1.DELETE("/projects/files/:id", s.handleDeleteProjectFile)

	s.apiV1.POST("/projects/:project_id/comments", s.handleCreateProjectComment)
	s.apiV1.GET("/projects/:project_id/comments", s.handleListProjectComments)
	s.apiV1.DELETE("/projects/comments/:id", s.handleDeleteProjectComment)

	s.apiV1.POST("/projects/:project_id/links", s.handleCreateProjectLink)
	s.apiV1.GET("/projects/:project_id/links", s.handleListProjectLinks)
	s.apiV1.DELETE("/projects/links/:id", s.handleDeleteProjectLink)

	s.apiV1.POST("/projects/:project_id/tags", s.handleAddProjectTag)
	s.apiV1.GET("/projects/:project_id/tags", s.handleListProjectTags)
	s.apiV1.DELETE("/projects/:project_id/tags/:tag_id", s.handleDeleteProjectTag)
}

func (s *Server) setupTagRoutes() {
	s.apiV1.POST("/tags", s.handleCreateTag)
	s.apiV1.GET("/tags/:id", s.handleGetTag)
	s.apiV1.GET("/tags", s.handleListTags)
	s.apiV1.DELETE("/tags/:id", s.handleDeleteTag)
}

func (s *Server) setupFundingTransactionRoutes() {
	s.apiV1.POST("/funding-transactions", s.handleCreateFundingTransaction)
	s.apiV1.GET("/funding-transactions/:id", s.handleGetFundingTransaction)
	s.apiV1.GET("/funding-transactions", s.handleListFundingTransactions)
	s.apiV1.PUT("/funding-transactions/:id/status", s.handleUpdateFundingTransactionStatus)
	s.apiV1.DELETE("/funding-transactions/:id", s.handleDeleteFundingTransaction)
}

func (s *Server) setupMeetingRoutes() {
	s.apiV1.POST("/meetings", s.handleCreateMeeting)
	s.apiV1.GET("/meetings/:id", s.handleGetMeeting)
	s.apiV1.GET("/meetings", s.handleListMeetings)
	s.apiV1.PUT("/meetings/:id", s.handleUpdateMeeting)
	s.apiV1.DELETE("/meetings/:id", s.handleDeleteMeeting)
}

func (s *Server) setupHealthRoutes() {
	s.apiV1.GET("/health", s.handleHealthCheck)
}

// Start listening at the given address.
//
// Example:
//
// s := server.New()
// log.Fatal(s.Listen(":8080")) // listen on port 8080
func (s *Server) Listen(address string) error {
	return s.echoInstance.Start(address)
}
