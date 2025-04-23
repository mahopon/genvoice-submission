package util

import (
	"encoding/base64"
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var SigningKey []byte

type Claims struct {
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func init() {
	var err error
	SigningKey, err = base64.StdEncoding.DecodeString(os.Getenv("JWT_SECRET"))
	if err != nil {
		log.Fatal("Error decoding signing key: ", err)
	}
}

func CreateRefreshToken(userID uuid.UUID, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  userID.String(),
		"role": role,
		"exp":  time.Now().Add(24 * time.Hour).Unix(),
	})
	return token.SignedString(SigningKey)
}

func CreateAccessToken(userID uuid.UUID, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  userID.String(),
		"role": role,
		"exp":  time.Now().Add(15 * time.Minute).Unix(),
	})
	return token.SignedString(SigningKey)
}

func ValidateJWT(tokenStr string) (*Claims, error) {
	// Parse the token
	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return SigningKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
			return nil, errors.New("token has expired")
		}
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
