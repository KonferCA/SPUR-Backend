package server

import (
	"mime"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) setupStaticRoutes() {
	// add mime types
	mime.AddExtensionType(".js", "application/javascript")
	mime.AddExtensionType(".css", "text/css")
	mime.AddExtensionType(".html", "text/html")

	// get static files directory from env or use default
	staticDir := os.Getenv("STATIC_FILES_DIR")
	if staticDir == "" {
		if os.Getenv("APP_ENV") == "development" {
			staticDir = "static/dist" // local development
		} else {
			staticDir = "/app/static" // docker environment
		}
	}

	// serve static files
	s.echoInstance.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       staticDir,
		Index:      "index.html",
		HTML5:      true,
		Browse:     false,
		IgnoreBase: true,
	}))

	// serve assets with correct mime types
	s.echoInstance.GET("/assets/*", func(c echo.Context) error {
		path := filepath.Join(staticDir, "assets", c.Param("*"))
		return c.File(path)
	})

	// catch all route
	s.echoInstance.GET("/*", func(c echo.Context) error {
		if strings.HasPrefix(c.Path(), "/api") {
			return echo.NotFoundHandler(c)
		}
		return c.File(filepath.Join(staticDir, "index.html"))
	})
}
