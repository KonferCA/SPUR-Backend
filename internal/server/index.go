package server

import (
	"github.com/labstack/echo/v4"
)

type Server struct {
	echoInstance *echo.Echo
}

// Create a new Server instance and registers all routes and middlewares.
func New() *Server {
	server := &Server{}
	server.echoInstance = echo.New()

	return server
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
