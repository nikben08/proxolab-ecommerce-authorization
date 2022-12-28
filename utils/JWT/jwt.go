package utils

import (
	"strconv"
	"time"
	"wiaoj/authorization/config"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func GenerateJwtToken(userID uuid.UUID, Email string, Name string, AccessLevel int) (string, error) {
	var secretKey = []byte(config.GetJwtSecretKey())
	var expirationMinutes, err = strconv.Atoi(config.GetJwtExpirationMinutes())

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = Email
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(expirationMinutes)).Unix()
	claims["level"] = AccessLevel
	claims["userId"] = userID

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}
