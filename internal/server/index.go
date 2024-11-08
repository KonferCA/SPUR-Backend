package server

import (
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/KonferCA/NoKap/db"
	"github.com/KonferCA/NoKap/internal/middleware"
)

type Server struct {
	DBPool       *pgxpool.Pool
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

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(echoMiddleware.Recover())

	e.Validator = &CustomValidator{
		validator: validator.New(),
	}

	server := &Server{
		DBPool:       pool,
		echoInstance: e,
	}

	// setup api routes
	server.setupV1Routes()
	server.setupStartupRoutes()

	// setup static routes
	server.setupStaticRoutes()

	return server, nil
}

func (s *Server) setupV1Routes() {
	s.apiV1 = s.echoInstance.Group("/api/v1")

	s.apiV1.GET("/health", s.handleHealthCheck)

	s.echoInstance.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "Server do be running...",
		})
	})

	for _, route := range s.echoInstance.Routes() {
		s.echoInstance.Logger.Printf("Route: %s %s", route.Method, route.Path)
	}
}

func (s *Server) setupStartupRoutes() {
	s.apiV1.POST("/companies", s.handleCreateCompany)
	s.apiV1.GET("/companies/:id", s.handleGetCompany)
	s.apiV1.GET("/companies", s.handleListCompanies)
	s.apiV1.GET("/companies/:id", s.handleDeleteCompany)
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
