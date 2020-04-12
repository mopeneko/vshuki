package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4/middleware"
	"time"
)

type jwtCustomClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

func GenerateConfig(secret []byte) middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &jwtCustomClaims{},
		SigningKey: secret,
	}
}

func GenerateJWT(userID string, jwtSecret []byte) (string, error) {
	claims := &jwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(jwtSecret)

	return t, err
}
