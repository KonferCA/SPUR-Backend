package middleware

import (
	"net/http"
	"strings"

	"github.com/KonferCA/NoKap/internal/jwt"
	"github.com/labstack/echo/v4"
)

const JWT_CLAIMS = "MIDDLEWARE_JWT_CLAIMS"

// Middleware that validate the "Authorization" header for a Bearer token.
func ProtectAPI() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authorization := c.Request().Header.Get(echo.HeaderAuthorization)
			parts := strings.Split(authorization, " ")
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid authorization header. Only accept Bearer token.")
			}
			claims, err := jwt.VerifyToken(parts[0])
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
			}
			c.Set(JWT_CLAIMS, claims)
			return next(c)
		}
	}
}
