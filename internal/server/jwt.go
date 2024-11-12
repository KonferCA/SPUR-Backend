package server

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTClaims struct {
	UserID string `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func generateJWT(userID string, role string) (string, string, error) {
	accessTokenClaims := JWTClaims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(10 * time.Minute).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	refreshTokenClaims := JWTClaims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			// expire in 1 week
			ExpiresAt: time.Now().Add(24 * time.Hour * 7).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	jwtSecret := []byte(os.Getenv("JWT_SECRET"))

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenStr, err := accessToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenStr, err := refreshToken.SignedString(jwtSecret)
	if err != nil {
		return "", "", err
	}

	return accessTokenStr, refreshTokenStr, nil
}
