package server

import (
	"context"
	"net/http"

	"github.com/KonferCA/NoKap/db"
	"github.com/emicklei/pgtalk/convert"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func (s *Server) setupAuthRoutes() {
	auth := s.apiV1.Group("/auth")
	auth.POST("/signup", s.handleSignup)
	auth.POST("/signin", s.handleSignin)
}

func (s *Server) handleSignup(c echo.Context) error {
	var req SignupRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	ctx := context.Background()
	existingUser, err := s.queries.GetUserByEmail(ctx, req.Email)
	if err == nil && existingUser.ID.Valid {
		return echo.NewHTTPError(http.StatusConflict, "email already registered")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to hash password")
	}

	user, err := s.queries.CreateUser(ctx, db.CreateUserParams{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		FirstName:    pgtype.Text{String: req.FirstName, Valid: true},
		LastName:     pgtype.Text{String: req.LastName, Valid: true},
		Role:         req.Role,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create user")
	}

	userID := convert.UUIDToString(user.ID)
	accessToken, refreshToken, err := generateJWT(userID, user.Role)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to generate token")
	}

	return c.JSON(http.StatusCreated, AuthResponse{
		AcessToken:   accessToken,
		RefreshToken: refreshToken,
		User: User{
			ID:            userID,
			Email:         user.Email,
			FirstName:     user.FirstName.String,
			LastName:      user.LastName.String,
			Role:          user.Role,
			WalletAddress: getStringPtr(user.WalletAddress),
		},
	})
}

func (s *Server) handleSignin(c echo.Context) error {
	var req SigninRequest
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body")
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	ctx := context.Background()
	user, err := s.queries.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "invalid credentials")
	}

	userID := convert.UUIDToString(user.ID)
	accessToken, refreshToken, err := generateJWT(userID, user.Role)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to generate token")
	}

	return c.JSON(http.StatusOK, AuthResponse{
		AcessToken:   accessToken,
		RefreshToken: refreshToken,
		User: User{
			ID:            userID,
			Email:         user.Email,
			FirstName:     user.FirstName.String,
			LastName:      user.LastName.String,
			Role:          user.Role,
			WalletAddress: getStringPtr(user.WalletAddress),
		},
	})
}

// helper function to convert pgtype.Text to *string
func getStringPtr(t pgtype.Text) *string {
	if !t.Valid {
		return nil
	}
	return &t.String
}
