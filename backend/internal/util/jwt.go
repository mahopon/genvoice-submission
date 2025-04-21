package util

import (
	"encoding/base64"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var SigningKey []byte

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
