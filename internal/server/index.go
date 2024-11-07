package server

import (
	"fmt"
	"os"

	"github.com/KonferCA/NoKap/db"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type Server struct {
	DBPool       *pgxpool.Pool
	echoInstance *echo.Echo
}

// Create a new Server instance and registers all routes and middlewares.
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

	server := &Server{
		DBPool: pool,
	}
	server.echoInstance = echo.New()

	return server, nil
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
