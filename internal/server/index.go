package server

import (
	"fmt"
	"os"

	"github.com/KonferCA/NoKap/db"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.Use(middleware.Recover())

	server := &Server{
		DBPool: pool,
	}
	server.echoInstance = e

	server.setupV1Routes()

	return server, nil
}

func (s *Server) setupV1Routes() {
	s.apiV1 = s.echoInstance.Group("/api/v1")

	s.echoInstance.GET("/", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "Server do be running...",
		})
	})

	s.apiV1.GET("/health", s.handleHealthCheck)

	for _, route := range s.echoInstance.Routes() {
		s.echoInstance.Logger.Printf("Route: %s %s", route.Method, route.Path)
	}
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
