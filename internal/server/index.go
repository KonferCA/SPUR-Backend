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
