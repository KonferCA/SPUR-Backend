package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/KonferCA/NoKap/internal/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestProtectAPIMiddleware(t *testing.T) {
	os.Setenv("JWT_SECRET", "secret")
	e := echo.New()
	e.Use(ProtectAPI())

	e.GET("/protected", func(c echo.Context) error {
		return c.String(http.StatusOK, "protected resource")
	})

	// generate valid tokens
	userID := "user-id"
	role := "user-role"
	validAccessToken, validRefreshToken, err := jwt.Generate(userID, role)
	assert.Nil(t, err)

	// generate invalid tokens
	os.Setenv("JWT_SECRET", "wrong-secret")
	invalidAccessToken, invalidRefreshToken, err := jwt.Generate(userID, role)
	assert.Nil(t, err)

	// reset the secret
	os.Setenv("JWT_SECRET", "secret")

	tests := []struct {
		name         string
		expectedCode int
		token        string
	}{
		{
			name:         "Valid access token",
			expectedCode: http.StatusOK,
			token:        validAccessToken,
		},
		{
			name:         "Valid refresh token",
			expectedCode: http.StatusOK,
			token:        validRefreshToken,
		},
		{
			name:         "Invalid access token",
			expectedCode: http.StatusUnauthorized,
			token:        invalidAccessToken,
		},
		{
			name:         "Invalid refresh token",
			expectedCode: http.StatusUnauthorized,
			token:        invalidRefreshToken,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/protected", nil)
			rec := httptest.NewRecorder()
			req.Header.Set(echo.HeaderAuthorization, fmt.Sprintf("Bearer %s", test.token))
			e.ServeHTTP(rec, req)
			assert.Equal(t, test.expectedCode, rec.Code)
		})
	}
}
