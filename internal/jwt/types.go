package jwt

import golangJWT "github.com/golang-jwt/jwt/v5"

type JWTClaims struct {
	UserID    string `json:"user_id"`
	Role      string `json:"role"`
	TokenType string `json:"token_type"`
	golangJWT.RegisteredClaims
}
